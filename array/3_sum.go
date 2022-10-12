package array

import "sort"

//15. 三数之和-https://leetcode.cn/problems/3sum/
func threeSum(nums []int) [][]int {
	//0. 排序
	sort.Ints(nums)

	ans := make([][]int, 0)
	//1.将三数转成两个数
	for first := 0; first < len(nums); first++ {

		//2. 去重
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		target := 0 - nums[first]
		third := len(nums) - 1 //每次重置
		//3. 双指针
		for second := first + 1; second < len(nums); second++ {
			//去重
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			//4.处理大于target 情况
			for second < third && nums[second]+nums[third] > target {
				third--
			}

			//5.已经达到中间
			if second == third {
				break
			}

			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
			//继续寻找下一个&默认是second++情况
		}

	}
	return ans
}
