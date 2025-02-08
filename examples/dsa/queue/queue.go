package main

import (
	"fmt"
)

type Queue struct {
	items []int
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}
	dequeuedItem := q.items[0]
	q.items = q.items[1:]
	return dequeuedItem
}

func main() {
	fmt.Println("Queue Implementation")
	queue := Queue{}
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)

	fmt.Print("Queue :: ")
	fmt.Println(queue.items)

	fmt.Print("Dequeued :: ")
	fmt.Println(queue.Dequeue())

	fmt.Print("Queue :: ")
	fmt.Println(queue.items)
}