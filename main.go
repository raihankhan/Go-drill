package main

import (
	"fmt"
	"github.com/raihankhan/Go-drill/queue"
)

func main() {
	q := queue.New()
	q.Insert(3, "dass", 4)
	fmt.Println(q)
}
