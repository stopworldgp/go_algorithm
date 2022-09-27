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
	//如果尝试在每层递归都设置随机数种子，将会导致超时，因为这个操作太耗时了。我测试的结果，每次加时间戳种子的耗时是不加种子的452倍左右。
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
