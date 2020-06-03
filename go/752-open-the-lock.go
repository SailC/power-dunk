package main

import "strconv"

func openLock(deadends []string, target string) int {
	que := []string{"0000"}
	step := 0
	visited := make(map[string]bool)
	for _, deadLock := range deadends {
		visited[deadLock] = true
	}
	if _, exist := visited["0000"]; exist {
		return -1
	}
	visited["0000"] = true
	const digitNum = 4
	for len(que) > 0 {
		nextQue := []string{}
		for _, lock := range que {
			if lock == target {
				return step
			}
			for i := 0; i < digitNum; i++ {
				for _, newLock := range nextLocks(lock, i) {
					if _, exist := visited[newLock]; exist {
						continue
					}
					visited[newLock] = true
					nextQue = append(nextQue, newLock)
				}
			}
		}
		step++
		que = nextQue
	}
	return -1
}

func nextLocks(lock string, i int) []string {
	digitStr := string(lock[i])
	digitInt, _ := strconv.Atoi(digitStr)
	var plusOne int
	if digitInt == 9 {
		plusOne = 0
	} else {
		plusOne = digitInt + 1
	}
	plusOneLock := lock[0:i] + strconv.Itoa(plusOne) + lock[i+1:]
	var minusOne int
	if digitInt == 0 {
		plusOne = 9
	} else {
		plusOne = digitInt - 1
	}
	minusOneLock := lock[0:i] + strconv.Itoa(minusOne) + lock[i+1:]
	return []string{plusOneLock, minusOneLock}
}
