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
	if len(nums) == 1 {
		return nums[0]
	}
	var min int64 = math.MaxInt64
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	if min == math.MaxInt64 {
		fmt.Println("????????")
	}
	return min
}

func Max(nums []int64) int64 {
	if len(nums) == 1 {
		return nums[0]
	}
	var max int64 = math.MinInt64
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	if max == math.MinInt64 {
		fmt.Println("????????")

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
