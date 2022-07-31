package queue

type queue struct {
	data []interface{}
}

func New() *queue {
	return &queue{
		data: []interface{}{},
	}
}

func (q *queue) Clear() {
	q.data = []interface{}{}
}

func (q *queue) Front() interface{} {
	if q.data == nil {
		return nil
	}
	return q.data[0]
}

func (q *queue) Push(args ...interface{}) {
	for i := 0; i < len(args); i++ {
		q.data = append(q.data, args[i])
	}
}

func (q *queue) Pop() {
	if len(q.data) == 0 {
		return
	}
	q.data = q.data[1:]
}

func (q *queue) Size() int {
	return len(q.data)
}

func (q *queue) Get() []interface{} {
	return q.data
}
