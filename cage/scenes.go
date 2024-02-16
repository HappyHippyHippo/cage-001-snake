package cage

import (
	"sync"
)

// SceneFactory @todo doc
type SceneFactory func() ITraversable

// ISceneManager @todo doc
type ISceneManager interface {
	ITraversable
	IPubSub

	Close() error
	Update(args UpdateArgs) error
	Render(args RenderArgs) error

	HasScene(id string) bool
	AddScene(id string, factory SceneFactory) error
	QueueScene(id string) error
	RemoveScene(id string) error
}

type sceneManagerReg struct {
	factory SceneFactory
	scene   ITraversable
}

// SceneManager @todo doc
type SceneManager struct {
	Traversable
	PubSub

	self       ISceneManager
	logger     ILogger
	scenesLock sync.Locker
	scenes     map[string]sceneManagerReg
	isQueued   bool
	queued     string
	current    ITraversable
}

var (
	_ ITraversable  = &SceneManager{}
	_ ISceneManager = &SceneManager{}
)

// Init @todo doc
func (sm *SceneManager) Init(logger ILogger, self ...ISceneManager) ISceneManager {
	var ref ISceneManager = sm
	if len(self) > 0 {
		ref = self[0]
	}

	sm.Traversable.Init("scenes", ref)
	sm.PubSub.Init(ref)

	sm.logger = logger
	sm.scenesLock = &sync.Mutex{}
	sm.scenes = map[string]sceneManagerReg{}
	sm.isQueued = false
	sm.queued = ""
	sm.current = nil

	return ref
}

// Close @todo doc
func (sm *SceneManager) Close() error {
	if sm.current != nil {
		if pub, ok := sm.current.(IPubSub); ok {
			sm.logger.Signal(LogInfo, "cage", "sleeping current scene : %s", LogCtxMsgArgs(sm.current.ID()))
			_ = pub.Publish(SignalSleep, nil, PropagatePrePublish)
		}
		sm.logger.Signal(LogDebug, "cage", "removing current scene from tree : %s", LogCtxMsgArgs(sm.current.ID()))
		_ = sm.RemoveChild(sm.current.ID())
		sm.current = nil
	}

	for id, reg := range sm.scenes {
		if reg.scene != nil {
			if pub, ok := reg.scene.(IPubSub); ok {
				sm.logger.Signal(LogInfo, "cage", "closing scene : %s", LogCtxMsgArgs(id))
				_ = pub.Publish(SignalClose, nil, PropagatePrePublish)
			}
		}
	}

	sm.scenes = map[string]sceneManagerReg{}
	return nil
}

// HasScene @todo doc
func (sm *SceneManager) HasScene(id string) bool {
	_, exists := sm.scenes[id]
	return exists
}

// AddScene @todo doc
func (sm *SceneManager) AddScene(id string, factory SceneFactory) error {
	if _, exists := sm.scenes[id]; exists {
		return errDuplicateScene(id)
	}

	sm.scenes[id] = sceneManagerReg{
		factory: factory,
		scene:   nil,
	}

	return nil
}

// QueueScene @todo doc
func (sm *SceneManager) QueueScene(id string) error {
	_, exists := sm.scenes[id]
	switch {
	case !exists:
		return errSceneNotFound(id)
	case sm.isQueued && sm.queued == id:
	case sm.current != nil && sm.current.ID() == id:
		return nil
	}

	sm.isQueued = true
	sm.queued = id
	return nil
}

// RemoveScene @todo doc
func (sm *SceneManager) RemoveScene(id string) error {
	_, exists := sm.scenes[id]
	switch {
	case !exists:
		return errSceneNotFound(id)
	case sm.isQueued && sm.queued == id:
		return errSceneQueued(id)
	case sm.current != nil && sm.current.ID() == id:
		return errSceneRunning(id)
	}

	if reg := sm.scenes[id]; reg.scene != nil {
		if pub, ok := reg.scene.(IPubSub); ok {
			sm.logger.Signal(LogInfo, "cage", "closing scene : %s", LogCtxMsgArgs(id))
			_ = pub.Publish(SignalClose, nil, PropagatePrePublish)
		}
		delete(sm.scenes, id)
	}
	return nil
}

// Update @todo doc
func (sm *SceneManager) Update(args UpdateArgs) error {
	sm.scenesLock.Lock()
	defer func() { sm.scenesLock.Unlock() }()

	sm.processQueue()
	if sm.current == nil {
		return ErrBreak
	}

	if pub, ok := sm.current.(IPubSub); ok {
		sm.logger.Signal(LogDebug, "cage", "updating scene : %s", LogCtxMsgArgs(sm.current.ID()))
		return pub.Publish(SignalUpdate, args, PropagatePostPublish)
	}
	return nil
}

// Render @todo doc
func (sm *SceneManager) Render(args RenderArgs) error {
	sm.scenesLock.Lock()
	defer func() { sm.scenesLock.Unlock() }()

	if sm.current == nil {
		return ErrBreak
	}

	if pub, ok := sm.current.(IPubSub); ok {
		sm.logger.Signal(LogDebug, "cage", "rendering scene : %s", LogCtxMsgArgs(sm.current.ID()))
		return pub.Publish(SignalRender, args, PropagatePostPublish)
	}
	return nil
}

func (sm *SceneManager) processQueue() {
	if !sm.isQueued {
		return
	}

	if sm.current != nil {
		if pub, ok := sm.current.(IPubSub); ok {
			sm.logger.Signal(LogInfo, "cage", "sleeping current scene : %s", LogCtxMsgArgs(sm.current.ID()))
			_ = pub.Publish(SignalSleep, nil, PropagatePrePublish)
		}
		_ = sm.RemoveChild(sm.current.ID())
		sm.current = nil
	}

	reg := sm.scenes[sm.queued]
	if reg.scene == nil {
		sm.scenes[sm.queued] = sceneManagerReg{
			factory: reg.factory,
			scene:   reg.factory(),
		}
	}

	sm.isQueued = false
	sm.current = sm.scenes[sm.queued].scene
	sm.logger.Signal(LogDebug, "cage", "adding the new current scene into the tree : %s", LogCtxMsgArgs(sm.current.ID()))
	_ = sm.AddChild(sm.current)

	if pub, ok := sm.current.(IPubSub); ok {
		sm.logger.Signal(LogInfo, "cage", "awaking the new current scene : %s", LogCtxMsgArgs(sm.current.ID()))
		_ = pub.Publish(SignalAwake, nil, PropagatePostPublish)
	}
}
