package array

//283. 移动零 -https://leetcode.cn/problems/move-zeroes/
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
	return
}
