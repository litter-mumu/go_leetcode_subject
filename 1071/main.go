package main

import (
	"bytes"
	"strings"
)

/*
	//TODO 对于字符串 S 和 T，只有在 S = T + ... + T（T 与自身连接 1 次或多次）时，我们才认定 “T 能除尽 S”。
	返回最长字符串 X，要求满足 X 能除尽 str1 且 X 能除尽 str2。

	示例 1：
	输入：str1 = "ABCABC", str2 = "ABC"
	输出："ABC"

	示例 2：
	输入：str1 = "ABABAB", str2 = "ABAB"
	输出："AB"

	示例 3：
	输入：str1 = "LEET", str2 = "CODE"
	输出：""

	提示：
	1 <= str1.length <= 1000
	1 <= str2.length <= 1000
	str1[i] 和 str2[i] 为大写英文字母
*/

func main() {

}

func gcdOfStrings(str1 string, str2 string) (maxYin string) {
	var answer []string
	var str3, str4 string
	if len(str1) > len(str2) {
		str3, str4 = str2, str1
	} else {
		str3, str4 = str1, str2
	}
	//rs := []rune(str)
	//由执行时间可知，这种获取[]rune方式效率更高
	rs := bytes.Runes(bytes.NewBufferString(str3).Bytes())
	for i, _ := range str3 {
		sub := string(rs[:i+1])

		if len(str3)%len(sub) == 0 {
			canTime := len(str3) / len(sub)
			if canTime == strings.Count(str3, sub) {
				answer = append(answer, sub)
			}
		} else {
			continue
		}
	}

	for _, v := range answer {
		if len(str4)%len(v) == 0 {
			canTime := len(str4) / len(v)
			if canTime == strings.Count(str4, v) && len(v) > len(maxYin) {
				maxYin = v
			}
		} else {
			continue
		}
	}
	return
}