package main

type MinStack struct {
	Value []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{Value: []int{}, min: []int{}}
}

func (this *MinStack) Push(val int) {
	this.Value = append(this.Value, val)
	if len(this.min) == 0 {
		this.min = append(this.min, val)
	} else if val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	if this.Value[len(this.Value)-1] == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}

	this.Value = this.Value[:len(this.Value)-1]
}

func (this *MinStack) Top() int {
	return this.Value[len(this.Value)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
