package main

import (
	"fmt"
	"library-management/service"
	structs "library-management/struct"

	"github.com/google/uuid"
)

func main() {
	var choice int
	var books = make(map[uuid.UUID]structs.BookStruct)
	var persons = make(map[uuid.UUID]structs.PersonStruct)

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
			service.AddBookService(books)
		case 2:
			service.ListBookService(books)
		case 3:
			service.AddPersonService(persons)
		case 4:
			service.ListPersonService(persons)
		case 5:
			service.BorrowBookService(books, persons)
		case 6:
			service.ListBorrowPersonService(books, persons)
		case 7:
			service.ReturnBookService(books, persons)
		case 8:
			service.SearchBookService(books)
		case 9:
			fmt.Println("\n👋 Cam on ban da su dung chuong trinh!")
			fmt.Println("🚪 Thoat chuong trinh...")
			return
		default:
			fmt.Println("❗ Lua chon khong hop le, vui long chon lai!")
		}
	}
}
