package Day18

import (
	"fmt"
	"strconv"
)

func Eighteen(input string) int {
	s := NewStream(input)
	s.NextChar()
	t := s.GetUntilNext("[", "]")
	s = NewStream(t)

	n := genTree(s)
	fmt.Println(n)

	return 0
}

func genTree(s *Stream) *node {
	n := NewNode()
	char := s.NextChar()
	if char == "[" {
		leftPart := s.GetUntilNext("[", "]")
		leftStream := NewStream(leftPart)
		n.SetLeft(genTree(leftStream))
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
		n.SetRight(genTree(rightStream))
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
