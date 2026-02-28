package main

import (
	"fmt"
	"sync"
)

func RaceConditionGoRoutine() {
	tokens := 0
	var wg sync.WaitGroup
	var mtx sync.Mutex

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			mtx.Lock()
			tokens++
			mtx.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(tokens)
}
