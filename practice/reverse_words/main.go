package main

import "fmt"

func main() {
	fmt.Println(reverseWords("the sky is blue"))
}
func reverseWords(s string) string {

	var res string
	begin := len(s) - 1
	end := begin
	for {
		if begin < 0 {
			break
		}

		fmt.Println(begin, end)

		if s[begin] == ' ' {
			begin = begin - 1
			continue
		} else {
			// 从begin开始找第一个空格
			end = begin
			for {
				if end < 0 {
					// 找完
					break
				}
				if s[end] != ' ' {
					end--
				} else {
					break
				}
			}
			res = res + s[end+1:begin+1] + " "
			begin = end
		}

	}
	return res[:len(res)-1]
}
