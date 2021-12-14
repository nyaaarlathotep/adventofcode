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

func fourteenPartTwo(input string) int {
	parts := strings.Split(input, "\n\n")
	polymerTemplate := parts[0]
	rules := util.Get2dString(parts[1], "\n", " -> ")
	ruleMap := make(map[string][]string)
	res := make(map[string]int)
	for _, v := range rules {
		ruleMap[v[0]] = []string{string(v[0][0]) + v[1], v[1] + string(v[0][1])}
	}
	for i := 1; i < len(polymerTemplate); i++ {
		charsNow := string(polymerTemplate[i-1]) + string(polymerTemplate[i])
		res[charsNow] = res[charsNow] + 1
	}

	for c := 0; c < 40; c++ {
		temp := make(map[string]int)
		for k := range res {
			temp[ruleMap[k][0]] = res[k] + temp[ruleMap[k][0]]
			temp[ruleMap[k][1]] = res[k] + temp[ruleMap[k][1]]
		}
		res = temp
	}

	resCount := make(map[string]int)
	for k, v := range res {
		resCount[string(k[0])] = resCount[string(k[0])] + v
		resCount[string(k[1])] = resCount[string(k[1])] + v
	}
	resCount[string(polymerTemplate[0])] = resCount[string(polymerTemplate[0])] + 1
	resCount[string(polymerTemplate[len(polymerTemplate)-1])] = resCount[string(polymerTemplate[len(polymerTemplate)-1])] + 1
	for k := range resCount {
		resCount[k] = resCount[k] / 2
	}
	max := 0
	min := 984258259007
	for _, v := range resCount {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println(resCount)
	fmt.Println(max - min)
	return 0
}
