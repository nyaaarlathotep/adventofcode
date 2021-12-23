package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day22(input string) int {

	reg := regexp.MustCompile(`(-)*\d+`)
	if reg == nil {
		fmt.Println("MustCompile err")
		return 0
	}
	lines := strings.Split(input, "\r\n")
	commands := make([]command, 0)

	for _, line := range lines {
		parts := strings.Split(line, " ")
		isOn := parts[0] == "on"
		res := reg.FindAllStringSubmatch(parts[1], -1)
		x1, _ := strconv.Atoi(res[0][0])
		x2, _ := strconv.Atoi(res[1][0])
		y1, _ := strconv.Atoi(res[2][0])
		y2, _ := strconv.Atoi(res[3][0])
		z1, _ := strconv.Atoi(res[4][0])
		z2, _ := strconv.Atoi(res[5][0])
		commands = append(commands, command{
			x1: x1,
			x2: x2,
			y1: y1,
			y2: y2,
			z1: z1,
			z2: z2,
			on: isOn,
		})
	}

	return Day22partTwo(commands)

}

func Day22PartOne(commands []command) int {
	tempCube := [101][101][101]int{}
	for _, command := range commands {
		for i := command.x1 + 50; i <= command.x2+50; i++ {
			for j := command.y1 + 50; j <= command.y2+50; j++ {
				for k := command.z1 + 50; k <= command.z2+50; k++ {
					if command.on {
						tempCube[i][j][k] = 1
					} else {
						tempCube[i][j][k] = 0
					}
				}
			}
		}
	}

	count := 0
	for i := 0; i < 101; i++ {
		for j := 0; j < 101; j++ {
			for k := 0; k < 101; k++ {
				if tempCube[i][j][k] == 1 {
					count++
				}
			}
		}
	}

	return count
}

func Day22partTwo(commands []command) int {
	//area := [200000]plane{}
	for _, command := range commands {
		for i := command.x1 + 100000; i <= command.x2+100000; i++ {
		}
	}
	var counter int

	for i := 0; i < 200000; i++ {
	}

	return counter
}

type command struct {
	x1, x2, y1, y2, z1, z2 int
	on                     bool
}

type plane struct {

}

func linesAndNewLine(lines *[][2]int, newStart int, newEnd int, on bool) *[][2]int {
	line := [20000]int{}
	for _, l := range *lines {
		for i := l[0]; i <= l[1]; i++ {
			line[i] = 1
		}
	}
	if on {
		for i := newStart; i <= newEnd; i++ {
			line[i] = 1
		}
	} else {
		for i := newStart; i <= newEnd; i++ {
			line[i] = 0
		}
	}

	newStart = -1
	var newLines [][2]int
	for i := range line {
		if line[i] == 0 {
			if newStart == -1 {
				continue
			} else {
				newLines = append(newLines, [2]int{newStart, i})
				newStart = -1
				continue
			}
		}
		if newStart == -1 {
			newStart = i
		} else {
			continue
		}
	}
	return &newLines
}
