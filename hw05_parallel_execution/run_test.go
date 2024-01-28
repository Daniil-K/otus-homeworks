package hw05parallelexecution

import (
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("if were errors in first M tasks, than finished not more N+M tasks", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32

		for i := 0; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			tasks = append(tasks, func() error {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
				atomic.AddInt32(&runTasksCount, 1)
				return err
			})
		}

		workersCount := 10
		maxErrorsCount := 23
		err := Run(tasks, workersCount, maxErrorsCount)

		require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
		require.LessOrEqual(t, runTasksCount, int32(workersCount+maxErrorsCount), "extra tasks were started")
	})

	t.Run("tasks without errors", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 5
		maxErrorsCount := 1

		start := time.Now()
		err := Run(tasks, workersCount, maxErrorsCount)
		elapsedTime := time.Since(start)
		require.NoError(t, err)

		require.Equal(t, runTasksCount, int32(tasksCount), "not all tasks were completed")
		require.LessOrEqual(t, int64(elapsedTime), int64(sumTime/2), "tasks were run sequentially?")
	})
}

func TestRunFirstTaskWithError(t *testing.T) {
	defer goleak.VerifyNone(t)
	t.Run("many tasks - first error, single worker", func(t *testing.T) {
		tasksCount := 10
		tasks := make([]Task, 0, tasksCount)
		var runTasksCount int32
		tasks = append(tasks, makeErrorTask(&runTasksCount))

		for i := 0; i < tasksCount; i++ {
			tasks = append(tasks, makeSuccessTask(&runTasksCount))
		}

		workersCount := 1
		maxErrorsCount := 1
		err := Run(tasks, workersCount, maxErrorsCount)

		require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
		require.Equal(t, runTasksCount, int32(1), "not all tasks were completed")
	})

	t.Run("many tasks - first error, many worker", func(t *testing.T) {
		tasksCount := 10
		tasks := make([]Task, 0)
		var runTasksCount int32

		tasks = append(tasks, makeErrorTask(&runTasksCount))
		for i := 0; i < tasksCount; i++ {
			tasks = append(tasks, makeSuccessTask(&runTasksCount))
		}

		workersCount := 4
		maxErrorsCount := 1
		err := Run(tasks, workersCount, maxErrorsCount)

		require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
		require.LessOrEqual(t, runTasksCount, int32(1+workersCount*2), "not all tasks were completed")
	})
}

func TestRunComplete(t *testing.T) {
	defer goleak.VerifyNone(t)
	t.Run("many tasks - many error, many worker", func(t *testing.T) {
		tasksCount := 40
		tasks := make([]Task, 0)
		var runTasksCount int32
		for i := 0; i < tasksCount; i++ {
			if i%4 == 0 {
				tasks = append(tasks, makeErrorTask(&runTasksCount))
			} else {
				tasks = append(tasks, makeSuccessTask(&runTasksCount))
			}
		}

		workersCount := 4
		maxErrorsCount := 5
		err := Run(tasks, workersCount, maxErrorsCount)

		require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
		require.LessOrEqual(t, runTasksCount, int32(maxErrorsCount*4+workersCount), "not all tasks were completed")
	})
}

func makeSuccessTask(runTasksCount *int32) func() error {
	taskSleep := time.Millisecond * 10

	f := func() error {
		time.Sleep(taskSleep)
		atomic.AddInt32(runTasksCount, 1)
		return nil
	}

	return f
}

func makeErrorTask(runTasksCount *int32) func() error {
	taskSleep := time.Millisecond * 10

	f := func() error {
		time.Sleep(taskSleep)
		atomic.AddInt32(runTasksCount, 1)
		return errors.New("Bad task here")
	}

	return f
}
