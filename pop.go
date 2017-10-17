package deque

func (q *Deque) Pop() interface{} {
	if q.Len() == 0 {
		panic("pop from an empty deque")
	}
	n := len(q.view) - 1
	val := q.view[n]

	q.view[n] = nil
	q.view = q.view[:n]
	q.j -= 1
	return val
}

func (q *Deque) PopLeft() interface{} {
	if q.Len() == 0 {
		panic("pop from an empty deque")
	}
	val := q.view[0]
	q.view[0] = nil

	q.view = q.view[1:]
	q.i += 1
	return val
}
