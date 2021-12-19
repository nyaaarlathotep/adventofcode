package Day18

import (
	"fmt"
	"strconv"
	"strings"
)

func Eighteen(input string) int {
	lines := strings.Split(input, "\n")
	var root *node
	for i, v := range lines {
		if root != nil {
			printTree(root)
		}
		if i == 0 {
			root = lineToTree(v, nil, true)
			cleanTree(root)
		} else {
			newRoot := NewNode(nil, false)
			rightNode := lineToTree(v, newRoot, false)
			root.father = newRoot
			newRoot.SetLeft(root)
			newRoot.SetRight(rightNode)
			cleanTree(newRoot)
			root = newRoot
		}
	}

	return 0
}

func lineToTree(line string, father *node, isFatherLeft bool) *node {
	s := NewStream(line)
	s.NextChar()
	t := s.GetUntilNext("[", "]")
	s = NewStream(t)
	root := genTree(s, father, isFatherLeft)
	return root
}

func genTree(s *Stream, father *node, isFatherLeft bool) *node {
	n := NewNode(father, isFatherLeft)
	char := s.NextChar()
	if char == "[" {
		leftPart := s.GetUntilNext("[", "]")
		leftStream := NewStream(leftPart)
		n.SetLeft(genTree(leftStream, n, true))
		comma := s.NextChar()
		if comma != "," {
			fmt.Println("?????")
		}
	} else {
		numStr := s.GetUntilNext("~", ",")
		num, e := strconv.Atoi(char + numStr)
		if e != nil {
			fmt.Println(e)
		}
		n.LeftNum = num
	}

	char = s.NextChar()
	if char == "[" {
		rightPart := s.GetUntilNext("[", "]")
		rightStream := NewStream(rightPart)
		n.SetRight(genTree(rightStream, n, false))
	} else {
		numStr := s.GetUntilEnd()
		num, e := strconv.Atoi(char + numStr)
		if e != nil {
			fmt.Println(e)
		}
		n.RightNum = num
	}

	return n
}

func cleanTree(top *node) {

	for true {
		printTree(top)
		exploded := testExplode(top, 1)
		if exploded {
			continue
		}
		splited := testSplit(top)
		if !exploded && !splited {
			break
		}
	}
}

func testExplode(top *node, floor int) bool {
	exploded := false
	if floor == 5 {
		explodeToLeft(top.father, top.LeftNum, top.isFatherLeft)
		explodeToRight(top.father, top.RightNum, top.isFatherLeft)
		if top.isFatherLeft {
			top.father.leftBombed()
		} else {
			top.father.rightBombed()
		}
		exploded = true
	} else {
		floor++
		if top.hasLeftNode {
			tExplode := testExplode(top.left, floor)
			if tExplode {
				return tExplode
			}
		}
		if top.hasRightNode {
			tExplode := testExplode(top.right, floor)
			if tExplode {
				return tExplode
			}
		}
	}
	return exploded
}

func explodeToLeft(n *node, num int, fromLeft bool) {
	if fromLeft {
		if n.father != nil {
			explodeToLeft(n.father, num, n.isFatherLeft)
		}
	} else {
		if !n.hasLeftNode {
			n.LeftNum = n.LeftNum + num
		} else {
			addToRight(n.left, num)
		}
	}
}

func explodeToRight(n *node, num int, fromLeft bool) {
	if !fromLeft {
		if n.father != nil {
			explodeToRight(n.father, num, n.isFatherLeft)
		}
	} else {
		if !n.hasRightNode {
			n.RightNum = n.RightNum + num
		} else {
			addToLeft(n.right, num)
		}
	}
}

func addToLeft(n *node, num int) {
	if n.hasLeftNode {
		addToLeft(n.left, num)
	} else {
		n.LeftNum = n.LeftNum + num
	}
}
func addToRight(n *node, num int) {
	if n.hasRightNode {
		addToRight(n.right, num)
	} else {
		n.RightNum = n.RightNum + num
	}
}

func testSplit(top *node) bool {
	if top.hasLeftNode {
		splited := testSplit(top.left)
		if splited {
			return splited
		}
	} else if top.LeftNum > 9 {
		leftChild := NewNode(top, true)
		leftChild.LeftNum = top.LeftNum / 2
		leftChild.RightNum = top.LeftNum - top.LeftNum/2
		top.SetLeft(leftChild)
		return true
	}

	if top.hasRightNode {
		splited := testSplit(top.right)
		if splited {
			return splited
		}
	} else if top.RightNum > 9 {
		rightChild := NewNode(top, false)
		rightChild.LeftNum = top.RightNum / 2
		rightChild.RightNum = top.RightNum - top.RightNum/2
		top.SetRight(rightChild)
		return true
	}
	return false
}

func printTree(n *node) {
	recPrint(n)
	fmt.Println()
}

func recPrint(n *node) {
	fmt.Printf("[")
	if n.hasLeftNode {
		recPrint(n.left)
	} else {
		fmt.Printf(strconv.Itoa(n.LeftNum))
	}
	fmt.Printf(",")
	if n.hasRightNode {
		recPrint(n.right)
	} else {
		fmt.Printf(strconv.Itoa(n.RightNum))
	}
	fmt.Printf("]")

}
