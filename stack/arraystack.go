package stack

type stack struct {
	data []interface{}
}

func New() *stack {
	return &stack{
		data: []interface{}{},
	}
}

func (st *stack) Clear() {
	st.data = []interface{}{}
}

func (st *stack) Top() interface{} {
	if st.data == nil {
		return nil
	}
	topIndex := len(st.data) - 1
	return st.data[topIndex]
}

func (st *stack) Push(args ...interface{}) {
	for i := 0; i < len(args); i++ {
		st.data = append(st.data, args[i])
	}
}

func (st *stack) Pop() {
	if len(st.data) == 0 {
		return
	}
	topIndex := len(st.data) - 1
	st.data = st.data[:topIndex]
}

func (st *stack) Size() int {
	return len(st.data)
}

func (st *stack) Get() []interface{} {
	return st.data
}
