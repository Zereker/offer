package main

import (
	"errors"
	"fmt"
)

/*
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

 

示例 1：

输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
示例 2：

输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
提示：

1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type Stack interface {
	Empty() bool
	Pop() (int, error)
	Push(value int)
	Size() int
	Top() (int, error)
}

type stack struct {
	slices []int
}

func NewStack() Stack {
	slices := make([]int, 0)
	return &stack{
		slices: slices,
	}
}

func (s *stack) Empty() bool {
	return len(s.slices) == 0
}

func (s *stack) Pop() (int, error) {
	if s.Size() == 0 {
		return 0, errors.New("empty stack")
	}

	size := s.Size()
	result := s.slices[size-1]
	s.slices = s.slices[0 : size-1]
	return result, nil
}

func (s *stack) Push(value int) {
	s.slices = append(s.slices, value)
}

func (s *stack) Size() int {
	return len(s.slices)
}

func (s *stack) Top() (int, error) {
	if s.Size() == 0 {
		return 0, errors.New("empty stack")
	}

	length := len(s.slices)
	return s.slices[length-1], nil
}

type CQueue struct {
	inStack  Stack
	outStack Stack
}

func Constructor() CQueue {
	return CQueue{
		inStack:  NewStack(),
		outStack: NewStack(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack.Push(value)
}

func (this *CQueue) reverse() {
	for this.inStack.Size() != 0 {
		data, err := this.inStack.Pop()
		if err != nil {
			panic(err)
		}

		this.outStack.Push(data)
	}
}

func (this *CQueue) DeleteHead() int {
	if this.outStack.Size() == 0 {
		this.reverse()
	}

	pop, err := this.outStack.Pop()
	if err != nil {
		return -1
	}

	return pop
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func main() {
	queue := Constructor()
	queue.AppendTail(1)
	queue.AppendTail(2)

	head := queue.DeleteHead()
	fmt.Println(head)
}
