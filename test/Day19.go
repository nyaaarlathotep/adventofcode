package main

import (
	"strconv"
	"strings"
)

func day19(input string) int {
	scannerParts := strings.Split(input, "\n\n")
	scanners := make([][][]int, 0)
	for _, v := range scannerParts {
		scanner := make([][]int, 0)
		lines := strings.Split(v, "\n")
		for i := 1; i < len(lines); i++ {
			line := lines[i]
			location := make([]int, 0)
			nums := strings.Split(line, ",")
			for _, strNum := range nums {
				num, _ := strconv.Atoi(strNum)
				location = append(location, num)
			}
			scanner = append(scanner, location)
		}
		scanners = append(scanners, scanner)
	}

	for i := 0; i < len(scanners)-1; i++ {
		for j := i + 1; j < len(scanners); j++ {

			//beaconA := scanners[0][i]
			//beaconB := scanners[0][j]

		}
	}

	return 0
}
