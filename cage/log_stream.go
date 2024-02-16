package cage

import (
	"sort"
)

// ILogStreamer @todo doc
type ILogStreamer interface {
	ITraversable
	IPubSub
	ILoggable

	GetLevel() LogLevel
	SetLevel(level LogLevel)

	SetFormatter(formatter LogFormatter)

	ListChannels() []string
	HasChannel(channel string) bool
	AddChannel(channel string)
	RemoveChannel(channel string)
}

type logStream struct {
	Traversable
	PubSub

	level     LogLevel
	formatter LogFormatter
	channels  []string
}

var (
	_ ITraversable = &logStream{}
	_ IPubSub      = &logStream{}
	_ ILoggable    = &logStream{}
	_ ILogStreamer = &logStream{}
)

type logSignalMsgArgs struct {
	level   LogLevel
	channel string
	msg     string
	context LogCtx
}

type logBroadcastMsgArgs struct {
	level   LogLevel
	msg     string
	context LogCtx
}

func (s *logStream) Init(id string, self ILogStreamer) ILogStreamer {
	s.Traversable.Init(id, self)
	s.PubSub.Init(self)

	s.level = LogError
	s.formatter = logFormatter
	s.channels = []string{}

	_ = self.Subscribe(SignalLogChannel, "cage", func(args interface{}) error {
		if a, ok := args.(logSignalMsgArgs); ok {
			self.Signal(a.level, a.channel, a.msg, a.context)
		}
		return nil
	})
	_ = self.Subscribe(SignalLogBroadcast, "cage", func(args interface{}) error {
		if a, ok := args.(logBroadcastMsgArgs); ok {
			self.Broadcast(a.level, a.msg, a.context)
		}
		return nil
	})

	return self
}

func (s *logStream) GetLevel() LogLevel {
	return s.level
}

func (s *logStream) SetLevel(level LogLevel) {
	s.level = level
}

func (s *logStream) SetFormatter(formatter LogFormatter) {
	s.formatter = formatter
}

func (s *logStream) ListChannels() []string {
	return s.channels
}

func (s *logStream) HasChannel(channel string) bool {
	i := sort.SearchStrings(s.channels, channel)
	return i < len(s.channels) && s.channels[i] == channel
}

func (s *logStream) AddChannel(channel string) {
	i := sort.SearchStrings(s.channels, channel)
	if ok := i < len(s.channels) && s.channels[i] == channel; !ok {
		s.channels = append(s.channels, channel)
		sort.Strings(s.channels)
	}
}

func (s *logStream) RemoveChannel(channel string) {
	i := sort.SearchStrings(s.channels, channel)
	if ok := i < len(s.channels) && s.channels[i] == channel; ok {
		s.channels = append(s.channels[:i], s.channels[i+1:]...)
	}
}

func (s *logStream) Signal(level LogLevel, channel, msg string, context ...LogCtx) {
	i := sort.SearchStrings(s.channels, channel)
	if i == len(s.channels) || s.channels[i] != channel {
		return
	}

	s.self.(ILogStreamer).Broadcast(level, msg, context...)
}

func (*logStream) Broadcast(_ LogLevel, _ string, _ ...LogCtx) {
	// no-op
}
