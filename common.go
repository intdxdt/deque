package deque

func maxInt(x, y int) int {
    if y > x {
        return y
    }
    return x
}

func initQue(initSize int) ([]interface{}, []interface{}, int, int) {
	i := initSize / 2
	j := i
	base := make([]interface{}, initSize, initSize)
	view := base[i:j]
	return base, view, i, j
}

func (q *Deque) expandBase() {
	bn := len(q.base)
	vn := len(q.view)

	nn := 2 * bn
	if vn+(nn/2-vn/2) >= nn {
		nn = 2 * nn //not big enough
	}

	k := nn / 2
	mid := vn / 2

	ii := k - mid
	jj := ii + vn

	newBase := make([]interface{}, nn)
	copy(newBase[k:], q.view[mid:])
	copy(newBase[k-mid:k], q.view[0:mid])
	q.base = newBase

	q.i, q.j = ii, jj
	q.view = q.base[q.i: q.j]
}

func (q *Deque) atPivot() bool {
	n := len(q.base)
	return q.i == q.j && (q.i >= 0 && q.i < n)
}
