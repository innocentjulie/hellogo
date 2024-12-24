/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */

// @lc code=start
func thirdMax(nums []int) int {
	var f = math.MinInt
	var s = math.MinInt
	var t = math.MinInt
	for i := 0; i < len(nums); i++ {
		if nums[i] > f {
			if f > s {
				if s > t {
					t = s
				}
				s = f
			}
			f = nums[i]
		} else if nums[i] == f {
			continue
		} else if nums[i] > s {
			if s > t {
				t = s
			}
			s = nums[i]
		} else if nums[i] == s {
			continue
		} else if nums[i] >= t {
			t = nums[i]
		}
	}
	if t == math.MinInt {
		return f
	}
	return t
}

// @lc code=end

