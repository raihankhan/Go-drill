package main

import (
	"fmt"
	"github.com/raihankhan/Go-drill/queue"
)

func main() {
	q := queue.New()
	q.Push(3, "dass", 4)
	fmt.Println(q.Front())
	fmt.Println(q.Get())

	q.Pop()
	fmt.Println(q.Front())
	fmt.Println(q.Get())

	q.Clear()
	fmt.Println(q.Size())

}
