package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

/*type timer struct {
	tb *timersBucket
	i  int

	when   int64
	period int64
	f      func(interface{}, uintptr)
	arg    interface{}
	seq    uintptr
}

type Timer struct {
	C <-chan Time
	r runTimeTimer
}

type timersBucket struct {
	lock        mutex
	gp          *g
	created     bool
	sleeping    bool
	recheduling bool
	sleepUntil  int64
	waitnote    note
	t           []*timer
}*/

func main() {
	time.NewTicker(time.Duration(60))
}

func RunTimers(count int) {
	durationCh := make(chan time.Duration, count)

	wg := sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			startedAt := time.Now()
			time.AfterFunc(10*time.Millisecond, func() {
				defer wg.Done()
				durationCh <- time.Since(startedAt)
			})
		}()

	}
	wg.Wait()

	close(durationCh)

	durations := []time.Duration{}
	totalDuration := 0 * time.Millisecond
	for duration := range durationCh {
		durations = append(durations, duration)
		totalDuration += duration
	}
	averageDuration := totalDuration / time.Duration(count)
	sort.Slice(durations, func(i, j int) bool {
		return durations[i] < durations[j]
	})

	fmt.Printf("run %v timers with average=%v, pct50=%v, pct99=%v\n", count, averageDuration, durations[count/2], durations[int(float64(count)*0.99)])
}
