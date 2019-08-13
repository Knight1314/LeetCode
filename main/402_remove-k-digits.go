package main

import (
	"strings"
)

/**
https://leetcode-cn.com/problems/remove-k-digits/submissions/

402. 移掉K位数字
给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。
注意:

num 的长度小于 10002 且 ≥ k。
num 不会包含任何前导零。
示例 1 :

输入: num = "1432219", k = 3
输出: "1219"
解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
示例 2 :

输入: num = "10200", k = 1
输出: "200"
解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
示例 3 :

输入: num = "10", k = 2
输出: "0"
解释: 从原数字移除所有的数字，剩余为空就是0。
*/
func removeKdigits(num string, k int) string {
	stack := []byte{}
	for i := 0; i < len(num); i++ {
		if i == 0 {
			stack = append(stack, num[i])
			continue
		}
		if len(stack) == 0 {
			stack = append(stack, num[i])
			continue
		}
		lstN := stack[len(stack)-1]
		if k > 0 && lstN > num[i] && len(stack) > 0 {
			stack = stack[:len(stack)-1]
			k--
			i--
		} else {
			stack = append(stack, num[i])
		}
	}
	stack = stack[:len(stack)-k]
	ret := strings.TrimLeft(string(stack), "0")
	if len(ret) == 0 {
		return "0"
	}
	return ret
}
