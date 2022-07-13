package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

// according to the Semaphore concurrensy pattern lecture slide
func getBatch(n int64, pool int64) (res []user) {
	var mu sync.Mutex
	res = make([]user, 0, n)
	var wg sync.WaitGroup
	sem := make(chan struct{}, pool)
	for i := int64(0); i < n; i++ {
		wg.Add(1)
		sem <- struct{}{} // right position. pool running goroutines
		go func(i int64) {
			// sem <- struct{}{}  // was inside the go func in the video (wrong).
			// 100% running goroutine. pool - working
			defer wg.Done()
			user := getOne(i)
			mu.Lock()
			res = append(res, user)
			mu.Unlock()
			<-sem
		}(i)
	}
	wg.Wait()

	return res
}
