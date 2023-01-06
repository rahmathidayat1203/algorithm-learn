package main

import "fmt"

func main() {
	arr := []int{1, 1, 2}
	containsDuplicate(arr)
}

func containsDuplicate(nums []int) bool {
	var el int
	for i := 0; i < len(nums); i++ {
		el = nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == el {
				fmt.Println(nums[j])
				return true

			}

		}
	}
	return false
}
