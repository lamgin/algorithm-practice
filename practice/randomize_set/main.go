package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomizedSet := Constructor()
	fmt.Println(randomizedSet.Insert(1)) // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
	randomizedSet.Print()

	fmt.Println(randomizedSet.Insert(2)) // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
	randomizedSet.Print()

	fmt.Println(randomizedSet.GetRandom()) // getRandom 应随机返回 1 或 2 。
	randomizedSet.Print()

	fmt.Println(randomizedSet.Remove(1)) // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
	randomizedSet.Print()

	fmt.Println(randomizedSet.Insert(2)) // 2 已在集合中，所以返回 false 。
	randomizedSet.Print()

	fmt.Println(randomizedSet.GetRandom()) // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
	randomizedSet.Print()

}

type RandomizedSet struct {
	m    map[int]int
	data []int
}

func (this *RandomizedSet) Print() {
	fmt.Println(this.m)
	fmt.Println(this.data)
}

func Constructor() *RandomizedSet {
	return &RandomizedSet{
		m:    make(map[int]int), // value idx
		data: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if this == nil {
		return false
	}

	if _, ok := this.m[val]; ok {
		return false
	}

	this.data = append(this.data, val)
	idx := len(this.data) - 1
	this.m[val] = idx

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if this == nil {
		return false
	}
	if _, ok := this.m[val]; !ok {
		return false
	}

	idx := this.m[val]
	temp := this.data[len(this.data)-1]
	this.data[idx], this.data[len(this.data)-1] = this.data[len(this.data)-1], this.data[idx]
	this.m[temp] = idx
	this.data = this.data[:len(this.data)-1]
	delete(this.m, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	max := len(this.data)
	if max <= 0 {
		return 0
	}

	idx := rand.Intn(max)
	return this.data[idx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
