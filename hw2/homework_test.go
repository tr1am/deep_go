package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// go test -v homework_test.go

type CircularQueue struct {
	values []int
	startq int
	endq   int
	// need to implement
}

func NewCircularQueue(size int) CircularQueue {
	var CQ CircularQueue
	CQ.values = make([]int, size)
	CQ.startq = 0
	CQ.endq = 0
	// init all elements
	for i := 0; i < size; i++ {
		CQ.values[i] = -1
	}
	return CQ
}
func (q *CircularQueue) Push(value int) bool {
	// calculate pointer for next element and the pointer to the last element of slice
	// to check borders of slice
	next := q.endq + 1
	last := len(q.values) - 1

	// Cannot add new element due to slice is full
	if q.Full() {
		return false
	}
	// Write the first value. Don't need to calculate new pointer
	// It's already point to first empty element
	if q.Empty() {
		q.endq = 0
		q.startq = 0
		q.values[0] = value
		return true
	}
	// move the End Queue pointer to the beginning and write first element
	if (q.endq == last) && (q.values[0] == -1) {
		q.endq = 0
		q.values[0] = value
		return true
	}
	// move the End Queue pointer to next element and write it
	if q.values[next] == -1 {
		q.values[next] = value
		q.endq = next
		return true
	}
	return false
}
func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}
	// if it's the last element in the slice, move to the first element
	if q.startq == len(q.values)-1 {
		q.values[q.startq] = -1
		q.startq = 0
		return true
	}
	q.values[q.startq] = -1
	q.startq++
	return true
}
func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.startq]
}
func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.endq]
}
func (q *CircularQueue) Empty() bool {
	for i := 0; i < len(q.values); i++ {
		if q.values[i] != -1 {
			return false
		}
	}
	return true
}
func (q *CircularQueue) Full() bool {
	for i := 0; i < len(q.values); i++ {
		if q.values[i] == -1 {
			return false
		}
	}
	return true
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
