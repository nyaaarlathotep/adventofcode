package main

import (
	"fmt"
	"strings"
	"test/util"
)

func fourteenPartOne(input string) int {
	parts := strings.Split(input, "\r\n\r\n")
	polymerTemplate := parts[0]
	rules := util.Get2dString(parts[1], "\r\n", " -> ")
	ruleMap := make(map[string]string)
	for _, v := range rules {
		ruleMap[v[0]] = v[1] + string(v[0][1])
		//strings.Join(strings.Split(v[0], ""), v[1])
	}
	for c := 0; c < 10; c++ {
		temp := string(polymerTemplate[0])
		for i := 1; i < len(polymerTemplate); i++ {
			temp = temp + ruleMap[string(polymerTemplate[i-1])+string(polymerTemplate[i])]
		}
		polymerTemplate = temp
	}
	resMap := make(map[string]int)

	for _, v := range polymerTemplate {
		resMap[string(v)] = resMap[string(v)] + 1
	}
	fmt.Println(resMap)

	return 0
}
