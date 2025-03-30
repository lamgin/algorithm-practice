package main

import "fmt"
func main() {
	var nums = []int{1,1,2,2,3,4,5,6,7}
    var count = removeDuplicates(nums)
    fmt.Println(count)
}

func removeDuplicates(nums []int) int {

    var last,lastIdx int
    for idx, num := range nums {
        if idx == 0 {
            // init
            last = num
            continue
        }
        if num != last {
			lastIdx = lastIdx + 1
            last = num
			nums[lastIdx] = num
        }
    }

    return lastIdx + 1
}