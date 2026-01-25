package teacherservice

import (
	"fmt"
	"slices"
)

type TeacherStruct struct {
	ID      int
	Name    string  `json:"name"`
	Subject string  `json:"subject"`
	Salary  float64 `json:"salary"`
}

func GetTeacher(arrTeachers *[]TeacherStruct) {
	if len(*arrTeachers) == 0 {
		fmt.Println("📋 Danh sach giao vien trong!")
		return
	}

	fmt.Println("\n📚 ========== DANH SACH GIAO VIEN ==========")
	for i, teacher := range *arrTeachers {
		fmt.Printf("\n👨‍🏫 Giao vien #%d\n", i+1)
		fmt.Println("─────────────────────────────────────────")
		fmt.Printf("  🆔 ID:         %d\n", teacher.ID)
		fmt.Printf("  📝 Ho Ten:     %s\n", teacher.Name)
		fmt.Printf("  📖 Mon Hoc:    %s\n", teacher.Subject)
		fmt.Printf("  💰 Luong:      %.2f VND\n", teacher.Salary)
		fmt.Println("─────────────────────────────────────────")
	}
	fmt.Printf("\n📊 Tong so giao vien: %d\n\n", len(*arrTeachers))
}

func AddNewTeacher(arrTeachers *[]TeacherStruct) {
	fmt.Println("Adding new teacher...")
	var teacher *TeacherStruct = &TeacherStruct{}
	fmt.Print("Enter teacher name: ")
	fmt.Scanln(&teacher.Name)
	fmt.Print("Enter teacher subject: ")
	fmt.Scanln(&teacher.Subject)
	fmt.Print("Enter teacher salary: ")
	fmt.Scanln(&teacher.Salary)
	id := len(*arrTeachers) + 1
	newTeacher := TeacherStruct{
		ID:      id,
		Name:    teacher.Name,
		Subject: teacher.Subject,
		Salary:  teacher.Salary,
	}
	*arrTeachers = append(*arrTeachers, newTeacher)
	fmt.Printf("New teacher %s added successfully!\n", teacher.Name)
}

func UpdateTeacher(arrTeachers *[]TeacherStruct) {
	fmt.Println("Press the ID you want to update...")
	var id int
	fmt.Scanln(&id)

	if id < 1 {
		fmt.Println("ID Không hợp lệ")
		return
	}

	idx := GetTeacherIndexByID(arrTeachers, id)
	if idx == -1 {
		return
	}
	var choice int
	for {
		fmt.Println("\n🎓 ═════ CHỌN FIELD BẠN MUỐN UPDATE ═════")
		fmt.Println("┌─────────────────────────────────────────┐")
		fmt.Println("│  👨‍🏫  1. Tên                             │")
		fmt.Println("│  👨‍🎓  2. Môn học                         │")
		fmt.Println("│  👨‍🎓  3. Lương                           │")
		fmt.Println("│  🚪  4. Thoat chuong trinh              │")
		fmt.Println("└─────────────────────────────────────────┘")
		fmt.Print("👉 Nhap lua chon cua ban: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			UpdateTeacherField("Name", arrTeachers, "Name", idx)
		case 2:
			UpdateTeacherField("Subject", arrTeachers, "Subject", idx)
		case 3:
			UpdateTeacherField("Salary", arrTeachers, "Salary", idx)
		case 4:
			fmt.Println("\n👋 Cam on ban da su dung chuong trinh!")
			fmt.Println("🚪 Thoat chuong trinh...")
			return
		default:
			fmt.Println("❗ Lua chon khong hop le, vui long chon lai!")
		}
	}
}

func UpdateTeacherField(text string, arrTeachers *[]TeacherStruct, field string, idx int) {
	fmt.Print("Enter new teacher " + text + " to update: ")
	switch field {
	case "Name":
		var name string
		fmt.Scanln(&name)
		(*arrTeachers)[idx].Name = name
	case "Subject":
		var subject string
		fmt.Scanln(&subject)
		(*arrTeachers)[idx].Subject = subject
	case "Salary":
		var salary float64
		fmt.Scanln(&salary)
		(*arrTeachers)[idx].Salary = salary
	}

}

func DeleteTeacher(arrTeachers *[]TeacherStruct) {
	fmt.Print("Enter the ID you want to delete: ")
	var id int
	fmt.Scanln(&id)

	idx := GetTeacherIndexByID(arrTeachers, id)
	if idx == -1 {
		return
	}
	*arrTeachers = slices.Delete(*arrTeachers, idx, idx+1)
	fmt.Println("Delete successful!!")
}

func FindTeacherByID(arrTeachers *[]TeacherStruct, id int) {
	idx := GetTeacherIndexByID(arrTeachers, id)
	if idx == -1 {
		return
	}
	fmt.Println("Teacher you find is: ", (*arrTeachers)[idx])
}

func GetTeacherIndexByID(arrTeachers *[]TeacherStruct, id int) int {
	idx := -1
	for i, teacher := range *arrTeachers {
		if teacher.ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println("Cannot find the teacher you want")
		return -1
	}
	return idx
}
