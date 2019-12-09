package main

import "fmt"

func main() {
	// fmt.Println(isValid("()[]{}"))
	// fmt.Println(isValid("(]"))
	// fmt.Println(isValid("([)]"))
	// fmt.Println(isValid("{[]}"))
	// fmt.Println(removeKdigits("10200", 1))
	// fmt.Println(removeKdigits("1573421", 2))
	// fmt.Println(removeKdigits("1573421", 3))
	// fmt.Println(removeKdigits("1573421", 4))
	// fmt.Println(removeKdigits("1573421", 5))
	// fmt.Println(removeKdigits("1573421", 6))
	// fmt.Println(removeKdigits("1573421", 7))

	// fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))

	// fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
	// fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))

	// fmt.Println(asteroidCollision([]int{5, 10, -5}))
	// fmt.Println(asteroidCollision([]int{8, -8}))
	// fmt.Println(asteroidCollision([]int{10, 2, -5}))
	// fmt.Println(asteroidCollision([]int{-2, -1, 1, 2}))

	// fmt.Println(decodeAtIndex("leet2code3", 10))
	// fmt.Println(decodeAtIndex("ha22", 5))
	// fmt.Println(decodeAtIndex("a2345678999999999999999", 1))
	// fmt.Println(decodeAtIndex("a2b3c4d5e6f7g8h9", 10))
	// t1 := (time.Now())
	// fmt.Println(decodeAtIndex("cpmxv8ewnfk3xxcilcmm68d2ygc88daomywc3imncfjgtwj8nrxjtwhiem5nzqnicxzo248g52y72v3yujqpvqcssrofd99lkovg", 480551547))
	// fmt.Println(decodeAtIndex("y75lgfqyn4re8treuyrs9pqxfgvf3rrtqxr6lrm8ymxawwf97jcm5itz8dpvjig3o9iartdajjeoo58ipu6wmuozzpjzzfzrciy6", 292404628))
	// t2 := (time.Now())
	// fmt.Println(t2.Sub(t1).Seconds())

	// fmt.Println(simplifyPath("/home/"))
	// fmt.Println(simplifyPath("/../"))
	// fmt.Println(simplifyPath("/home//foo/"))
	// fmt.Println(simplifyPath("/a/./b/../../c/"))
	// fmt.Println(simplifyPath("/a/../../b/../c//.//"))
	// fmt.Println(simplifyPath("/a//b////c/d//././/.."))

	// fmt.Println(toLowerCase("HeLlo"))
	// fmt.Println([]byte("hello"))
	// fmt.Println([]byte("Hello"))
	// fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))

	// fmt.Println(convert("AB", 2))
	// fmt.Println(convert("LEETCODEISHIRING", 3))

	// fmt.Println(compress([]byte("aabbccc")))
	// fmt.Println(compress([]byte("abbbbbbbbbbbb")))
	// fmt.Println(compress([]byte("abc")))

	// fmt.Println(letterCombinations("234"))

	// fmt.Println(restoreIpAddresses("25525511135"))
	// fmt.Println(restoreIpAddresses("010010"))

	// cache := Constructor(1)
	// cache.Put(2, 1)
	// fmt.Println(cache.string())

	// fmt.Println(cache.Get(2))
	// fmt.Println(cache.string())

	// cache.Put(3, 2)
	// fmt.Println(cache.string())

	// fmt.Println(cache.Get(2))
	// fmt.Println(cache.string())

	// fmt.Println(cache.Get(3))
	// fmt.Println(cache.string())

	// fmt.Println(cache.Get(1)) // 返回  1
	// cache.Put(3, 3)           // 该操作会使得密钥 2 作废
	// fmt.Println(cache.Get(2)) // 返回 -1 (未找到)
	// cache.Put(4, 4)           // 该操作会使得密钥 1 作废
	// fmt.Println(cache.Get(1)) // 返回 -1 (未找到)
	// fmt.Println(cache.Get(3)) // 返回  3
	// fmt.Println(cache.Get(4)) // 返回  4

	// fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

	// fmt.Println(climbStairs(10))
	// fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
