package main

import (
	"fmt"
	"reflect"
	"sync"
)

type WorkPool interface {
	Run()
	worker()
}

type GenericPool struct {
	Tasks       []Task
	Concurrency int
	taskChan    chan Task
	wg          sync.WaitGroup
}

func (wp *GenericPool) worker() {
	for task := range wp.taskChan {
		task.run()
		wp.wg.Done()
	}
}

func (wp *GenericPool) Run() {
	// Initialize the tasks channel
	wp.taskChan = make(chan Task, len(wp.Tasks))

	// Start workers
	for i := 0; i < wp.Concurrency; i++ {
		go wp.worker()
	}

	// Send tasks to the tasks channel
	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		wp.taskChan <- task
	}
	close(wp.taskChan)

	// Wait for all tasks to finish
	wp.wg.Wait()
}

type Task interface {
	run() error
}

type GenericTask struct {
	Fn interface{}
}

func (g *GenericTask) run() error {
	fn := reflect.ValueOf(g.Fn)
	if fn.Kind() != reflect.Func {
		return fmt.Errorf("task function must be a function")
	}

	// Check if the function has any input arguments
	if fn.Type().NumIn() > 0 {
		// Create input arguments
		in := make([]reflect.Value, fn.Type().NumIn())
		for i := 0; i < len(in); i++ {
			in[i] = reflect.New(fn.Type().In(i)).Elem()
		}

		// Call the function
		out := fn.Call(in)
		if len(out) > 0 && !out[0].IsNil() {
			return out[0].Interface().(error)
		}
	} else {
		// Call the function without arguments
		out := fn.Call(nil)
		if len(out) > 0 && !out[0].IsNil() {
			return out[0].Interface().(error)
		}
	}
	return nil
}

func ab(a int, b int) int {
	return a + b
}

func main() {
	add := func(a int, b int) int {
		return a + b
	}

	tasks := []Task{
		&GenericTask{add},
		&GenericTask{ab},
	}

}
