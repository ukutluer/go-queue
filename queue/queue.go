package queue

import (
	"errors"
	"fmt"
)

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append([]T{item}, q.items...)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if !q.IsEmpty() {
		willBeReturnedItem := q.items[len(q.items)-1]
		q.items = q.items[:len(q.items)-1]
		return willBeReturnedItem, nil
	}
	var result T
	return result, errors.New("THERE IS NO ITEM IN QUEUE")
}

func (q *Queue[T]) Peek() (T, error) {
	if !q.IsEmpty() {
		return q.items[len(q.items)-1], nil
	}
	var result T
	return result, errors.New("THERE IS NO ITEM IN QUEUE")
}

func (q *Queue[T]) IsEmpty() bool {
	return !(len(q.items) > 0)
}

func (q *Queue[T]) Print() {
	fmt.Printf("Queue has %d items. \n", len(q.items))
	fmt.Println(q.items)
}
func NewQueue[T any]() *Queue[T] {
	queue := &Queue[T]{}
	return queue
}
