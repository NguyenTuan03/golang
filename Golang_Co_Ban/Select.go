package main

import (
	"fmt"
	"time"
)

func Select() {
	// Select là 1 cơ chế để chọn nhiều channel để nhận dữ liệu
	// Các loại Select
	// Unbuffered Channel: Chỉ chứa được 1 gtri tại 1 thời điểm
	// Buffered Channel: Chứa được nhiều gtri tại 1 thời điểm
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()
	for i := range 2 {
		fmt.Println(i)
		select {
		case msg1 := <-ch1:
			fmt.Println("received", msg1)
		case msg2 := <-ch2:
			fmt.Println("received", msg2)
		}
	}

}
