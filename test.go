package main

import "fmt"

func main() {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	heap_sort(a)
	fmt.Println(a)
}

func heap_sort(nums []int) {
	lens := len(nums) - 1
	for i := lens / 2; i >= 0; i-- { // 建堆 O(n) lens/2后面都是叶子节点，不需要向下调整
		down(nums, i, lens)
	}
	for j := lens; j >= 1; j-- { //堆排序（升序）:堆顶(最大值)交换到末尾
		nums[0], nums[j] = nums[j], nums[0]
		lens--
		down(nums, 0, lens)
	}
}
func down(nums []int, i, lens int) { //O(logn)大根堆，如果堆顶节点小于叶子，向下调整
	max := i
	if i<<1+1 <= lens && nums[i<<1+1] > nums[max] {
		max = i<<1 + 1
	}
	if i<<1+2 <= lens && nums[i<<1+2] > nums[max] {
		max = i<<1 + 2
	}
	if max != i {
		nums[max], nums[i] = nums[i], nums[max]
		down(nums, max, lens)
	}
}
