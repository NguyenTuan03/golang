package actions

import "fmt"

func StudentActions() {
	var studentChoice int
	for {
		fmt.Println("\n👨‍🎓 ═══════════ QUAN LY HOC SINH ═══════════")
		fmt.Println("┌────────────────────────────────────────────┐")
		fmt.Println("│  ➕  1. Them hoc sinh                      │")
		fmt.Println("│  📋  2. Hien thi danh sach hoc sinh        │")
		fmt.Println("│  ✏️  3. Cap nhat thong tin hoc sinh        │")
		fmt.Println("│  ❌  4. Xoa hoc sinh                       │")
		fmt.Println("│  🔍  5. Tim kiem hoc sinh                  │")
		fmt.Println("│  ⬅️  6. Quay lai menu chinh                │")
		fmt.Println("└────────────────────────────────────────────┘")
		fmt.Print("👉 Nhap lua chon cua ban: ")
		fmt.Scanln(&studentChoice)
		switch studentChoice {
		case 1:
			// Them hoc sinh
		case 2:
			// Hien thi danh sach hoc sinh
		case 3:
			// Cap nhat thong tin hoc sinh
		case 4:
			// Xoa hoc sinh
		case 5:
			// Tim kiem hoc sinh
		case 6:
			return
		default:
			fmt.Println("❗ Lua chon khong hop le, vui long chon lai!")
		}
	}
}
