package queue

type queue struct {
	data []interface{}
}

func New() queue {
	return queue{
		data: []interface{}{},
	}
}

func (q queue) Front() interface{} {
	if q.data == nil {
		return nil
	}
	return q.data[0]
}

func (q queue) Push(args ...interface{}) queue {
	newQueue := q.data
	for i := 0; i < len(args); i++ {
		newQueue = append(newQueue, args[i])
	}
	return queue{
		data: newQueue,
	}
}

func (q queue) Pop() queue {
	if len(q.data) == 0 {
		return q
	}
	return queue{
		data: q.data[1:],
	}
}

func (q queue) Size() int {
	return len(q.data)
}

func (q queue) Get() []interface{} {
	return q.data
}
