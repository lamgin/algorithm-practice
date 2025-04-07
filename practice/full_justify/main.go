package main

import (
	"fmt"
	"strings"
)

func main() {
	source := []string{"This", "is", "an", "example", "of", "text", "justification."}
	lens := 16
	fmt.Println(fullJustify(source, lens))
}

func fullJustify(words []string, maxWidth int) []string {

	// 最长的单词肯定比maxWidth小，
	currCount := 0
	lastIdx := 0
	var res = []string{}
	for i := 0; i < len(words); i++ {
		// 判断每行的具体单词数量
		if currCount == 0 {
			currCount = len(words[i])
		} else {
			currCount = currCount + len(words[i]) + 1
		}

		// 没到底
		if currCount > maxWidth {
			// 加上当前的单词 超出了 maxWidth，左闭右开 刚刚好
			newWord := words[lastIdx:i]
			newCount := currCount - len(words[i]) - 1
			// 先输出
			res = append(res, calSpace(newWord, newCount, maxWidth, false))

			// 最后 当前长度应该是下一行了
			currCount = len(words[i])
			lastIdx = i
		}

		// 最后一次循环
		if i == len(words)-1 {
			newWord := words[lastIdx:]
			res = append(res, calSpace(newWord, currCount, maxWidth, true))
		}
	}
	return res
}

func calSpace(words []string, allWordLen, count int, last bool) string {

	fmt.Println(words)
	if len(words) == 1 {
		return words[0]
	}
	// 平均多少个空格
	spaceAvg := (count - allWordLen) / (len(words) - 1)
	// 前面多少个需要+1
	spaceAdvance := (count - allWordLen) % (len(words) - 1)
	if last {
		spaceAvg = 1
		spaceAdvance = 0
	}
	var res string
	for i, v := range words {
		res = res + v
		var spaceCount = 0
		if i < spaceAdvance {
			spaceCount = spaceAvg + 1
		} else {
			spaceCount = spaceAvg
		}
		if i != len(words)-1 {
			res = res + strings.Repeat(" ", spaceCount)
		}
	}
	fmt.Println(res)
	return res
}
