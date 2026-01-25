package main

import (
	"fmt"
	"teacher_students_management/actions"
)

func main() {
	var choice int
	for {
		fmt.Println("\n🎓 ═════ QUAN LY HOC SINH - GIAO VIEN ═════")
		fmt.Println("┌──────────────────────────────────────────┐")
		fmt.Println("│  👨‍🏫  1. Quan ly giao vien               │")
		fmt.Println("│  👨‍🎓  2. Quan ly hoc sinh                │")
		fmt.Println("│  🚪  3. Thoat chuong trinh               │")
		fmt.Println("└──────────────────────────────────────────┘")
		fmt.Print("👉 Nhap lua chon cua ban: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			actions.TeacherActions()
		case 2:
			actions.StudentActions()
		case 3:
			fmt.Println("\n👋 Cam on ban da su dung chuong trinh!")
			fmt.Println("🚪 Thoat chuong trinh...")
			return
		default:
			fmt.Println("❗ Lua chon khong hop le, vui long chon lai!")
		}
	}

}
