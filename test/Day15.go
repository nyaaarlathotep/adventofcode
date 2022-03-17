package main

import (
	"fmt"
	"test/util"
)

const INT_MAX = int(^uint(0) >> 1)

func fifteenPartOne(input string) int {
	routeMap := util.TwoDStringToInt(util.Get2dString(input, "\r\n", ""))
	distance := make([][500]int, len(routeMap))

	//updateDistance(&distance, &routeMap, 0, 0)
	distance[0][0] = 0

	return distance[len(routeMap)-1][len(routeMap[0])-1]

}

func fifteenPartTwo(input string) int {

	routeMap := util.TwoDStringToInt(util.Get2dString(input, "\n", ""))
	bigRouteMap := fiveTimeBigger(&routeMap)

	distance := make([][500]int, len(*bigRouteMap))
	InitializeDistance(&distance)
	updateDistance(&distance, bigRouteMap, 0, 0, 0)

	return distance[len(*bigRouteMap)-1][len((*bigRouteMap)[0])-1]

}

func fiveTimeBigger(routeMap *[][]int) *[][]int {
	smallLen := len((*routeMap)[0])
	for i := 0; i < len(*routeMap); i++ {
		tempSlice := make([]int, 0)

		for k := 0; k < 5; k++ {
			for j := 0; j < smallLen; j++ {
				temp := (*routeMap)[i][j] + k
				if temp > 9 {
					temp = temp - 9
				}
				tempSlice = append(tempSlice, temp)
			}
		}
		(*routeMap)[i] = tempSlice
	}
	firstLen := len(*routeMap)
	for k := 1; k < 5; k++ {

		for i := 0; i < firstLen; i++ {
			tempSlice := make([]int, 0)
			for j := 0; j < len((*routeMap)[0]); j++ {
				temp := (*routeMap)[i][j] + k
				if temp > 9 {
					temp = temp - 9
				}
				tempSlice = append(tempSlice, temp)
			}
			*routeMap = append(*routeMap, tempSlice)
		}
	}
	return routeMap
}

func updateDistance(distance *[][500]int, routeMap *[][]int, i int, j int, totalLength int) {
	realRouteMap := *routeMap
	realDistance := *distance
	height := len(realRouteMap)
	width := len(realRouteMap[0])

	realDistance[i][j] = totalLength

	if i == height-1 && j == width-1 {
		return
	}

	if i > 0 {
		temp := totalLength + realRouteMap[i-1][j]
		if temp < realDistance[i-1][j] {
			updateDistance(distance, routeMap, i-1, j, temp)
		}

	}
	if i < height-1 {
		temp := totalLength + realRouteMap[i+1][j]
		if temp < realDistance[i+1][j] {
			updateDistance(distance, routeMap, i+1, j, temp)
		}
	}

	if j > 0 {
		temp := totalLength + realRouteMap[i][j-1]
		if temp < realDistance[i][j-1] {
			updateDistance(distance, routeMap, i, j-1, temp)

		}
	}
	if j < len((*routeMap)[0])-1 {
		temp := totalLength + realRouteMap[i][j+1]
		if temp < realDistance[i][j+1] {
			updateDistance(distance, routeMap, i, j+1, temp)

		}
	}
}

func print2DSlice(s *[][50]int) {
	for _, h := range *s {
		for _, ss := range h {
			fmt.Print(ss)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func InitializeDistance(s *[][500]int) {
	for i := range *s {
		for j := range (*s)[0] {
			(*s)[i][j] = INT_MAX
		}
	}
}
