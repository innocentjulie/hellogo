/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */
//version1 = "1.01", version2 = "1.001"

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	s1 := strings.Split(version1, ".")
	s2 := strings.Split(version2, ".")
	// fmt.Println(s1, s2)
	l1 := len(s1)
	l2 := len(s2)
	// fmt.Println(l1, l2)
	var slen = l1
	if l2 < l1 {
		slen = l2
	}
	for i := 0; i < slen; i++ {
		// result, _ = strconv.Atoi(singleStr)
		var v1, _ = strconv.Atoi(s1[i])
		var v2, _ = strconv.Atoi(s2[i])
		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}
	var vv = 0
	if l1 > l2 {
		for j := l2; j < l1; j++ {
			vv, _ = strconv.Atoi(s1[j])
			if vv > 0 {
				return 1
			}
		}
		return 0
	} else if l1 < l2 {
		for j := l1; j < l2; j++ {
			vv, _ = strconv.Atoi(s2[j])
			if vv > 0 {
				return -1
			}
		}
		return 0
	}
	return 0
}

// @lc code=end

