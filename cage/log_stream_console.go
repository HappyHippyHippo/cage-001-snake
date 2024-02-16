package cage

import (
	"fmt"
	"os"
)

type logStreamConsole struct {
	logStream
}

var (
	_ ITraversable = &logStreamConsole{}
  _ IPubSub      = &logStream{}
	_ ILoggable    = &logStream{}
	_ ILogStreamer = &logStreamConsole{}
)

// NewLogStreamConsole @todo doc
func NewLogStreamConsole() ILogStreamer {
	s := &logStreamConsole{}
	s.Init(s)
	return s
}

func (s *logStreamConsole) Init(self ...ILogStreamer) ILogStreamer {
	var ref ILogStreamer = s
	if len(self) > 0 {
		ref = self[0]
	}

	return s.logStream.Init("console", ref)
}

func (s *logStreamConsole) Broadcast(level LogLevel, msg string, context ...LogCtx) {
	if s.level == LogNone || s.level > level {
		return
	}

	if s.formatter != nil {
		msg = s.formatter(level, msg, context...)
	}

	_, _ = fmt.Fprintln(os.Stdout, msg)
}
