package scheduler

import (
	"context"
	"testing"
	"time"
)

func TestInitScheduler(t *testing.T) {

	init := 0
	runnable := func(ctx context.Context) {
		init++
	}
	sched := NewScheduler()
	sched.Add(context.Background(), runnable, 1*time.Second)

	time.Sleep(3 * time.Second)
	sched.Stop()
	if init != 3 {
		t.Errorf("Expected a total of %d but got %d", 3, init)
	}
}
