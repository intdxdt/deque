package deque

import (
	"bytes"
	"fmt"
)

const N = 32

type Deque struct {
	base     []interface{}
	view     []interface{}
	i        int
	j        int
	initSize int
}

//Construct a new Deque
func NewDeque(initSize ...int) *Deque {
	var iSize = N
	if len(initSize) > 0 {
		iSize = maxInt(1, initSize[0])
	}
	base, view, i, j := initQue(iSize)
	return &Deque{
		base:     base,
		view:     view,
		i:        i,
		j:        j,
		initSize: iSize,
	}
}

//Clone Deque
func (q *Deque) Clone() *Deque {
	base := make([]interface{}, len(q.base))
	copy(base, q.base)
	view := base[q.i:q.j]
	return &Deque{
		base: base,
		view: view,
		i:    q.i,
		j:    q.j,
	}
}

//reserve enough space left or right
// sufficient to contain elements on insert
func (q *Deque) Reserve(left, right bool) {
	if left && q.i == 0 {
		q.expandBase()
	}

	if right && q.j == len(q.base) {
		q.expandBase()
	}
}

func (q *Deque) DataRange() (*int, *int) {
	return &q.i, &q.j
}

func (q *Deque) RawSlice() []interface{} {
	return q.base
}

func (q *Deque) DataView() *[]interface{} {
	return &q.view
}

//Reverse Deque in-pace
func (q *Deque) Reverse() *Deque {
	for i, j := 0, len(q.view)-1; i < j; i, j = i+1, j-1 {
		q.view[i], q.view[j] = q.view[j], q.view[i]
	}
	return q
}

//Append to right side of Deque
func (q *Deque) Append(o interface{}) *Deque {
	q.Reserve(false, true)
	q.base[q.j] = o
	q.j += 1
	q.view = q.base[q.i:q.j]
	return q
}

//Extend Deque given list of values as params
func (q *Deque) Extend(values ...interface{}) *Deque {
	for _, v := range values {
		q.Append(v)
	}
	return q
}

//Extend with deque another deque from the right
func (q *Deque) ExtendWithDeque(dq *Deque) *Deque {
	for _, v := range dq.view {
		q.Append(v)
	}
	return q
}

//Concat two Deque and returns a new Deque
func (q *Deque) Concat(dq *Deque) *Deque {
	concat := q.Clone()
	for _, v := range dq.view {
		concat.Append(v)
	}
	return concat
}

//AppendLeft: appends to left of Deque
func (q *Deque) AppendLeft(o interface{}) *Deque {
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

//Extend leftside of deque with given values as params oder
func (q *Deque) ExtendLeft(values ...interface{}) *Deque {
	for i := len(values) - 1; i >= 0; i-- {
		q.AppendLeft(values[i])
	}
	return q
}

//Extend left width deque
func (q *Deque) ExtendLeftWithDeque(dq *Deque) *Deque {
	for i := len(dq.view) - 1; i >= 0; i-- {
		q.AppendLeft(dq.view[i])
	}
	return q
}

//First value in Deque
func (q *Deque) Get(idx int) interface{} {
	if idx < 0 {
		idx += len(q.view)
	}
	return q.view[idx]
}

//First value in Deque
func (q *Deque) First() interface{} {
	return q.Get(0)
}

//Last value in Deque
func (q *Deque) Last() interface{} {
	return q.Get(-1)
}

//Length of number of items in Deque
func (q *Deque) Len() int {
	return len(q.view)
}

//Checks if Deque empty
func (q *Deque) IsEmpty() bool {
	return q.Len() == 0
}

//Clear everything in Deque
func (q *Deque) Clear() *Deque {
	q.base, q.view, q.i, q.j = initQue(q.initSize)
	return q
}

func (q *Deque) String() string {
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

//Loop through items in the queue with a callback
// if callback returns bool. Break looping with callback
// return as false
func (q *Deque) ForEach(fn func(interface{}, int) bool) {
	for i, v := range q.view {
		if !fn(v, i) {
			break
		}
	}
}
