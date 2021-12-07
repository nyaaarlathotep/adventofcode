package main

import (
	"math"
	"strconv"
	"strings"
)

func Seven(input string, partOne bool) int {

	strCrab := strings.Split(input, ",")
	crabSwarm := make(map[int]float64)
	min := 100
	max := 0
	for _, v := range strCrab {
		pos, _ := strconv.Atoi(v)
		crabSwarm[pos] = crabSwarm[pos] + 1
		if pos < min {
			min = pos
		}
		if pos > max {
			max = pos
		}
	}

	minTotal := 10000000000000.0
	var total float64
	for i := min; i < max; i++ {
		total = 0
		for k, v := range crabSwarm {
			if partOne {
				total = total + partOneFuel(k, i)*v
			} else {
				total = total + partTwoFuel(k, i)*v

			}
		}
		if total < minTotal {
			minTotal = total
		}
	}

	return int(minTotal)
}

func partOneFuel(start int, end int) float64 {
	return math.Abs(float64(start - end))
}

func partTwoFuel(start int, end int) float64 {
	return recFuel(math.Abs(float64(start - end)))
}

func recFuel(step float64) float64 {
	if step == 0 {
		return 0
	}
	if step == 1 {
		return 1
	} else {
		return recFuel(step-1) + step
	}
}
