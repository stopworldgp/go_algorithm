package main

import (
	"fmt"
	"go_algorithm/gp_queue"
)

func main() {

	q := gp_queue.NewLKqueue()
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(q.Dequeue())
	}
}
