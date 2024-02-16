package cage

// LogLevel @todo doc
type LogLevel int

const (
	// LogAll @todo doc
	LogAll LogLevel = iota

	// LogTrace @todo doc
	LogTrace

	// LogDebug @todo doc
	LogDebug

	// LogInfo @todo doc
	LogInfo

	// LogWarning @todo doc
	LogWarning

	// LogError @todo doc
	LogError

	// LogFatal @todo doc
	LogFatal

	// LogNone @todo doc
	LogNone
)

// LogLevelName @todo doc
var LogLevelName = map[LogLevel]string{
	LogAll:     "all",
	LogTrace:   "trace",
	LogDebug:   "debug",
	LogInfo:    "info",
	LogWarning: "warning",
	LogError:   "error",
	LogFatal:   "fatal",
	LogNone:    "none",
}
