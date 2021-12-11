package main

import (
	"test/util"
)

func elevPartOne(input string) int {
	fish := util.TwoDStringToInt(util.Get2dString(input, "\n", ""))
	count := 0
	for i := 0; i < 100; i++ {
		energyIncrease(&fish)
		for in := range fish {
			for jn := range (fish)[0] {
				if fish[in][jn] == 10 {

					flash(&fish, in, jn)
					//fmt.Println(in, "==================", jn)
					//for _, line := range fish {
					//	for _, v := range line {
					//		if v > 9 {
					//			fmt.Printf("! ")
					//		} else {
					//			fmt.Printf(strconv.Itoa(v) + " ")
					//		}
					//	}
					//	fmt.Println()
					//}
				}
			}
		}
		checkFlash(&fish, &count)
	}

	return count
}

func flash(fish *[][]int, i int, j int) {
	(*fish)[i][j]++
	//if i ==0 && j == 2 {
	//	fmt.Println(len((*fish)[0])-1)
	//}
	if i > 0 {
		(*fish)[i-1][j]++
		if (*fish)[i-1][j] == 10 || (*fish)[i-1][j] == 11 {
			flash(fish, i-1, j)
		}
		if j > 0 {
			(*fish)[i-1][j-1]++
			if (*fish)[i-1][j-1] == 10 || (*fish)[i-1][j-1] == 11 {
				flash(fish, i-1, j-1)
			}

		}
		if j < len((*fish)[0])-1 {
			(*fish)[i-1][j+1]++
			if (*fish)[i-1][j+1] == 10 || (*fish)[i-1][j+1] == 11 {

				flash(fish, i-1, j+1)
			}
		}
	}
	if i < len(*fish)-1 {
		(*fish)[i+1][j]++
		if (*fish)[i+1][j] == 10 || (*fish)[i+1][j] == 11 {

			flash(fish, i+1, j)
		}
		if j > 0 {
			(*fish)[i+1][j-1]++
			if (*fish)[i+1][j-1] == 10 || (*fish)[i+1][j-1] == 11 {

				flash(fish, i+1, j-1)
			}
		}
		if j < len((*fish)[0])-1 {
			(*fish)[i+1][j+1]++
			if (*fish)[i+1][j+1] == 10 || (*fish)[i+1][j+1] == 11 {

				flash(fish, i+1, j+1)
			}
		}
	}

	if j > 0 {
		(*fish)[i][j-1]++

		if (*fish)[i][j-1] == 10 || (*fish)[i][j-1] == 11 {

			flash(fish, i, j-1)
		}
	}
	if j < len((*fish)[0])-1 {
		(*fish)[i][j+1]++

		if (*fish)[i][j+1] == 10 || (*fish)[i][j+1] == 11 {

			flash(fish, i, j+1)
		}
	}
}

func energyIncrease(fish *[][]int) {
	for i := range *fish {
		for j := range (*fish)[0] {
			(*fish)[i][j]++
		}
	}
}

func checkFlash(fish *[][]int, count *int) {
	for i := range *fish {
		for j := range (*fish)[0] {
			if (*fish)[i][j] > 9 {
				(*fish)[i][j] = 0
				*count++
			}
		}
	}
}
