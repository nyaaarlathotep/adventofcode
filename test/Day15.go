package main

import (
	"fmt"
	"test/util"
)

func fifteenPartOne(input string) int {
	routeMap := util.TwoDStringToInt(util.Get2dString(input, "\r\n", ""))
	distance := make([][500]int, len(routeMap))

	updateDistance(&distance, &routeMap, 0, 0)
	distance[0][0] = 0

	return distance[len(routeMap)-1][len(routeMap[0])-1]

}

func fifteenPartTwo(input string) int {
	routeMap := util.TwoDStringToInt(util.Get2dString(input, "\r\n", ""))
	bigRouteMap := fiveTimeBigger(&routeMap)
	fmt.Println(len(*bigRouteMap))
	fmt.Println(len((*bigRouteMap)[0]))

	distance := make([][500]int, len(*bigRouteMap))

	updateDistance(&distance, bigRouteMap, 0, 0)
	distance[0][0] = 0

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

func updateDistance(distance *[][500]int, routeMap *[][]int, i int, j int) {
	if i==len(*routeMap)-1 && j==len((*routeMap)[0])-1 {
		return
	}

	if (*distance)[i][j] == 0 {
		(*distance)[i+1][j] = (*routeMap)[i+1][j]
		(*distance)[i][j+1] = (*routeMap)[i][j+1]
		updateDistance(distance, routeMap, i+1, j)
		//updateDistance(distance,routeMap,i,j+1)
	} else {
		if i == 0 && j == 0 {
			return
		}
		if i > 0 {
			temp := (*distance)[i][j] + (*routeMap)[i-1][j]
			if (*distance)[i-1][j] == 0 || (*distance)[i-1][j] > temp {
				(*distance)[i-1][j] = temp
				updateDistance(distance, routeMap, i-1, j)
			}
		}
		if i < len(*routeMap)-1 {
			temp := (*distance)[i][j] + (*routeMap)[i+1][j]
			if (*distance)[i+1][j] == 0 || (*distance)[i+1][j] > temp {
				(*distance)[i+1][j] = temp
				updateDistance(distance, routeMap, i+1, j)
			}
		}

		if j > 0 {
			temp := (*distance)[i][j] + (*routeMap)[i][j-1]
			if (*distance)[i][j-1] == 0 || (*distance)[i][j-1] > temp {
				(*distance)[i][j-1] = temp
				updateDistance(distance, routeMap, i, j-1)

			}
		}
		if j < len((*routeMap)[0])-1 {
			temp := (*distance)[i][j] + (*routeMap)[i][j+1]
			if (*distance)[i][j+1] == 0 || (*distance)[i][j+1] > temp {
				(*distance)[i][j+1] = temp
				updateDistance(distance, routeMap, i, j+1)

			}
		}

	}
}
