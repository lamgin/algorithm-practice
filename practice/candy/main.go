package main

import "fmt"

func main() {
	fmt.Println(candy([]int{1, 0, 2}))
}

func candy(ratings []int) int {
	// 首先 给孩子这么评分是十分不人道的行为 我表示严厉谴责

	left := make([]int, len(ratings))
	right := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		left[i] = 1
		right[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}

	for i := len(ratings) - 1 - 1; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}

	var sum int
	for i := 0; i < len(ratings); i++ {
		sum += max(left[i], right[i])
	}
	fmt.Println(left)
	fmt.Println(right)
	return sum

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
