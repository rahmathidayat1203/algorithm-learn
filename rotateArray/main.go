package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr, 3)
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	result := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	for i := 0; i < len(nums); i++ {
		nums[i] = result[i]
	}

	fmt.Println(nums)
}
