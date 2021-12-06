package DaySix

import (
	"strconv"
	"strings"
)

func SixPart(input string) int64 {

	strFish := strings.Split(input, ",")
	fishes := make([]Fish, 0)

	for _, v := range strFish {
		day, _ := strconv.Atoi(v)
		fish := Fish{day, 0}
		fishes = append(fishes, fish)
	}

	return rec(fishes)
}

func rec(fish []Fish) int64 {
	var counter int64
	for _, singleFish := range fish {
		fishLive(&singleFish, &counter, 256)
	}
	return counter
}

func fishLive(fish *Fish, counter *int64, limit int) {
	if fish.LiveDay() == limit {
		*counter++
		return
	}
	if fish.Live() {
		fishSon := Fish{8, fish.LiveDay()}
		fishLive(&fishSon, counter, limit)
	}
	if fish.LiveDay() == limit {
		*counter = *counter + int64(1)
	} else {
		fishLive(fish, counter, limit)
	}
}
