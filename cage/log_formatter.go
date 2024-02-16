package cage

import (
	"fmt"
	"strings"
	"time"
)

const (
	// LogCtxMsgArgsField @todo doc
	LogCtxMsgArgsField = "args"
)

// LogCtx @todo doc
type LogCtx map[string]interface{}

// LogCtxMsgArgs @todo doc
func LogCtxMsgArgs(args ...any) LogCtx {
	return LogCtx{LogCtxMsgArgsField: args}
}

// LogFormatter @todo doc
type LogFormatter func(level LogLevel, mdg string, context ...LogCtx) string

func logFormatter(level LogLevel, msg string, context ...LogCtx) string {
	data := LogCtx{}
	for _, ctx := range context {
		for k, v := range ctx {
			data[k] = v
		}
	}
	if ctx, ok := data[LogCtxMsgArgsField].([]any); ok {
		msg = fmt.Sprintf(msg, ctx...)
	}
	return fmt.Sprintf(
		"%7s [%s] %s",
		strings.ToUpper(LogLevelName[level]),
		time.Now().Format("2006-01-02T15:04:05.000-0700"),
		msg,
	)
}
