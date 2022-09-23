package search

//二分查找
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	//0.left > right
	for left <= right {
		//1. mid
		midIndex := (right + left) / 2
		//2. midValue
		midValue := nums[midIndex]
		if midValue == target {
			return midIndex
		}
		//3.binarySearch
		if target > midValue {
			left = midIndex + 1
		} else {
			right = midIndex - 1
		}
	}
	return -1
}
