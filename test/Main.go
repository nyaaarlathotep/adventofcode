package main

import (
	"fmt"
	"io/ioutil"
	"test/Day19"
	"time"
)

func main() {
	f, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("read fail", err)
	}

	start := time.Now()
	Day19.Day19(string(f))
	//fmt.Println(Day19.Day19(string(f)))
	elapsed := time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}
