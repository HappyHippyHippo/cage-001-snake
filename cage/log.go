package cage

import (
	"io"
)

// ILogger @todo doc
type ILogger interface {
	ITraversable
	IPubSub
	ILoggable

	Close() error

	HasStream(id string) bool
	AddStream(stream ILogStreamer) error
	RemoveStream(id string) error
}

// Logger @todo doc
type Logger struct {
	Traversable
	PubSub
}

var (
	_ ITraversable = &Logger{}
	_ IPubSub      = &Logger{}
	_ ILoggable    = &Logger{}
	_ ILogger      = &Logger{}
)

// Init @todo doc
func (l *Logger) Init(self ...ILogger) ILogger {
	var ref ILogger = l
	if len(self) > 0 {
		ref = self[0]
	}

	l.Traversable.Init("logger", ref)
	l.PubSub.Init(ref)

	_ = l.Subscribe(SignalLogChannel, "cage", func(args interface{}) error {
		if a, ok := args.(logSignalMsgArgs); ok {
			ref.Signal(a.level, a.channel, a.msg, a.context)
		}
		return nil
	})
	_ = l.Subscribe(SignalLogBroadcast, "cage", func(args interface{}) error {
		if a, ok := args.(logBroadcastMsgArgs); ok {
			ref.Broadcast(a.level, a.msg, a.context)
		}
		return nil
	})

	return ref
}

// Close @todo doc
func (l *Logger) Close() error {
	for _, stream := range l.Traversable.children {
		if closer, ok := stream.(io.Closer); ok {
			_ = closer.Close()
		}
	}
	l.Traversable.children = []ITraversable{}
	return nil
}

// HasStream @todo doc
func (l *Logger) HasStream(id string) bool {
	return l.self.HasChild(id)
}

// AddStream @todo doc
func (l *Logger) AddStream(stream ILogStreamer) error {
	return l.self.AddChild(stream)
}

// RemoveStream @todo doc
func (l *Logger) RemoveStream(id string) error {
	return l.self.RemoveChild(id)
}

// Signal @todo doc
func (l *Logger) Signal(level LogLevel, channel, msg string, context ...LogCtx) {
	args := logSignalMsgArgs{level: level, channel: channel, msg: msg}
	if len(context) > 0 {
		args.context = context[0]
	}
	l.self.(IPubSub).Publish(SignalLogChannel, args)
}

// Broadcast @todo doc
func (l *Logger) Broadcast(level LogLevel, msg string, context ...LogCtx) {
	args := logBroadcastMsgArgs{level: level, msg: msg}
	if len(context) > 0 {
		args.context = context[0]
	}
	l.self.(IPubSub).Publish(SignalLogBroadcast, args)
}
