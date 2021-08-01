package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				s := []int{i, j}
				return s
				//return i,j
			}

		}

	}
	return nil
}

//注意下面这些都要写在主函数内，不要直接独立写出来会报错
func main() {
	nums := make([]int, 4)
	//var a int
	//for i := 0; i < len(nums); i++ {
	for i := 0; i < 4; i++ {
		fmt.Scan(&nums[i])
		//fmt.Scan(&a)
		//nums = append(nums, a)
		//fmt.scan(&nums[i])
	}

	fmt.Println(nums)
	var target int
	fmt.Scan(&target)
	fmt.Println(target)

	res := twoSum(nums, target)
	fmt.Println(res)
}
