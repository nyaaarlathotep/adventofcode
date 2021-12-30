package Day24

import (
	"fmt"
	"strings"
)

func Day24(input string) int {
	//commands := strings.Split(input, "\r\n")
	//for _, command := range commands {
	//	parse(command)
	//}
	inp := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	stop := 0
	c := make([]int, len(inp))
	copy(c, inp)
	fmt.Printf("run %v\n", c)
	var cont = true
	for cont {
		z := execute(c)
		if z==0{
			fmt.Printf("bingo inp=%v\n", c)
			return 0
		}
		cont = decrement(c, len(c)-1, stop)

	}
	return 0
}

func increment(inp []int, index int, stop int) bool {

	if inp[index] == 9 {
		inp[index] = 1
		return increment(inp, index-1, stop)
	} else {
		inp[index]++
	}
	if index < stop {
		return false
	}
	return true
}

func decrement(inp []int, index int, stop int) bool {

	if inp[index] == 1 {
		inp[index] = 9
		return decrement(inp, index-1, stop)
	} else {
		inp[index]--
	}
	if index < stop {
		return false
	}
	return true
}


func parse(command string) {
	commands := strings.Split(command, " ")
	if len(commands) == 2 {
		if commands[0] != "inp" {
			fmt.Println("???")
		}
		fmt.Println(commands[1] + " = " + "inp[index]")
		fmt.Println("index ++")
	} else {
		switch commands[0] {
		case "add":
			fmt.Println(commands[1] + " = " + commands[1] + " + " + commands[2])
		case "mul":
			fmt.Println(commands[1] + " = " + commands[1] + " * " + commands[2])
		case "div":
			fmt.Println(commands[1] + " = " + commands[1] + " / " + commands[2])
		case "mod":
			fmt.Println(commands[1] + " = " + commands[1] + " % " + commands[2])
		case "eql":
			fmt.Println("if " + commands[1] + " == " + commands[2] + "{")
			fmt.Println(commands[1] + " = " + "1")
			fmt.Println("} else {")
			fmt.Println(commands[1] + " = " + "0")
			fmt.Println("}")
		}
	}
}

func valueOfIt(str string, aluMap *map[string]string) string {
	if string(str[0]) == "x" || string(str[0]) == "y" || string(str[0]) == "z" || string(str[0]) == "w" {
		if (*aluMap)[str] == "" {
			return "0"
		}
		return (*aluMap)[str]
	}
	return str
}

func getMapEle(key string, aluMap *map[string]string) string {
	if (*aluMap)[key] == "" {
		return "0"
	} else {
		return (*aluMap)[key]
	}
}
