package actions

import (
	"fmt"
	"teacher_students_management/services/teacherService"
)

func TeacherActions() {
	var teacherChoice int
	arrTeachers := []teacherservice.TeacherStruct{}

	for {
		fmt.Println("\n👨‍🏫 ═══════════ QUAN LY GIAO VIEN ═══════════")
		fmt.Println("┌────────────────────────────────────────────┐")
		fmt.Println("│  ➕  1. Them giao vien                     │")
		fmt.Println("│  📋  2. Hien thi danh sach giao vien       │")
		fmt.Println("│  ✏️  3. Cap nhat thong tin giao vien       │")
		fmt.Println("│  ❌  4. Xoa giao vien                      │")
		fmt.Println("│  🔍  5. Tim kiem giao vien                 │")
		fmt.Println("│  ⬅️  6. Quay lai menu chinh                │")
		fmt.Println("└────────────────────────────────────────────┘")
		fmt.Print("👉 Nhap lua chon cua ban: ")
		fmt.Scanln(&teacherChoice)
		switch teacherChoice {
		case 1:
			teacherservice.AddNewTeacher(&arrTeachers)
		case 2:
			teacherservice.GetTeacher(&arrTeachers)
		case 3:
			teacherservice.UpdateTeacher(&arrTeachers)
		case 4:
			teacherservice.DeleteTeacher(&arrTeachers)
		case 5:
			var id int
			fmt.Scanln(&id)
			teacherservice.FindTeacherByID(&arrTeachers, id)
		case 6:
			return
		default:
			fmt.Println("❗ Lua chon khong hop le, vui long chon lai!")
		}
	}

}
