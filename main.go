package main

import (
	"fmt"
	"github.com/raihankhan/Go-drill/queue"
	"github.com/raihankhan/Go-drill/stack"
)

func main() {

	// ........... Queue ................
	fmt.Println(".......Queue........")
	q := queue.New()
	q.Push(3, "dass", 4)
	fmt.Println(q.Front())
	fmt.Println(q.Get())

	q.Pop()
	fmt.Println(q.Front())
	fmt.Println(q.Get())

	q.Clear()
	fmt.Println(q.Size())

	// ........... Stack ................
	fmt.Println(".......Stack........")
	st := stack.New()
	st.Push(3, "dass", 4)
	fmt.Println(st.Top())
	fmt.Println(st.Get())

	st.Pop()
	fmt.Println(st.Top())
	fmt.Println(st.Get())

	st.Clear()
	fmt.Println(st.Size())

}
