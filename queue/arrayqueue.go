package queue

type queue struct {
	data []interface{}
}

func New() queue {
	return queue{
		data: []interface{}{},
	}
}

func (q *queue) Size() int {
	return len(q.data)
}

func (q *queue) Insert(args ...interface{}) queue {
	newQueue := append(q.data, args)
	return queue{
		data: newQueue,
	}
}
