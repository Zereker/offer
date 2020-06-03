package main

type MaxQueue struct {
	slice []int
	max   int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.slice) == 0 {
		return -1
	}

	return this.max
}

func (this *MaxQueue) Push_back(value int) {
	if len(this.slice) == 0 {
		this.max = value
		this.slice = append(this.slice, value)
		return
	}

	if this.max <= value {
		this.slice = append(this.slice, this.max)
		this.max = value
		this.slice = append(this.slice, value)
		return
	}

	this.slice = append(this.slice, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.slice) == 0 {
		return -1
	}

	value := this.slice[0]
	if value == this.max {

	}
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */

func main() {

}
