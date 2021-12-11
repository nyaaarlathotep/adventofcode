package util

import (
	"bytes"
	"strconv"
	"strings"
)

func GetDecimalFromBinaryInt(binarySlice []int) int {
	var buffer bytes.Buffer
	for _, v := range binarySlice {
		buffer.WriteString(strconv.Itoa(v))
	}
	res, _ := strconv.ParseInt(buffer.String(), 2, 64)
	return int(res)
}

func GetDecimalFromBinaryString(binarySlice []string) int {
	binary := strings.Join(binarySlice, "")
	res, _ := strconv.ParseInt(binary, 2, 64)
	return int(res)
}

func GetSliceFromString(str string) []string {
	return strings.Split(str, "")
}

func StringAt(str string, index int) string {
	tmp := strings.Split(str, "")
	return tmp[index]
}

func Get2dString(input string, sep1 string, sep2 string) [][]string {
	input = strings.Replace(input, "  ", " ", -1)
	lines := strings.Split(input, sep1)
	tDString := make([][]string, 0)
	for _, v := range lines {
		v = strings.Trim(v, " ")
		line := strings.Split(v, sep2)
		tDString = append(tDString, line)
	}
	return tDString
}

func GetStringSlice(input string, sep string) []string {
	return strings.Split(input, sep)
}

func TwoDStringToInt(s [][]string) [][]int {

	res := make([][]int, 0)
	for i := range s {
		res = append(res, make([]int, 0))
		for j := range s[0] {
			v, _ := strconv.Atoi(s[i][j])
			res[i] = append(res[i], v)
		}
	}
	return res
}
