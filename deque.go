package deque

import (
	"bytes"
	"fmt"
)

const N = 32

type Deque[T any] struct {
	base     []T
	view     []T
	i        int
	j        int
	initSize int
}

func NewDeque[T any](initSize ...int) *Deque[T] {
	var iSize = N
	if len(initSize) > 0 {
		iSize = max(1, initSize[0])
	}
	var base, view, i, j = initQue[T](iSize)
	return &Deque[T]{
		base:     base,
		view:     view,
		i:        i,
		j:        j,
		initSize: iSize,
	}
}

// Clone Deque
func (q *Deque[T]) Clone() *Deque[T] {
	var base = make([]T, len(q.base))
	copy(base, q.base)
	view := base[q.i:q.j]
	return &Deque[T]{
		base: base,
		view: view,
		i:    q.i,
		j:    q.j,
	}
}

// Reserve enough space left or right
// sufficient to contain elements on insert
func (q *Deque[T]) Reserve(left, right bool) {
	if left && q.i == 0 {
		q.expandBase()
	}

	if right && q.j == len(q.base) {
		q.expandBase()
	}
}

func (q *Deque[T]) DataRange() (*int, *int) {
	return &q.i, &q.j
}

func (q *Deque[T]) RawSlice() []T {
	return q.base
}

func (q *Deque[T]) DataView() *[]T {
	return &q.view
}

// Reverse Deque in-pace
func (q *Deque[T]) Reverse() *Deque[T] {
	for i, j := 0, len(q.view)-1; i < j; i, j = i+1, j-1 {
		q.view[i], q.view[j] = q.view[j], q.view[i]
	}
	return q
}

// Append to right side of Deque
func (q *Deque[T]) Append(o T) *Deque[T] {
	q.Reserve(false, true)
	q.base[q.j] = o
	q.j += 1
	q.view = q.base[q.i:q.j]
	return q
}

// Extend Deque given list of values as params
func (q *Deque[T]) Extend(values ...T) *Deque[T] {
	for _, v := range values {
		q.Append(v)
	}
	return q
}

// ExtendWithDeque - extends deque with another deque from the right
func (q *Deque[T]) ExtendWithDeque(dq *Deque[T]) *Deque[T] {
	for _, v := range dq.view {
		q.Append(v)
	}
	return q
}

// Concat two Deque and returns a new Deque
func (q *Deque[T]) Concat(dq *Deque[T]) *Deque[T] {
	concat := q.Clone()
	for _, v := range dq.view {
		concat.Append(v)
	}
	return concat
}

// AppendLeft - appends to left of Deque
func (q *Deque[T]) AppendLeft(o T) *Deque[T] {
	q.Reserve(true, false)

	if q.atPivot() {
		q.j += 1
	} else {
		q.i -= 1
	}
	q.base[q.i] = o

	q.view = q.base[q.i:q.j]
	return q
}

// ExtendLeft - extend leftside of deque with given values as params oder
func (q *Deque[T]) ExtendLeft(values ...T) *Deque[T] {
	for i := len(values) - 1; i >= 0; i-- {
		q.AppendLeft(values[i])
	}
	return q
}

// ExtendLeftWithDeque - extend left width deque
func (q *Deque[T]) ExtendLeftWithDeque(dq *Deque[T]) *Deque[T] {
	for i := len(dq.view) - 1; i >= 0; i-- {
		q.AppendLeft(dq.view[i])
	}
	return q
}

// Get first value in Deque
func (q *Deque[T]) Get(idx int) interface{} {
	if idx < 0 {
		idx += len(q.view)
	}
	return q.view[idx]
}

// First value in Deque
func (q *Deque[T]) First() interface{} {
	return q.Get(0)
}

// Last value in Deque
func (q *Deque[T]) Last() interface{} {
	return q.Get(-1)
}

// Len - length of number of items in Deque
func (q *Deque[T]) Len() int {
	return len(q.view)
}

// IsEmpty - checks if Deque empty
func (q *Deque[T]) IsEmpty() bool {
	return q.Len() == 0
}

// Clear everything in Deque
func (q *Deque[T]) Clear() *Deque[T] {
	q.base, q.view, q.i, q.j = initQue[T](q.initSize)
	return q
}

func (q *Deque[T]) String() string {
	var buffer bytes.Buffer
	n := q.Len()
	buffer.WriteString("[")
	for i, o := range q.view {
		token := fmt.Sprintf("%v", o)
		if i < n-1 {
			token += ", "
		}
		buffer.WriteString(token)
	}
	buffer.WriteString("]")
	return buffer.String()
}

// ForEach : Loop through items in the queue with a callback
// if callback returns bool. Break looping with callback
// return as false
func (q *Deque[T]) ForEach(fn func(T, int) bool) {
	for i, v := range q.view {
		if !fn(v, i) {
			break
		}
	}
}
