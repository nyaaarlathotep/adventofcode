package main

import (
	"fmt"
	"io/ioutil"
	"test/DaySix"
	"time"
)

func main() {
	f, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("read fail", err)
	}

	start := time.Now()
	fmt.Println(DaySix.SixPart(string(f)))
	elapsed := time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}
