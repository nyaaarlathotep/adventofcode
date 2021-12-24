package main

import (
	"fmt"
	"strconv"
	"strings"
	"test/DataStruct"
	"unicode"
)

func day24(input string) int {
	var aluMap map[string]int
	aluMap = make(map[string]int)
	commands := strings.Split(input, "\r\n")
	max := 99999999999999
	for true {
		p := DataStruct.CacheNewPipe(strconv.Itoa(max))
		for _, command := range commands {
			execute(command, p, &aluMap)
		}
		if aluMap["z"] == 0 {
			break
		} else {
			max--
		}
	}

	return 0
}

func execute(command string, monad *DataStruct.Pipe, aluMap *map[string]int) {
	commands := strings.Split(command, " ")
	if len(commands) == 2 {
		if commands[0] != "inp" {
			fmt.Println("???")
		}
		next, _ := strconv.Atoi(monad.NextChar())
		(*aluMap)[commands[1]] = next
	} else {
		switch commands[0] {
		case "add":
			(*aluMap)[commands[1]] = valueOfIt(commands[2], aluMap) + (*aluMap)[commands[1]]
		case "mul":
			(*aluMap)[commands[1]] = valueOfIt(commands[2], aluMap) * ((*aluMap)[commands[1]])
		case "div":
			(*aluMap)[commands[1]] = (*aluMap)[commands[1]] / valueOfIt(commands[2], aluMap)
		case "mod":
			(*aluMap)[commands[1]] = (*aluMap)[commands[1]] % valueOfIt(commands[2], aluMap)
		case "eql":
			if (*aluMap)[commands[1]] == valueOfIt(commands[2], aluMap) {
				(*aluMap)[commands[1]] = 1
			} else {
				(*aluMap)[commands[1]] = 0
			}
		}
	}
}
func valueOfIt(str string, aluMap *map[string]int) int {
	if unicode.IsNumber(rune(str[0])) {
		temp, _ := strconv.Atoi(str)
		return temp
	} else {
		return (*aluMap)[str]
	}
}
