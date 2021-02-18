package scheduler

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewScheduler(t *testing.T) {

	init := 0
	runnable := func(ctx context.Context) {
		init++
	}
	sched := NewScheduler()
	sched.Schedule(context.Background(), runnable, 1*time.Second)

	time.Sleep(3 * time.Second)
	sched.Stop()
	assert.Equal(t, 3, init)
}

func TestDefaultScheduler(t *testing.T) {

	init := 0
	job := func(ctx context.Context) {
		init++
	}
	sched := GetDefaultScheduler()
	sched.Schedule(context.Background(), job, 1*time.Second)

	time.Sleep(3 * time.Second)
	sched.Stop()
	assert.Equal(t, 3, init)
}
