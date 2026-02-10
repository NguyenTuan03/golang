package main

import (
	"fmt"
	"sync"
	"time"
)

// Goroutine là 1 luồng nhẹ mà Go tạo ra để chạy 1 tác vụ nào đó đồng thời với chương trình chính
// Sử dụng concurrent: khi có các công việc nhẹ và không cần chờ đợi kết quả ngay lập tức
// Sử dụng parallel: khi có các công việc nặng và cần chờ đợi kết quả ngay lập tức
func Task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task", id, "is running")
	time.Sleep(1 * time.Second)
	fmt.Println("Task", id, "is done")
}
// wg.Add(1) để cho biết có 1 goroutine mới
// wg.Done() để cho biết goroutine đã hoàn thành
// wg.Wait() để chờ đợi tất cả goroutine hoàn thành
func Goroutine() {
	start := time.Now()
	var wg sync.WaitGroup

	for i := range 3 {
		wg.Add(1)
		go Task(i, &wg)
	}

	wg.Wait()	
	fmt.Println("Total time:", time.Since(start))
}