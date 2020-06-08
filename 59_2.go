package main

type MaxQueue struct {
	slice []int
	maxs  []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.maxs) == 0 {
		return -1
	}

	return this.maxs[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.slice = append(this.slice, value)

	for len(this.maxs) != 0 && this.maxs[len(this.maxs)-1] < value {
		this.maxs = this.maxs[:len(this.maxs)-1]
	}

	this.maxs = append(this.maxs, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.slice) == 0 {
		return -1
	}

	value := this.slice[0]
	this.slice = this.slice[1:]

	if value == this.maxs[0] {
		this.maxs = this.maxs[1:]
	}

	return value
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */

func main() {
	queue := Constructor()
	/*
		println(queue.Max_value()) // -1
		queue.Push_back(1)
		println(queue.Max_value()) // 1
		queue.Push_back(2)
		println(queue.Max_value()) // 2
		queue.Push_back(3)
		println(queue.Max_value()) // 3
		println(queue.Pop_front()) // 3
		println(queue.Pop_front()) // 2
		println(queue.Pop_front()) // 1
		println(queue.Max_value()) // -1
	*/

	/*
		println(queue.Max_value()) // -1
		println(queue.Pop_front()) // -1
		println(queue.Max_value()) // -1
	*/

	/*
		println(queue.Max_value()) // -1
		queue.Push_back(1)
		queue.Push_back(2)

		println(queue.Max_value()) // 2
		println(queue.Pop_front()) // 1
		println(queue.Max_value()) // 2
	*/

	println(queue.Max_value()) // -1
	println(queue.Max_value()) // -1
	println(queue.Pop_front()) // -1
	println(queue.Pop_front()) // -1

	queue.Push_back(94)
	queue.Push_back(16)
	queue.Push_back(89)
	println(queue.Pop_front()) // 94
	queue.Push_back(22)
	println(queue.Pop_front()) // 16
}
