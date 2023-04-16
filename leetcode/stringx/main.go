package main

import (
	"fmt"
)

func main() {
	fmt.Println(partitionLabels("aebbedaddc"))
}

func partitionLabels(s string) []int {
	var arr [26]int
	for i := 0; i < len(s); i++ {
		ch := s[i] - 'a'
		arr[ch] = i
	}

	var ans []int
	i := 0
	l := 0
	for i < len(s) {
		ch := s[i] - 'a'
		if arr[ch] > i {
			i = arr[ch]
		} else {
			fmt.Println(arr[ch] == i)
			// 当前ch能满足需求了，需要看下从l到i-1是否能满足需求，可以的话，则结束，不可以的话，则继续loop
			for j := l; j < i; j++ {
				b := s[j] - 'a'
				if arr[b] > i {
					i = arr[b]
					break
				}
				fmt.Println(string(s[j]), b)
			}

			if i == arr[ch] {
				ans = append(ans, i-l+1)
				i++
				l = i
			}
		}
	}

	return ans
}

func judgeCircle(moves string) bool {
	var x, y int
	for _, m := range moves {
		switch m {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		case 'R':
			x++
		}
	}

	return x == 0 && y == 0
}
