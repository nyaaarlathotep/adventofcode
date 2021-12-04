package main

import (
	"fmt"
	"strconv"
	"strings"
)

func partTwo(commands string) int {
	aim := 0
	forward := 0
	depth := 0
	commandSlice := strings.Split(commands, "\n")

	for _, command := range commandSlice {
		twoWord := strings.Split(command, " ")
		if strings.Compare(twoWord[0], "forward") == 0 {

			temp, _ := strconv.Atoi(twoWord[1])
			fmt.Println("for", temp)

			forward = forward + temp
			depth = depth + aim*temp
		} else if strings.Compare(twoWord[0], "up") == 0 {

			temp, _ := strconv.Atoi(twoWord[1])
			fmt.Println("up", temp)

			aim = aim - temp
		} else if strings.Compare(twoWord[0], "down") == 0 {

			temp, _ := strconv.Atoi(twoWord[1])
			fmt.Println("down", temp)

			aim = aim + temp
		}
	}
	return forward * depth
}

func partOne(commands string) {
	depth := 0
	forward := 0
	commandSlice := strings.Split(commands, "\n")

	for _, command := range commandSlice {
		twoWord := strings.Split(command, " ")
		if strings.Compare(twoWord[0], "forward") == 0 {

			temp, _ := strconv.Atoi(twoWord[1])
			fmt.Println("for", temp)

			forward = forward + temp
		} else if strings.Compare(twoWord[0], "up") == 0 {

			temp, _ := strconv.Atoi(twoWord[1])
			fmt.Println("up", temp)

			depth = depth - temp
		} else if strings.Compare(twoWord[0], "down") == 0 {

			temp, _ := strconv.Atoi(twoWord[1])
			fmt.Println("down", temp)

			depth = depth + temp
		}
	}
}
