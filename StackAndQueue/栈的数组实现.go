package StackAndQueue

const Cap = 8

type ArrayStack struct {
	count int
	array [Cap]int
}

func (array *ArrayStack) Push(val int) {
	if array.count == Cap {
		panic("ArrayStack is full")
	}
	array.count += 1
	array.array[array.count] = val
}

func (array *ArrayStack) Pop() int {
	if array.IsEmpty() {
		panic("ArrayStack is empty")
	}

	element := array.array[array.count]
	array.count -= 1
	return element
}

func (array *ArrayStack) IsEmpty() bool {
	return array.count == 0
}

func (array *ArrayStack) Top() int {
	return array.array[array.count]
}
