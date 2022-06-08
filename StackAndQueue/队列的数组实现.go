package StackAndQueue

type ArrayQueue struct {
	start int
	end   int
	array [Cap]int
}

func (q *ArrayQueue) Push(val int) {
	if q.end == Cap {
		panic("ArrayQueue is full")
	}
	q.end += 1
	q.array[q.end] = val
}

func (q *ArrayQueue) Get() int {
	if q.IsEmpty() {
		panic("ArrayQueue is empty")
	}
	element := q.array[q.start]
	q.start += 1
	return element
}

func (q *ArrayQueue) IsEmpty() bool {
	if q.start > q.end {
		return true
	}
	return false
}
