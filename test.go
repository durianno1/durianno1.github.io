package main

import "fmt"

func main() {
	a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(sortArray(a))
}

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}
func quickSort(nums []int, left int, right int) {
	mid := getMid(nums, left, right)
	quickSort(nums, left, mid-1)
	quickSort(nums, mid+1, right)
}
func getMid(nums []int, left int, right int) int {
	pivot := nums[left]
	for left < right {
		for nums[right] >= pivot && right > left {
			right--
		}
		nums[left] = nums[right]
		for nums[left] <= pivot && right > left {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}
