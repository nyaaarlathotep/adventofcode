package main

import (
	"test/util"
)

func Day25(input string) int {

	rawHerdMap := util.Get2dString(input, "\r\n", "")
	seaCuMap := make([][]int, 0)

	for _, line := range rawHerdMap {
		seaLine := make([]int, 0)
		for _, v := range line {
			if v == "." {
				seaLine = append(seaLine, 0)
			} else if v == ">" {
				seaLine = append(seaLine, 1)
			} else if v == "v" {
				seaLine = append(seaLine, 2)
			}
		}
		seaCuMap = append(seaCuMap, seaLine)

	}

	counter := 0
	for true {
		counter++
		nextSeaCuMap := make([][]int, 0)
		for i := 0; i < len(seaCuMap); i++ {
			nextSeaCuMap = append(nextSeaCuMap, make([]int, len(seaCuMap[0])))
		}
		moved := false
		for i := range seaCuMap {
			for j := range seaCuMap[i] {
				if seaCuMap[i][j] == 1 {
					thisMoved := rightMove(&seaCuMap, &nextSeaCuMap, i, j)
					if !moved && thisMoved {
						moved = true
					}
				}
			}
		}

		//fmt.Println()
		//for _, v := range nextSeaCuMap {
		//	fmt.Println(v)
		//}

		for i := range seaCuMap {
			for j := range seaCuMap[i] {
				if seaCuMap[i][j] == 1 {
					seaCuMap[i][j] = 0

				}
			}
		}
		for i := range nextSeaCuMap {
			for j := range nextSeaCuMap[i] {
				if nextSeaCuMap[i][j] == 1 {
					seaCuMap[i][j] = 1
				}
			}
		}

		//fmt.Println()
		//for _, v := range seaCuMap {
		//	fmt.Println(v)
		//}

		for i := range seaCuMap {
			for j := range seaCuMap[i] {
				if seaCuMap[i][j] == 2 {
					thisMoved := downMove(&seaCuMap, &nextSeaCuMap, i, j)
					if !moved && thisMoved {
						moved = true
					}
				}
			}
		}

		seaCuMap = nextSeaCuMap

		if !moved {
			break
		}
	}

	return counter

}

func rightMove(secCuMap *[][]int, nextSecCuMap *[][]int, i int, j int) bool {
	moved := false
	nextJ := j + 1
	if nextJ >= len((*secCuMap)[i]) {
		nextJ = 0
	}
	if (*secCuMap)[i][nextJ] == 0 {
		(*nextSecCuMap)[i][nextJ] = 1
		moved = true
	} else {
		(*nextSecCuMap)[i][j] = 1

	}
	return moved
}

func downMove(secCuMap *[][]int, nextSecCuMap *[][]int, i int, j int) bool {
	moved := false

	nextI := i + 1
	if nextI >= len(*secCuMap) {
		nextI = 0
	}
	if (*secCuMap)[nextI][j] == 0 {
		(*nextSecCuMap)[nextI][j] = 2
		moved = true
	} else {
		(*nextSecCuMap)[i][j] = 2

	}
	return moved

}
