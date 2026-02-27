package main

import (
	"context"
	"fmt"
	"time"
)

func CookPho(ctx context.Context, chPho chan<- string) {
	fmt.Println("Bat dau nau pho")
	select {
	case <-time.After(1 * time.Second):
		chPho <- "pho"
	case <-ctx.Done():
		fmt.Println("Done")
		return
	}
}

func CookPizza(ctx context.Context, chPizza chan<- string) {
	fmt.Println("Bat dau nau pizza")
	select {
	case <-time.After(2 * time.Second):
		chPizza <- "pizza da xong !!"
	case <-ctx.Done():
		fmt.Println("Done")
		return
	}
}

func main() {
	// Context là 1 cơ chế giúp quản lí và kiểm soát thời gian sống cũng như hàng vi của các tác vụ trong chương trình
	// đặc biệt khi làm việc với GoRoutine. Chúng ta có thể hình dung Context như 1 bộ điêu khiển từ xa cho phép:
	// - Hủy bỏ các tác vụ
	// - Hạn chế thời gian sống của các tác vụ
	// - Truyền thông tin giữa các tác vụ
	chPho := make(chan string)
	chPizza := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	go CookPho(ctx, chPho)
	go CookPizza(ctx, chPizza)

	for i := 1; i <= 2; i++ {
		select {
		case pho := <-chPho:
			fmt.Println("Nhan duoc ", pho)
		case pizza := <-chPizza:
			fmt.Println("Nhan duoc ", pizza)
		case <-ctx.Done():
			fmt.Println("CLOSED !!!")
		}

	}

}
