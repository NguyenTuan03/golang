package main

import (
	"fmt"
	"sync"
	"time"
)

func goroutine() {
	var wg = sync.WaitGroup{}
	var m = sync.RWMutex{}

	wg.Add(2)
	go printMessage(&wg, &m, "tuan")
	go printMessage(&wg, &m, "Co")

	wg.Wait()
}

func printMessage(wg *sync.WaitGroup, m *sync.RWMutex, msg string) {
	defer wg.Done()
	for i := 0; i < len(msg); i++ {
		m.Lock()
		fmt.Println(msg, i)
		m.Unlock()
		time.Sleep(time.Second)
	}
}
