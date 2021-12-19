package Day18

type node struct {
	hasLeftNode  bool
	hasRightNode bool
	left         *node
	right        *node
	LeftNum      int
	RightNum     int
	father       *node
	isFatherLeft bool
}

func NewNode(fa *node, isFaLeft bool) *node {
	n := &node{
		hasLeftNode:  false,
		hasRightNode: false,
		left:         nil,
		right:        nil,
		LeftNum:      0,
		RightNum:     0,
		father:       fa,
		isFatherLeft: isFaLeft,
	}
	return n
}

func (n *node) SetRight(r *node) {
	r.isFatherLeft = false
	n.right = r
	n.hasRightNode = true
	n.RightNum = 0
}

func (n *node) SetLeft(l *node) {
	l.isFatherLeft = true
	n.left = l
	n.hasLeftNode = true
	n.LeftNum = 0
}

func (n *node) leftBombed() {
	n.left = nil
	n.LeftNum = 0
	n.hasLeftNode = false
}

func (n *node) rightBombed() {
	n.right = nil
	n.RightNum = 0
	n.hasRightNode = false
}
