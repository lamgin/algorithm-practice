package main
import "fmt"
func canJump(nums []int) bool {
    var mark = len(nums)-1
    var remark = mark
    // 最后一位前一位
    for {

        for i := mark-1; i>=0; i-- {
            if nums[i] >= mark - i {
				fmt.Println(i, mark)
                mark = i
                break
            }
        }
		if mark <= 0 {
            return true
        }
        // 没有变化 证明没有可以跳到这里的值
        if remark == mark {
            return false
        }
        remark = mark
    }
    return mark < 0
}
func main() {
	nums := []int{2,3,1,1,4}
	fmt.Println(canJump(nums))
}