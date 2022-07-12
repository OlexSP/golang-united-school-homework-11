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
	res = make([]user, 0, int(n))
	var wg sync.WaitGroup
	sem := make(chan struct{}, int(pool))
	for i := 0; i < int(n); i++ {
		wg.Add(1)
		sem <- struct{}{} // in the go func in the video (wrong). less running goroutines
		go func(i int) {  // it passes tests in bought position.
			defer wg.Done()
			user := getOne(int64(i))
			res = append(res, user)
			<-sem
		}(i)
	}
	wg.Wait()

	return res
}
