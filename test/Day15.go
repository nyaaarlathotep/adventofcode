package main

import (
	"test/util"
)

func fifteenPartOne(input string) int {
	routeMap := util.TwoDStringToInt(util.Get2dString(input, "\r\n", ""))
	distance := make([][100]int, len(routeMap))

	updateDistance(&distance, &routeMap,0, 0)
	distance[0][0] = 0

	return distance[len(routeMap)-1][len(routeMap[0])-1]

}

func updateDistance(distance *[][100]int, routeMap *[][]int, i int, j int) {
	if (*distance)[i][j] == 0 {
		(*distance)[i+1][j] = (*routeMap)[i+1][j]
		updateDistance(distance,routeMap,i+1,j)
		(*distance)[i][j+1] = (*routeMap)[i][j+1]
		updateDistance(distance,routeMap,i,j+1)
	} else {
		if i==0 && j==0{
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
