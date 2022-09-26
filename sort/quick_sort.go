package sort

import (
	"math/rand"
	"time"
)

//QuickSort 快排算法
func QuickSort(nums []int) []int {
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, left, right int) {
	//1.基础条件
	if left >= right {
		return
	}
	//2.递归条件
	pivot := partition(nums, left, right)
	sort(nums, left, pivot)
	sort(nums, pivot+1, right)
}

func partition(nums []int, left, right int) int {
	rand.Seed(time.Now().Unix())
	pivot := left + rand.Intn(right-left)
	pivotValue := nums[pivot]
	for left <= right {
		for nums[left] < pivotValue {
			left++
		}
		for nums[right] > pivotValue {
			right--
		}
		if left >= right {
			break
		}
		//swap
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	//通过上述左右转移，right为左区最右侧下标
	return right
}
