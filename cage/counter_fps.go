package cage

// FPS @todo doc
type FPS struct {
	Counter
}

var _ ICounter = &FPS{}

// Init @todo doc
func (fps *FPS) Init(self ...ICounter) ICounter {
	var ref ICounter = fps
	if len(self) > 0 {
		ref = self[0]
	}

	return fps.Counter.Init("FPS", ref)
}
