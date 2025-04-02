package main
import "fmt"
func main() {
	var nums = []int{64, 34, 25, 12, 22, 11, 90}
    bubbleSort(nums)
    fmt.Println(nums)
}

func bubbleSort(nums []int) {
    var count = len(nums)
    for i:=0;i<count;i++ {
        for j:=0;j<count-i-1;j++ {
            if nums[j] > nums[j+1] {
                nums[j], nums[j+1] = nums[j+1],nums[j]
            }
        }
    }
}