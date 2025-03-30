package main

import "fmt"
func main() {
	var nums = []int{0,0,1,1,1,1,2,3,3}
    var count = removeDuplicates(nums)
	
    fmt.Println(count)
}

func removeDuplicates(nums []int) int {
	var last, lastIdx, counter int

	for idx, num := range nums {
		if idx == 0 {
			last = num
			counter = 1
			continue
		}

		if num != last {
			lastIdx = lastIdx + 1
			last = num
			nums[lastIdx] = num
			counter = 1
		} else {
			counter++
			if counter > 2 {
				continue
			}
			lastIdx = lastIdx + 1
            nums[lastIdx] = num
		}
	}
	return lastIdx + 1
}