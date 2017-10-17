package deque

import (
	"fmt"
	"time"
	"testing"
	"math/rand"
	"github.com/franela/goblin"
)

func TestDeque(t *testing.T) {
	g := goblin.Goblin(t)
	rand.Seed(time.Now().Unix())

	g.Describe("Deque", func() {
		var dq *Deque
		var array []int
		var rArray []int
		var lArray []int
		g.BeforeEach(func() {
			array = []int{9, 5, 3, 2, 8, 1, 2, 3}
			dq = NewDeque()
			for _, v := range array {
				dq.Append(v)
			}
			dq.AppendLeft(0)
			fmt.Println(dq)
		})

		g.It("should test deque as queue", func() {
			g.Assert(dq.IsEmpty()).IsFalse()
			g.Assert(dq.Len()).Equal(len(array) + 1)
			g.Assert(dq.First().(int)).Equal(0)
			g.Assert(dq.PopLeft().(int)).Equal(0)

			g.Assert(dq.First().(int)).Equal(9)
			g.Assert(dq.PopLeft().(int)).Equal(9)
			g.Assert(dq.PopLeft().(int)).Equal(5)
			g.Assert(dq.PopLeft().(int)).Equal(3)
			g.Assert(dq.PopLeft().(int)).Equal(2)
			g.Assert(dq.Last().(int)).Equal(3)
			fmt.Println(dq)

			dq.Clear()
			g.Assert(dq.IsEmpty()).IsTrue()
			//print
			fmt.Println(dq)
		})

		g.It("should test deque", func() {
			n := len(array) + 1
			ri, rj := dq.DataRange()
			g.Assert([]int{*ri, *rj}).Eql([]int{dq.i, dq.j})
			g.Assert(dq.RawSlice()).Eql(dq.base)
			g.Assert(*dq.DataView()).Eql(dq.view)
			g.Assert(dq.IsEmpty()).IsFalse()
			g.Assert(dq.Len()).Equal(n)

			g.Assert(dq.First().(int)).Equal(0)
			g.Assert(dq.PopLeft().(int)).Equal(0)
			n--
			g.Assert(dq.Len()).Equal(n)

			g.Assert(dq.Last().(int)).Equal(3)
			g.Assert(dq.Pop().(int)).Equal(3)
			n--
			g.Assert(dq.Len()).Equal(n)

			g.Assert(dq.First().(int)).Equal(9)
			g.Assert(dq.PopLeft().(int)).Equal(9)
			n--
			g.Assert(dq.Len()).Equal(n)

			g.Assert(dq.Pop().(int)).Equal(2)
			g.Assert(dq.Pop().(int)).Equal(1)
			g.Assert(dq.Pop().(int)).Equal(8)
			g.Assert(dq.Last().(int)).Equal(2)
			fmt.Println(dq)
			dq.ExtendLeft(0, 9, 5)
			g.Assert(dq.First().(int)).Equal(0)
			g.Assert(dq.PopLeft().(int)).Equal(0)
			g.Assert(dq.PopLeft().(int)).Equal(9)
			g.Assert(dq.PopLeft().(int)).Equal(5)
			dq.Extend(0, 9, 5)
			fmt.Println(dq)
			g.Assert(dq.Last().(int)).Equal(5)
			g.Assert(dq.Pop().(int)).Equal(5)
			g.Assert(dq.Pop().(int)).Equal(9)
			g.Assert(dq.Pop().(int)).Equal(0)

			g.Assert(dq.Len()).Equal(3)
			dq.Clear()
			g.Assert(dq.IsEmpty()).IsTrue()
			//print
			fmt.Println(dq)

			dqA := NewDeque(1)
			for _, v := range []int{0, 1, 2, 3, 4, 5, 6} {
				dqA.Append(v)
			}
			dqB := dqA.Clone()
			dqB.Extend(7, 8, 9)
			dqC := NewDeque().Extend(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
			dqD := NewDeque().Extend(7, 8, 9)

			g.Assert(dqA.Len()).Equal(7)
			g.Assert(dqB.Len()).Equal(10)
			g.Assert(dqD.Len()).Equal(3)
			g.Assert(dqC.view).Eql(dqB.view)
			g.Assert(dqC.view).Eql(dqA.Concat(dqD).view)
			g.Assert(dqA.Len()).Equal(7)

			//modify dqA with another deque
			dqA.ExtendWithDeque(dqD)
			g.Assert(dqA.Len()).Equal(10)
			g.Assert(dqD.Len()).Equal(3)

			g.Assert(dqA.Get(-1).(int)).Equal(9)
			g.Assert(dqA.Get(-2).(int)).Equal(8)
			g.Assert(dqA.Get(-3).(int)).Equal(7)

			dqA.ExtendLeftWithDeque(dqD)
			g.Assert(dqA.Get(0).(int)).Equal(7)
			g.Assert(dqA.Get(1).(int)).Equal(8)
			g.Assert(dqA.Get(2).(int)).Equal(9)
			fmt.Println(dqA)

			g.Assert(dqD.Get(0).(int)).Equal(7)
			g.Assert(dqD.Get(1).(int)).Equal(8)
			g.Assert(dqD.Get(2).(int)).Equal(9)
			dqD.Reverse()
			g.Assert(dqD.Get(0).(int)).Equal(9)
			g.Assert(dqD.Get(1).(int)).Equal(8)
			g.Assert(dqD.Get(2).(int)).Equal(7)

			dqC = NewDeque().Extend(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
			cols := []int{}
			dqC.ForEach(func(v interface{}, i int) bool {
				val := v.(int)
				if val < 5 {
					cols = append(cols, val)
					return true
				}
				g.Assert(i).Equal(5)
				return false
			})
			g.Assert(cols).Eql([]int{0, 1, 2, 3, 4})
		})

		g.It("should test exceeding base size pop & popleft", func() {
			for i := 0; i < 1000; i++ {
				rArray = append(rArray, rand.Intn(10000))
				lArray = append(lArray, rand.Intn(10000))
			}
			lq, rq := NewDeque(-9), NewDeque(1)
			for _, v := range lArray {
				lq.AppendLeft(v)
			}
			for _, v := range rArray {
				rq.Append(v)
			}

			for i := len(lArray) - 1; i >= 0; i-- {
				g.Assert(lArray[i]).Equal(lq.PopLeft().(int))
			}

			g.Assert(lq.IsEmpty()).IsTrue()
			for i := len(rArray) - 1; i >= 0; i-- {
				g.Assert(rArray[i]).Equal(rq.Pop().(int))
			}
			g.Assert(rq.IsEmpty()).IsTrue()

			func() {
				defer func() {
					r := recover()
					g.Assert(r != nil).IsTrue()
				}()
				g.Assert(lq.IsEmpty()).IsTrue()
				lq.PopLeft()
			}()
			func() {
				defer func() {
					r := recover()
					g.Assert(r != nil).IsTrue()
				}()
				g.Assert(rq.IsEmpty()).IsTrue()
				rq.Pop()
			}()
		})
	})
}
