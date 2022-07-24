package main

import (
	"fmt"

	"github.com/ukutluer/go-queue/queue"
)

func main() {
	queue := queue.NewQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	queue.Print()
	v, _ := queue.Peek()
	fmt.Println("peek item --> ", v)
	v, _ = queue.Dequeue()
	fmt.Println(v)
	v, _ = queue.Dequeue()
	fmt.Println(v)
	v, _ = queue.Dequeue()
	fmt.Println(v)
	v, _ = queue.Dequeue()
	fmt.Println(v)
}
