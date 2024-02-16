package cage

const (
	// IAwakable signals

	// SignalAwake @todo doc
	SignalAwake = iota + 1
	// SignalSleep @todo doc
	SignalSleep
	// SignalClose @todo doc
	SignalClose

	// IRunnable signals

	// SignalRun @todo doc
	SignalRun
	// SignalPause @todo doc
	SignalPause
	// SignalUpdate @todo doc
	SignalUpdate

	// IVisible signals

	// SignalShow @todo doc
	SignalShow
	// SignalHide @todo doc
	SignalHide
	// SignalRender @todo doc
	SignalRender

  // Log signals

  // SignalLogChannel @todo doc
  SignalLogChannel

  // SignalLogBroadcast @todo doc
  SignalLogBroadcast

	// SignalImageLoaded @todo doc
	SignalImageLoaded
	// SignalImageLoadError @todo doc
	SignalImageLoadError
	// SignalTimerStarted @todo doc
	SignalTimerStarted
	// SignalTimerEnded @todo doc
	SignalTimerEnded
	// SignalTickerStarted @todo doc
	SignalTickerStarted
	// SignalTickerStep @todo doc
	SignalTickerStep
	// SignalTickerEnded @todo doc
	SignalTickerEnded
)
