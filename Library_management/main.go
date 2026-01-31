package main

import (
	"fmt"
	"library-management/service"
)

func main() {
	var choice int
	for {
		fmt.Println("\n🎓 ═════ CHUONG TRINH QUAN LY THU VIEN ═════")
		fmt.Println("┌──────────────────────────────────────────┐")
		fmt.Println("│  👨‍🏫  1. Thêm sách                        │")
		fmt.Println("│  👨‍🎓  2. Xem danh sách sách               │")
		fmt.Println("│  🚪  3. Thêm người mượn                  │")
		fmt.Println("│  🚪  4. Xem danh sách người mượn         │")
		fmt.Println("│  🚪  5. Mượn sách                        │")
		fmt.Println("│  🚪  6. Xem lịch sử mượn                 │")
		fmt.Println("│  🚪  7. Trả sách                         │")
		fmt.Println("│  🚪  8. Tìm kiếm sách                    │")
		fmt.Println("│  🚪  9. Thoát chương trình               │")
		fmt.Println("└──────────────────────────────────────────┘")
		fmt.Print("👉 Nhap lua chon cua ban: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			service.AddBookService()
		case 2:
			
		case 3:
			fmt.Println("\n👋 Cam on ban da su dung chuong trinh!")
			fmt.Println("🚪 Thoat chuong trinh...")
			return
		default:
			fmt.Println("❗ Lua chon khong hop le, vui long chon lai!")
		}
	}
}
