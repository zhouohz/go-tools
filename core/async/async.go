package async

import "sync"

// WaitAll 等待所有任务执行完毕
func WaitAll(done chan struct{}, tasks ...func() interface{}) {
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(t func() interface{}) {
			defer wg.Done()
			t()
		}(task)
	}

	go func() {
		wg.Wait()
		close(done)
	}()
}

// WaitAny 等待任何一个任务执行完毕
func WaitAny(done chan struct{}, result chan interface{}, tasks ...func() interface{}) {
	for _, task := range tasks {
		go func(t func() interface{}) {
			res := t()
			result <- res
			close(done)
		}(task)
	}
}
