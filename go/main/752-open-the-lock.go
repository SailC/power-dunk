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
				digits := nextDigits(lock[i])
				for _, digit := range digits {
					newLock := lock[0:i] + digit + lock[i+1:]
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

func nextDigits(digit byte) []string {
	result := []string{}
	digitStr := string(digit)
	digitInt, _ := strconv.Atoi(digitStr)
	if digitInt+1 == 10 {
		result = append(result, "0")
	} else {
		result = append(result, strconv.Itoa(digitInt+1))
	}
	if digitInt-1 == -1 {
		result = append(result, "9")
	} else {
		result = append(result, strconv.Itoa(digitInt-1))
	}
	return result
}

func main() {
	openLock([]string{"0000"}, "8888")
}
