package cage

// ILoggable @todo doc
type ILoggable interface {
	Signal(level LogLevel, channel, msg string, context ...LogCtx)
	Broadcast(level LogLevel, msg string, context ...LogCtx)
}

