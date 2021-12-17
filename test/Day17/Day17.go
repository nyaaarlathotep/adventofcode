package Day17

import (
	"fmt"
	"strconv"
	"strings"
	"test/util"
)

func Seventeen(input string) int {

	for i:=0;i<5;i++{
		input= strings.ReplaceAll(input,"  "," ")
	}

	nums := util.Get2dString(input, "\r\n", " ")
	res := make([]string, 0)
	for _, v := range nums {
		for _, single := range v {
			res = append(res, single)
		}
	}

	area := make([][]int, 0)
	xArea := []int{20, 30}
	yArea := []int{-10, -5}
	area = append(area, xArea)
	area = append(area, yArea)
	areaP := &area

	//maxHeight:=	partOneFlyTest(xArea[1]/2+1,-yArea[0],areaP)
	hitCount := partTwoFlyTest(xArea[1], yArea[0], -yArea[0], areaP)
	for _, v := range res {
		match := false
		for _, vv := range hitCount {
			if v == vv {
				match = true
				break
			}
		}
		if !match {
			fmt.Println(v)
		}
	}
	return 0
}

func partOneFlyTest(vXMax int, vYMax int, areaP *[][]int) int {
	maxHeight := -1

	for vX := 0; vX <= vXMax; vX++ {
		for vY := 0; vY <= vYMax; vY++ {
			probe := NewProbe(vX, vY)
			for !probe.Out(areaP) {
				probe.fly(areaP)
				if probe.Hit && probe.MaxHeight > maxHeight {
					maxHeight = probe.MaxHeight
					break
				}
			}
		}
	}
	return maxHeight
}

func partTwoFlyTest(vXMax int, vYMin int, vYMax int, areaP *[][]int) []string {
	hitCount := 0
	res := make([]string, 0)

	//7,-1
	//6,0
	for vX := 0; vX <= vXMax; vX++ {
		for vY := vYMin; vY <= vYMax; vY++ {
			probe := NewProbe(vX, vY)
			for !probe.Out(areaP) {
				probe.fly(areaP)
				if probe.Hit {
					vvY := strconv.Itoa(vX)
					vvX := strconv.Itoa(vY)
					res = append(res, vvY+","+vvX)
					hitCount++
					break
				}
			}
		}
	}
	fmt.Println(res)
	fmt.Println("--------------------")
	return res
}
