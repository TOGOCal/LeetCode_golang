package main

/*
给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，
满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。
*/

func containsNearbyDuplicate(nums []int, k int) bool {

	valueIndexMap := make(map[int]int)

	for index, value := range nums {

		getOrNot, ok := valueIndexMap[value]

		if !ok {

			valueIndexMap[value] = index
		} else {
			//检查是否符合题目要求

			if index-getOrNot <= k {

				return true
			}

			valueIndexMap[value] = index //更新
		}

	}

	return false
}
