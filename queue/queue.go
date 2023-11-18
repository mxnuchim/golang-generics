package queue

import (
	"bytes"
	"errors"
	"fmt"
)

type q[T any] struct {
	back int
	front int
	items []T
}

func New[T any]() *q[T] {
	return &q[T]{
		items: make([]T, 10),
		front: 0,
		back: -1,
	}
}

func (this *q[T]) Enqueue(item T) *q[T] {
	//if queue is full, create another slice for the items (remmber that the slice is initialized with a set length of q0)
	if this.back == len(this.items)-1 {
		fmt.Println("Expanding queue...")
		//create new list with longer length and copy existing list into it
		newList := make([]T, len(this.items)*2)
		copy(newList, this.items)
		this.items = newList
	}

	this.back++
	this.items[this.back] = item

	return this
}

func (this *q[T]) Dequeue() (T, error) {
	if this.back == -1 || this.front > this.back {
		var res T
		return res, errors.New("Queue is full")
	}

	res := this.items[this.front]
	this.front++
	return res, nil
}

func (this *q[T]) String() string {
	var buf bytes.Buffer

	for{
		v, err := this.Dequeue()
		if err != nil {
			break
		}
		buf.WriteString(fmt.Sprintf("%v ", v))
	}

	return buf.String()
}