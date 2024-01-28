package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrCountWorkersNull    = errors.New("count of workers not equal zero")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n <= 0 {
		return ErrCountWorkersNull
	}

	if m < 1 {
		m = 1
	}

	wg := &sync.WaitGroup{}
	queueCh := make(chan Task, len(tasks)+1)
	var errCount int32

	wg.Add(n)

	for i := 0; i < len(tasks); i++ {
		queueCh <- tasks[i]
	}
	close(queueCh)

	for i := 0; i < n; i++ {
		go worker(wg, queueCh, &errCount, m)
	}

	wg.Wait()

	var res error

	if int(errCount) >= m {
		res = ErrErrorsLimitExceeded
	}

	return res
}

// worker Работа воркеров.
func worker(wg *sync.WaitGroup, queue chan Task, errorsCount *int32, errorMaxCount int) {
	defer wg.Done()

	for task := range queue {
		err := task()
		if err != nil {
			atomic.AddInt32(errorsCount, 1)
		}

		if int(atomic.LoadInt32(errorsCount)) >= errorMaxCount {
			return
		}
	}
}
