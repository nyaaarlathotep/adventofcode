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
	var counter int64
	rec2(fishes, &counter, true)

	//pool := fishLive2(&fishes, 200)
	//return int64(len(pool))
	return counter
}

func rec2(fishPool []Fish, counter *int64, first bool) {
	if fishPool[0].LiveDay() >= 256 {
		*counter = *counter + int64(len(fishPool))
		return
	}
	for _, fish := range fishPool {
		singlePool := []Fish{fish}
		var afterPool []Fish
		if first {
			afterPool = fishLive2(&singlePool, 200)

		} else {
			afterPool = fishLive2(&singlePool, 56)
		}
		for _, afterPoolFish := range afterPool {
			newSinglePool := []Fish{afterPoolFish}
			rec2(newSinglePool, counter, false)
		}
	}

}

func fishLive2(fishPool *[]Fish, duration int) []Fish {

	for i := 0; i < duration; i++ {
		newFishCounter := 0
		for i := range *fishPool {
			if (*fishPool)[i].Live() {
				newFishCounter++
			}
		}
		for i := 0; i < newFishCounter; i++ {
			newFish := Fish{8, (*fishPool)[i].LiveDay()}
			*fishPool = append(*fishPool, newFish)
		}
	}
	return *fishPool

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
