package DaySix

import (
	"strconv"
	"strings"
)

func SixPart(input string) int64 {

	strFish := strings.Split(input, ",")
	fish := [...]int64{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, v := range strFish {
		day, _ := strconv.Atoi(v)
		fish[day]++
	}
	//var counter int64
	//pool := fishLive2(&fishes, 200)
	//return int64(len(pool))
	return clever(fish, 256)
}

func clever(fish [9]int64, limit int) int64 {
	newFish := [...]int64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < limit; i++ {
		newFish[8] = fish[0]
		newFish[7] = fish[8]
		newFish[6] = fish[7] + fish[0]
		newFish[5] = fish[6]
		newFish[4] = fish[5]
		newFish[3] = fish[4]
		newFish[2] = fish[3]
		newFish[1] = fish[2]
		newFish[0] = fish[1]
		fish = newFish
		newFish = [...]int64{0, 0, 0, 0, 0, 0, 0, 0, 0}

	}
	var res int64
	for _, v := range fish {
		res = res + v

	}
	return res
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
