package main

type b struct {
	st bool
	v  int
}

type x struct {
	a int
	c int
}

var mp map[x]b

func sumHere(n []int, xx x, yy x) (res b) {
	if _, ok := mp[xx]; ok {
		return mp[xx]
	}
	var nw int = 0
	if xx.a == yy.a-1 {
		if xx.c == yy.c {
			res.st = true
			res.v = 1
		} else {
			res.st = false
			res.v = 0
		}
		return
	} else {
		xx.a += 1
		var t int = xx.c
		xx.c = t + n[xx.a]
		res = sumHere(n, xx, yy)
		if res.st {
			nw += res.v
		}
		xx.c = t - n[xx.a]
		res = sumHere(n, xx, yy)
		if res.st {
			nw += res.v
		}
		if nw > 0 {
			res.st = true
		} else {
			res.st = false
		}
		res.v = nw
		xx.a -= 1
		xx.c = t
		mp[xx] = res
		return
	}
}

func findTargetSumWays(nums []int, S int) int {
	mp = make(map[x]b)
	var l, i int = len(nums), 0
	var d b
	var mm, nn x
	mm.a = 0
	mm.c = 0 + nums[0]
	nn.a = l
	nn.c = S
	d = sumHere(nums, mm, nn)
	i += d.v
	mm.c = 0 - nums[0]
	d = sumHere(nums, mm, nn)
	i += d.v
	return i
}
