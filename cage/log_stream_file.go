package cage

import (
	"fmt"
	"os"
)

type logStreamFile struct {
	logStream

	path string
}

var (
	_ ITraversable = &logStreamFile{}
  _ IPubSub      = &logStream{}
	_ ILoggable    = &logStream{}
	_ ILogStreamer = &logStreamFile{}
)

// NewLogStreamFile @todo doc
func NewLogStreamFile(path string) ILogStreamer {
	s := &logStreamFile{}
	s.Init(path, s)
	return s
}

func (s *logStreamFile) Init(path string, self ...ILogStreamer) ILogStreamer {
	var ref ILogStreamer = s
	if len(self) > 0 {
		ref = self[0]
	}

	s.path = path
	return s.logStream.Init("file", ref)
}

func (s *logStreamFile) GetLogPath() string {
	return s.path
}

func (s *logStreamFile) SetLogPath(path string) *logStreamFile {
	s.path = path
	return s
}

func (s *logStreamFile) Broadcast(level LogLevel, msg string, context ...LogCtx) {
	if s.level == LogNone || s.level > level {
		return
	}

	if s.formatter != nil {
		msg = s.formatter(level, msg, context...)
	}

	fd, e := os.OpenFile(s.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0x644)
	if e == nil {
		_, _ = fmt.Fprintln(fd, msg)
		_ = fd.Close()
	}
}
