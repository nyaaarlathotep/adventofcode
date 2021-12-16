package FuncPkg

import (
	"fmt"
	"math"
)

func Adder(nums []int64) int64 {
	var res int64 = 0
	for _, v := range nums {
		res = res + v
	}
	return res
}

func Multiplier(nums []int64) int64 {
	var res int64 = 1
	for _, v := range nums {
		res = res * v
	}
	return res
}

func Min(nums []int64) int64 {
	var min int64 = math.MaxInt64
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(nums []int64) int64 {
	var max int64 = math.MinInt64
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func GreaterThan(nums []int64) int64 {
	if len(nums) != 2 {
		fmt.Println("????????")
	}
	if nums[0] > nums[1] {
		return 1
	} else {
		return 0
	}
}

func LessThan(nums []int64) int64 {
	if len(nums) != 2 {
		fmt.Println("????????")
	}
	if nums[0] < nums[1] {
		return 1
	} else {
		return 0
	}
}

func Equal(nums []int64) int64 {
	if len(nums) != 2 {
		fmt.Println("????????")
	}
	if nums[0] == nums[1] {
		return 1
	} else {
		return 0
	}
}
