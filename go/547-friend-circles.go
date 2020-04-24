package main

func findCircleNum(M [][]int) int {
	circleCounted := make([]bool, len(M))
	var dfsFriend func(int)
	dfsFriend = func(i int) {
		if circleCounted[i] {
			return
		}
		circleCounted[i] = true
		for j := 0; j < len(M[i]); j++ {
			if j != i && M[i][j] == 1 {
				dfsFriend(j)
			}
		}
	}
	cnt := 0
	for i := 0; i < len(M); i++ {
		if circleCounted[i] {
			continue
		}
		cnt++
		dfsFriend(i)
	}
	return cnt
}
