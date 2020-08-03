package main

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func minSubArrayLen(s int, nums []int) int {
	var i, j, min, l, cs int
	i = 0
	j = 0
	l = len(nums)
	min = l + 1
	for j < l {
		cs += nums[j]
		if cs == s {
			min = Min(min, j-i+1)
			j += 1
		} else if cs > s {
			cs -= nums[i]
			cs -= nums[j] // since it will be calculated again
			min = Min(min, j-i+1)
			i += 1
		} else {
			j += 1
		}
	}
	j = l - 1
	for i <= j {
		cs -= nums[i]
		if cs >= s {
			min = Min(min, j-i+1)
		}
		i += 1
	}
	if min == l+1 {
		return 0
	}
	return min
}
