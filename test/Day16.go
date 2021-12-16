package main

import (
	"fmt"
	"strconv"
	"test/DataStruct"
)

func sixteenPartOne(input string) int {

	//i, _ := strconv.ParseInt("D2", 16, 32)
	//bits := strconv.FormatInt(i, 2)
	//
	//fmt.Println(i)
	//fmt.Println(bits)
	pipe := DataStruct.NewPipe(input)
	versions := make([]int, 0)
	fmt.Println(handlePackage(pipe, &versions))
	fmt.Println(versions)

	return 0
}

func handlePackage(pipe *DataStruct.Pipe, versions *[]int)  int {
	version, _ := strconv.ParseInt(pipe.GetChars(3), 2, 32)
	*versions = append(*versions, int(version))
	strType, _ := strconv.ParseInt(pipe.GetChars(3), 2, 32)
	if strType == 4 {
		resNum := handleNum(pipe)
		return resNum
	} else {
		//lengthType, _ := strconv.ParseInt(str[6:7], 2, 32)
		//if lengthType == 0 {
		//	length, _ := strconv.ParseInt(str[7:22], 2, 32)
		//	index, resNum := handleTwoPact(str[22:18+length], versions)
		//	return index, resNum
		//
		//} else if lengthType == 1 {
		//	pkgNum, _ := strconv.ParseInt(str[6:17], 2, 32)
		//	index, resNum := handleMultiPact(str[17:], versions, int(pkgNum))
		//	return index, resNum
		//} else {
		//	fmt.Println("????")
		//}
	}
	return 0
}

func handleNum(pipe *DataStruct.Pipe) int {
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
	res, _ := strconv.ParseInt(resNum, 2, 32)
	return int(res)
}

func handleMultiPact(nums string, versions *[]int, count int) (int, int) {

	index := 0
	//counter := 0
	//for counter != count {
	//	_, index = handlePackage(nums[index:], versions)
	//	//resss, _ := strconv.ParseInt(resNum, 2, 32)
	//	counter++
	//}
	return index, 0
}

func handleTwoPact(nums string, versions *[]int) (int, int) {
	//counter := 0
	var index int
	//for counter != 2 {
	//	_, index = handlePackage(nums[index:], versions)
	//	//resss, _ := strconv.ParseInt(resNum, 2, 32)
	//	counter++
	//}
	return index, 0
}
