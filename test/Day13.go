package main

import (
	"fmt"
	"strconv"
	"strings"
	"test/util"
)

func thirteenPartOne(input string) int {
	graph := [1500][1500]int{}
	parts := strings.Split(input, "\r\n\r\n")
	dots := util.Get2dString(parts[0], "\r\n", ",")
	instructs := make([][]string, 0)
	for _, v := range strings.Split(parts[1], "\r\n") {
		two := strings.Split(strings.Split(v, "fold along ")[1], "=")
		instructs = append(instructs, two)
	}
	for _, dot := range dots {
		y, _ := strconv.Atoi(dot[0])
		x, _ := strconv.Atoi(dot[1])
		graph[x][y] = 1
	}
	limitX := 1500
	limitY := 1500

	for _, inst := range instructs {
		if inst[0] == "x" {
			pos, _ := strconv.Atoi(inst[1])
			limitY = pos
			for i := 1; pos-i >= 0 && pos+i < len(graph[0]); i++ {
				for j := range graph {
					graph[j][pos-i] = graph[j][pos-i] + graph[j][pos+i]
				}
			}
		} else {
			pos, _ := strconv.Atoi(inst[1])
			limitX = pos
			for i := 1; pos-i >= 0 && pos+i < len(graph); i++ {
				for j := range graph[0] {
					graph[pos-i][j] = graph[pos-i][j] + graph[pos+i][j]
				}
			}
		}
	}
	count := 0
	for i := 0; i < limitX; i++ {
		for j := 0; j < limitY; j++ {
			if (graph[i][j]) > 0 {
				graph[i][j] = 1
			}
			fmt.Printf(strconv.Itoa(graph[i][j]) + " ")
			if graph[i][j] > 0 {
				count++
			}
		}
		fmt.Printf("\n")

	}

	return count
}
