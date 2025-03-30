package main

import "fmt"
func main() {
	var nums = []int{0,1,2,2,3,0,4,2}
    var val = 2
    var count = removeElement(nums, val)
    fmt.Println(count)
}

func removeElement(nums []int, val int) int {
    var count = len(nums)
    var idx = 0
    if count <= 0 {
        return 0
    }
    for {
        if idx > count-1 {
            return count
        }
        if nums[idx] == val {
            nums[idx], nums[count-1] = nums[count-1] ,nums[idx]
            count = count - 1
        } else {
            idx = idx + 1
        }
    }
}