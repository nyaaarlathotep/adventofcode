package Day17

func Seventeen(input string) int {

	area := make([][]int, 0)
	xArea := []int{217, 240}
	yArea := []int{-126, -69}
	area = append(area, xArea)
	area = append(area, yArea)
	areaP := &area
	maxHeight := -1

	for vX := 0; vX <= xArea[1]/2+1; vX++ {
		for vY := 0; vY <= -yArea[0]; vY++ {
			probe := NewProbe(vX, vY)
			for !probe.Out(areaP) {
				probe.fly(areaP)
				if probe.Hit && probe.MaxHeight > maxHeight {
					maxHeight = probe.MaxHeight
				}
			}
		}
	}

	return maxHeight
}
