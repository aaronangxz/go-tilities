package go_tilities

type Queue []interface{}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) Enqueue(d interface{}) {
	*q = append(*q, d)
}

func (q *Queue) Dequeue() interface{} {
	if len(*q) == 0 {
		return nil
	}
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) Front() interface{} {
	return (*q)[0]
}

func (q *Queue) Back() interface{} {
	index := len(*q) - 1
	return (*q)[index]
}
