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
	rec2(fishes, 80, &counter)
	return counter
}

func rec2(fishPool []Fish, limit int, counter *int64) {
	if fishPool[0].LiveDay() >= limit {
		*counter = *counter + int64(len(fishPool))
		return
	}
	duration := limit / 4
	for _, fish := range fishPool {
		singlePool := []Fish{fish}
		afterPool := fishLive2(&singlePool, duration)
		newSinglePool := []Fish{afterPool[0]}
		rec2(newSinglePool, limit, counter)
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
