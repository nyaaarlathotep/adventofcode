package main

import (
	"fmt"
	"strconv"
	"test/DataStruct"
	"test/FuncPkg"
)

var funcMap = make(map[int]func(nums []int64) int64)

func sixteenPartOne(input string) int64 {

	funcMap[0] = FuncPkg.Adder
	funcMap[1] = FuncPkg.Multiplier
	funcMap[2] = FuncPkg.Min
	funcMap[3] = FuncPkg.Max
	funcMap[5] = FuncPkg.GreaterThan
	funcMap[6] = FuncPkg.LessThan
	funcMap[7] = FuncPkg.Equal
	pipe := DataStruct.NewPipe(input)
	versions := make([]int, 0)

	return handlePackage(pipe, &versions)
}

func handlePackage(pipe *DataStruct.Pipe, versions *[]int) int64 {
	version, _ := strconv.ParseInt(pipe.GetChars(3), 2, 64)
	*versions = append(*versions, int(version))
	strType, _ := strconv.ParseInt(pipe.GetChars(3), 2, 64)
	if strType == 4 {
		resNum := handleNum(pipe)
		return resNum
	} else {
		lengthType := pipe.GetChars(1)
		if lengthType == "0" {
			length, _ := strconv.ParseInt(pipe.GetChars(15), 2, 64)
			resNum := handleLengthPact(pipe, versions, int(length), funcMap[int(strType)])
			return resNum

		} else if lengthType == "1" {
			pkgNum, _ := strconv.ParseInt(pipe.GetChars(11), 2, 64)
			resNum := handleMultiPact(pipe, versions, int(pkgNum), funcMap[int(strType)])
			return resNum
		} else {
			fmt.Println("????")
		}
	}
	return 0
}

// 12419742984
// 12419967741
func handleNum(pipe *DataStruct.Pipe) int64 {
	resNum := ""
	temp := ""
	for true {
		if pipe.NextChar() != "1" {
			temp = pipe.GetChars(4)
			resNum = resNum + temp
			break
		}
		temp = pipe.GetChars(4)
		resNum = resNum + temp
	}
	res, _ := strconv.ParseInt(resNum, 2, 64)
	return res
}

func handleMultiPact(pipe *DataStruct.Pipe, versions *[]int, count int, handler func(nums []int64) int64) int64 {
	counter := 0
	nums := make([]int64, 0)
	for counter != count {
		temp := handlePackage(pipe, versions)
		nums = append(nums, temp)
		counter++
	}
	return handler(nums)

}

func handleLengthPact(pipe *DataStruct.Pipe, versions *[]int, length int, handler func(nums []int64) int64) int64 {
	tempStr := pipe.GetChars(length)
	pp := DataStruct.CacheNewPipe(tempStr)
	nums := make([]int64, 0)
	for !pp.IsEnd() {
		temp := handlePackage(pp, versions)
		nums = append(nums, temp)
	}
	return handler(nums)
}
