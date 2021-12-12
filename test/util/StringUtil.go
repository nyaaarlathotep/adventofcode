package util

import (
	"bytes"
	"strconv"
	"strings"
	"unicode"
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

// IsUpper 判断字符 r 是否为大写格式
func IsUpper(s string) bool {
	for _, r := range s {
		// 判断字符是否为大写
		if unicode.IsUpper(r) {
			return true
		} else {
			return false
		}
	}
	return false
}

func CopyStringSlice(in []string) []string {
	re := make([]string, 0)
	for _, s := range in {
		re = append(re, s)
	}
	return re
}

func RemoveByLoop(slc []string) []string {
	var res []string
	for i := range slc {
		flag := true
		for j := range res {
			if slc[i] == res[j] {
				flag = false
				break
			}
			if flag {
				res = append(res, slc[i])
			}
		}

	}
	return res
}
