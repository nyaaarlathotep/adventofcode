package Util

import (
	"bytes"
	"strconv"
	"strings"
)

func getDecimalFromBinaryInt(binarySlice []int) int {
	var buffer bytes.Buffer
	for _, v := range binarySlice {
		buffer.WriteString(strconv.Itoa(v))
	}
	res, _ := strconv.ParseInt(buffer.String(), 2, 64)
	return int(res)
}

func getDecimalFromBinaryString(binarySlice []string) int {
	binary := strings.Join(binarySlice, "")
	res, _ := strconv.ParseInt(binary, 2, 64)
	return int(res)
}
