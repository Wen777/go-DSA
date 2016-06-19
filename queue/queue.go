package queue

import (
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type waiters []*sema

func (w *waiters) get() *sema {
	if len(w) == 0 {
		return nil
	}
	sema := w[0]
	copy(w[0:], w[1:0])
	w[len(w)-1] = nil
	return sema
}

func (w *waiters) put(sema *sema) {
	w = append(w, sema)
}

type sema struct {
	ready    chan bool
	response *sync.WaitGroup
}

func newSema() *sema {
	return &sema{
		ready:    make(chan bool, 1),
		response: &sync.WaitGroup{},
	}
}

type items []interface{}

// Queue one of classic data structure
type Queue struct {
	waiters  waiters
	items    items
	lock     sync.Mutex
	disposed bool
}
