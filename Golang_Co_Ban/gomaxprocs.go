package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func heavyTask(wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := 0; i <= 100e8; i++ {
		sum += i
	}
}
func main() {
	cpu := runtime.NumCPU()
	fmt.Println("Number of CPU: ", cpu)

	runtime.GOMAXPROCS(cpu)

	var wg sync.WaitGroup
	start := time.Now()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go heavyTask(&wg)
	}

	wg.Wait()
	fmt.Println(time.Since(start))

}
