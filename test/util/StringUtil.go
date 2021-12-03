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

func CharAt(str string, index int) string {
	tmp := strings.Split(str, "")
	return tmp[index]
}
