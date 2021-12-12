package main

import (
	"fmt"
	"strings"
	"test/DataStruct"
	"test/util"
)

func twelvePartOne(input string) int {
	ways := util.Get2dString(input, "\n", "-")
	spots := make([]*DataStruct.LinkTable, 0)
	for _, way := range ways {
		exist, beginSpot := findSpot(way[0], spots)
		if !exist {
			beginSpot = DataStruct.NewLinkTable(way[0])
			spots = append(spots, beginSpot)
		}

		exist, endSpot := findSpot(way[1], spots)
		if !exist {
			endSpot = DataStruct.NewLinkTable(way[1])
			spots = append(spots, endSpot)
		}
		beginSpot.AddWay(endSpot)
		endSpot.AddWay(beginSpot)

	}
	for _, v := range spots {
		fmt.Println(v)
	}
	_, start := findSpot("start", spots)
	paths := make([][]string, 0)
	findWay(start, &paths, make([]string, 0))
	for _, v := range paths {
		fmt.Println(v)
	}
	return len(paths)
}

func findSpot(name string, linkTables []*DataStruct.LinkTable) (bool, *DataStruct.LinkTable) {
	for _, lt := range linkTables {
		if lt.Name == name {
			return true, lt
		}
	}
	return false, nil
}

func findWay(t *DataStruct.LinkTable, paths *[][]string, path []string) {
	path = append(path, t.Name)

	if strings.Compare(t.Name, "end") == 0 {
		*paths = append(*paths, util.CopyStringSlice(path))
		return
	}

	for _, to := range t.Ptr {
		if to.IsBig() {
			findWay(to, paths, path)
		} else {
			hasBeen := false
			for _, pastSpot := range path {
				if to.Name == pastSpot {
					hasBeen = true
					break
				}
			}
			if !hasBeen {
				findWay(to, paths, path)
			}
		}

	}
	return
}
