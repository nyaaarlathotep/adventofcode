package DaySix

type Fish struct {
	birthDay int
	liveDay  int
}

func (fish *Fish) LiveDay() int {
	return fish.liveDay
}

func (fish *Fish) Live() bool {
	fish.liveDay++
	if fish.birthDay == 0 {
		(*fish).birthDay = 6
		return true
	} else {
		fish.birthDay = fish.birthDay - 1
		return false
	}
}
