package async

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {

	done := make(chan struct{})
	result := make(chan interface{})

	// 任务函数1，模拟一个耗时操作
	task1 := func() interface{} {
		time.Sleep(2 * time.Second)
		return "Task 1 completed"
	}

	// 任务函数2，模拟另一个耗时操作
	task2 := func() interface{} {
		time.Sleep(3 * time.Second)
		return "Task 2 completed"
	}

	// 启动任务
	go WaitAll(done, task1, task2)

	// 等待任何一个任务完成
	select {
	case res := <-result:
		fmt.Println("Received result:", res)
	case <-done:
		fmt.Println("All tasks are done.")
	}
}
