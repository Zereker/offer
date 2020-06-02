package main

type MinStack struct {
	data []int
	min  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.data) == 0 {
		this.data = append(this.data, x)
		this.min = x
		return
	}

	if x <= this.min {
		this.data = append(this.data, this.min, x)
		this.min = x
		return
	}

	this.data = append(this.data, x)
}

func (this *MinStack) Pop() {
	length := len(this.data)
	if length == 0 {
		return
	}

	data := this.data[length-1]
	this.data = this.data[:length-1]

	if data == this.min && len(this.data) != 0 {
		this.min = this.data[len(this.data)-1]
		this.data = this.data[:len(this.data)-1]
		return
	}

}

func (this *MinStack) Top() int {
	length := len(this.data)
	if length == 0 {
		return 0
	}

	return this.data[length-1]
}

func (this *MinStack) Min() int {
	if len(this.data) == 0 {
		return 0
	}

	return this.min
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	println(minStack.Min())

	minStack.Pop()
	println(minStack.Top())
	println(minStack.Min())
}
