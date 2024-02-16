package cage

import (
	"maps"
	"sync"
)

// ResLoaderOkCallback @todo doc
type ResLoaderOkCallback func(id, url string, bytes []byte) error

// ResLoaderErrorCallback @todo doc
type ResLoaderErrorCallback func(id string, e error) error

// IResLoader @todo doc
type IResLoader interface {
	ITraversable
	IPubSub

	Close() error
	Queue(id, url string, okCallback ResLoaderOkCallback, errorCallback ResLoaderErrorCallback) error
	Load() error
}

type resLoaderReg struct {
	id            string
	url           string
	okCallback    ResLoaderOkCallback
	errorCallback ResLoaderErrorCallback
}

// ResLoader @todo doc
type ResLoader struct {
	Traversable
	PubSub

	resQueueLock sync.Locker
	resQueue     map[string]resLoaderReg
	chLoad       chan string
	chClose      chan struct{}
	loadWG       sync.WaitGroup
	logger       ILogger
	images       ImageManager
}

var _ ITraversable = &ResLoader{}
var _ IResLoader = &ResLoader{}

// Init @todo doc
func (rl *ResLoader) Init(logger ILogger, self ...IResLoader) IResLoader {
	var ref IResLoader = rl
	if len(self) > 0 {
		ref = self[0]
	}

	rl.Traversable.Init("res", ref)
	rl.PubSub.Init(ref)

	rl.resQueueLock = &sync.Mutex{}
	rl.resQueue = map[string]resLoaderReg{}
	rl.chLoad = make(chan string)
	rl.chClose = make(chan struct{})
	rl.loadWG.Add(1)
	rl.logger = logger

	_ = ref.AddChild(rl.images.Init(rl))

	go func() {
		for {
			select {
			case id := <-rl.chLoad:
				_ = rl.load(id)
			case <-rl.chClose:
				rl.loadWG.Done()
				return
			}
		}
	}()

	return ref
}

// Close @todo doc
func (rl *ResLoader) Close() error {
	rl.chClose <- struct{}{}
	rl.loadWG.Wait()
	return nil
}

// Queue @todo doc
func (rl *ResLoader) Queue(id, url string, okCallback ResLoaderOkCallback, errorCallback ResLoaderErrorCallback) error {
	rl.resQueueLock.Lock()
	defer rl.resQueueLock.Unlock()

  if _, ok := rl.resQueue[id]; ok {
		return errResAlreadyQueued(id)
	}

  rl.resQueue[id] = resLoaderReg{id: id, url: url, okCallback: okCallback, errorCallback: errorCallback}
	return nil
}

// Load @todo doc
func (rl *ResLoader) Load() error {
	rl.resQueueLock.Lock()
	queue := maps.Clone(rl.resQueue)
	rl.resQueueLock.Unlock()

  for id := range queue {
		rl.chLoad <- id
	}
	return nil
}

func (rl *ResLoader) load(id string) error {
	rl.resQueueLock.Lock()
	defer rl.resQueueLock.Unlock()

  reg := rl.resQueue[id]
	bytes, e := rl.loadBytes(reg.url)
	if e != nil {
		rl.logger.Signal(LogError, "cage", "Error loading the resource '%s' : %s", LogCtxMsgArgs(id, e.Error()))
		if reg.errorCallback != nil {
			return reg.errorCallback(id, e)
		}
	}

  if reg.okCallback != nil {
		rl.logger.Signal(LogInfo, "cage", "resource '%s' loaded", LogCtxMsgArgs(id))
		if e := reg.okCallback(id, reg.url, bytes); e != nil {
			return e
		}
	}

  delete(rl.resQueue, id)
	return nil
}
