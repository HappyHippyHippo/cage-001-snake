package cage

// INode @todo doc
type INode interface {
	ITraversable
	IPubSub
	IAwakable
	IRunnable
	IVisible
}

// Node @todo doc
type Node struct {
	Traversable
	PubSub
	Awakable
	Runnable
	Visible
}

// Init @todo doc
func (n *Node) Init(id string, self ...INode) INode {
	var ref INode = n
	if len(self) > 0 {
		ref = self[0]
	}

	n.Traversable.Init(id, ref)
	n.PubSub.Init(ref)
	n.Awakable.Init(ref)
	n.Runnable.Init(ref)
	n.Visible.Init(ref)

	return ref
}
