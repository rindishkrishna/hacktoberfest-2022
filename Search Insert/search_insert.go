package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}

func searchInsert(nums []int, target int) (result int) {
	result = len(nums)
	for i, v := range nums {
		if v > target {
			result = i
			break
		} else if v == target {
			result = i
			break
		}
	}
	return
}
