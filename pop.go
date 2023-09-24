package deque

func (q *Deque[T]) Pop() T {
	var empty T
	if q.Len() == 0 {
		panic("pop from an empty deque")
	}
	n := len(q.view) - 1
	val := q.view[n]

	q.view[n] = empty
	q.view = q.view[:n]
	q.j -= 1
	return val
}

func (q *Deque[T]) PopLeft() T {
	var empty T
	if q.Len() == 0 {
		panic("pop from an empty que")
	}
	val := q.view[0]
	q.view[0] = empty

	q.view = q.view[1:]
	q.i += 1
	return val
}
