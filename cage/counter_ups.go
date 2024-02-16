package cage

// UPS @todo doc
type UPS struct {
	Counter
}

var _ ICounter = &UPS{}

// Init @todo doc
func (ups *UPS) Init(self ...ICounter) ICounter {
	var ref ICounter = ups
	if len(self) > 0 {
		ref = self[0]
	}

	return ups.Counter.Init("UPS", ref)
}
