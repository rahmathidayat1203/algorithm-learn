package main

func main() {
	arr := []int{1, 1, 2}
	removeDuplicates(arr)
}

func removeDuplicates(nums []int) int {
	var index int
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[index] {
			continue
		}
		index++
		nums[index] = nums[i]
	}
	return index
}

func removeDuplicatesArrayReturn(nums []int) []int {
	mapKey := make(map[int]bool, 0)
	list := []int{}
	for _, value := range nums {
		if _, v := mapKey[value]; !v {
			mapKey[value] = true
			list = append(list, value)
		}
	}
	return list
}
