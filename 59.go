package main

import (
	"math"
)

/*
面试题59 - I. 滑动窗口的最大值
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7


提示：

你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。

注意：本题与主站 239 题相同：https://leetcode-cn.com/problems/sliding-window-maximum/
*/

func maxSlidingWindow(nums []int, k int) []int {
	length := len(nums)

	if length*k == 0 {
		return []int{}
	}

	result := make([]int, 0, length-k+1)
	for i := 0; i < length-k+1; i++ {

		max := math.MinInt32
		for j := i; j < i+k; j++ {
			max = int(math.Max(float64(max), float64(nums[j])))
		}

		result = append(result, max)
	}

	return result
}

func main() {
	maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
}
