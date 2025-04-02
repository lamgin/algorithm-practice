package main
import "fmt"
import "math"
func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
func jump(nums []int) int {
	var dpArr = make([]int, len(nums))
	for i, _ := range dpArr {
		dpArr[i] = math.MaxInt
	}
	minStep := dp(nums, dpArr, len(nums)-1)
	fmt.Println(dpArr)
	return minStep
}
func dp(nums []int, dpArr []int, cur int) int {
	// 推出条件
	if cur == 0 {
		return 0
	}
	// 找到直接可达当前cur的，做递归，然后计算最小值
	var minStep = math.MaxInt
	for i := 0; i < cur; i++ {
		if nums[i]+i >= cur {
			var step int
			if dpArr[i] < math.MaxInt {
				step = dpArr[i] + 1
			} else {
				step = dp(nums, dpArr, i) + 1
				dpArr[i] = step-1
			}
			if step < minStep {
				minStep = step
			}
		}
	}
	return minStep
}