package main

func main() {
	arr := []int{5, 1, 1, 2, 2, 3, 3, 4, 4}
	singleNumber(arr)
}

func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		res ^= n
	}
	return res
}
