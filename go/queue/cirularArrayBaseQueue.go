package main

import (
	"errors"
	"fmt"
)

//Queue represents a queue that holds a slice
type Queue struct {
	array []int
	count int
}

// Enqueue adds a value at the end
func (q *Queue) Enqueue(i int) {

	q.array = append(q.array, i)
	q.count++
}

func remove(s []int, i int) []int {

	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

//Dequeue removes a value at the front
func (q *Queue) Dequeue() int {

	if len(q.array) > 0 {

		front := 0
		remveFrontItem := q.array[front]
		q.array = q.array[1:]
		q.count--
		return remveFrontItem

	}

	return -1
}

func (q *Queue) Peek() (int, error) {

	if q.IsEmpty() {

		return 0, errors.New("empty queue")

	}

	return q.array[0], nil
}

// LastElem - returns the last element of the queue
func (q Queue) IsEmpty() bool {

	return len(q.array) == 0
}

func main() {

	myQueue := Queue{}
	myQueue.Enqueue(123)
	myQueue.Enqueue(125)
	myQueue.Enqueue(124)
	fmt.Println(myQueue)

	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue)
}
