package Day17

func Seventeen(input string) int {

	area := make([][]int, 0)
	xArea := []int{217, 240}
	yArea := []int{-126, -69}
	area = append(area, xArea)
	area = append(area, yArea)
	areaP := &area

	//maxHeight:=	partOneFlyTest(xArea[1]/2+1,-yArea[0],areaP)
	hitCount := partTwoFlyTest(xArea[1]/2+1, yArea[0], -yArea[0], areaP)
	return hitCount
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

func partTwoFlyTest(vXMax int, vYMin int, vYMax int, areaP *[][]int) int {
	hitCount := 0

	for vX := 0; vX <= vXMax; vX++ {
		for vY := vYMin; vY <= vYMax; vY++ {
			probe := NewProbe(vX, vY)
			for !probe.Out(areaP) {
				probe.fly(areaP)
				if probe.Hit {
					hitCount++
					break
				}
			}
		}
	}
	return hitCount
}
