package main

import (
	"fmt"
	"github.com/raihankhan/Go-drill/queue"
)

func main() {
	q := queue.New()
	q = q.Push(3, "dass", 4)
	q.Pop()
	fmt.Println(q.Size())
	fmt.Println(q.Get())
	fmt.Println(q.Front())

}
