package sort

//SelectionSort 选择排序
func SelectionSort(nums []int) []int {
	//1. index 留出替换的位置
	for index := 0; index < len(nums)-1; index++ {
		small := index
		//比较index剩余的
		for i := index + 1; i < len(nums); i++ {
			if nums[small] > nums[i] {
				small = i
			}
		}
		swap(nums, index, small)
	}
	return nums
}

func swap(nums []int, index, small int) {
	tmp := nums[small]
	nums[small] = nums[index]
	nums[index] = tmp
}
