package main

/*

输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。



示例 1：

输入：head = [1,3,2]
输出：[2,3,1]


限制：

0 <= 链表长度 <= 10000
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	ints := make([]int, 0)

	for head != nil {
		ints = append(ints, head.Val)
		head = head.Next
	}

	result := make([]int, 0, len(ints))
	for i := len(ints) - 1; i >= 0; i-- {
		result = append(result, ints[i])
	}

	return result
}
