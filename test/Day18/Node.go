package Day18

type node struct {
	hasLeftNodeLeft bool
	hasRightNode    bool
	left            *node
	right           *node
	LeftNum         int
	RightNum        int
}

func NewNode() *node {
	n := &node{
		hasLeftNodeLeft: false,
		hasRightNode:    false,
		left:            nil,
		right:           nil,
		LeftNum:         0,
		RightNum:        0,
	}
	return n
}

func (n *node) SetRight(r *node) {
	n.right = r
	n.hasRightNode = true
}

func (n *node) SetLeft(l *node) {
	n.left = l
	n.hasLeftNodeLeft = true
}
