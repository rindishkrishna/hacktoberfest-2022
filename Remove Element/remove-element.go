package main

import "fmt"

// https://leetcode.com/problems/remove-element/

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {
	cleanElement := make([]int, 0)
	for _, num := range nums {
		if num != val {
			cleanElement = append(cleanElement, num)
		}
	}
	return len(cleanElement)
}
