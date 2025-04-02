package main
import "fmt"
func main() {
	var nums = []int{2,2,1,1,1,2,2}
    majorityElement(nums)
    fmt.Println(nums)
}

// 摩尔投票法
// 核心思想为对拼消耗。
// 玩一个诸侯争霸的游戏，假设你方人口超过总人口一半以上，并且能保证每个人口出去干仗都能一对一同归于尽。最后还有人活下来的国家就是胜利。
// 那就大混战呗，最差所有人都联合起来对付你（对应你每次选择作为计数器的数都是众数），或者其他国家也会相互攻击（会选择其他数作为计数器的数），但是只要你们不要内斗，最后肯定你赢。
func majorityElement(nums []int) int {

	var target, count int
	for idx, num := range nums {
		if idx == 0 {
			target = num
			count = 1
            continue
		}

		if count == 0 {
			target = num
		}

		if num == target {
			count++
		}

		if num != target {
			count = count - 1
		}
	}

	return target
}