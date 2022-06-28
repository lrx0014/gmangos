package fifo

import (
	"container/list"
	"sync"
)

type Queue struct {
	list   *list.List
	cap    int
	locker sync.RWMutex
}

func New(cap int) *Queue {
	return &Queue{
		list: list.New(),
		cap:  cap,
	}
}

func (q *Queue) Push(val []byte) {
	q.locker.Lock()
	defer q.locker.Unlock()
	if q.list.Len() > q.cap {
		e := q.list.Front()
		q.list.Remove(e)
	}

	q.list.PushBack(val)
}

func (q *Queue) Pop() (val []byte) {
	q.locker.Lock()
	defer q.locker.Unlock()
	if q.list.Len() > 0 {
		e := q.list.Front()
		val, _ = e.Value.([]byte)
		q.list.Remove(e)
	}

	return
}

func (q *Queue) All() (values [][]byte) {
	values = make([][]byte, q.list.Len())
	i := 0
	for e := q.list.Front(); e != nil; e = e.Next() {
		values[i], _ = e.Value.([]byte)
		i++
	}

	return
}

func (q *Queue) Write(p []byte) (n int, err error) {
	q.Push(p)
	return len(p), nil
}
