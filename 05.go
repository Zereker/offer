package main

/*
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

 

示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."
 

限制：

0 <= s 的长度 <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func replaceSpace(s string) string {
	bytes := make([]rune, 0, len(s))
	for _, b := range []rune(s) {
		if b == ' ' {
			bytes = append(bytes, '%')
			bytes = append(bytes, '2')
			bytes = append(bytes, '0')
		} else {
			bytes = append(bytes, b)
		}
	}

	return string(bytes)
}
