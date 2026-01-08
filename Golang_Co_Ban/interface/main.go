package main

import (
	"fmt"
	"interface/cat"
	"interface/dog"
)

// Animal là một INTERFACE - định nghĩa "hợp đồng" mà các struct phải tuân theo
// Bất kỳ type nào có đủ các method bên dưới sẽ tự động implement interface này
// KHÔNG CẦN khai báo tường minh "implements Animal" như Java/C#
type Animal interface {
	GetName() string // Method: Trả về tên động vật
	Speak() string   // Method: Trả về tiếng kêu
}

// DogInterface là EMBEDDED INTERFACE - kế thừa tất cả method từ Animal
// Và thêm method riêng của Dog
// Type implement DogInterface PHẢI có: GetName(), Speak(), Hunting()
type DogInterface interface {
	Animal           // Nhúng interface Animal (kế thừa GetName, Speak)
	Hunting() string // Method bổ sung: Chỉ Dog có
}

// MakeSound nhận vào BẤT KỲ type nào implement Animal
// Đây là POLYMORPHISM (đa hình): một function xử lý nhiều type khác nhau
func MakeSound(a Animal) {
	fmt.Println(a.GetName()) // Gọi method GetName() - mọi Animal đều có
	fmt.Println(a.Speak())   // Gọi method Speak() - mọi Animal đều có

	// TYPE ASSERTION: Kiểm tra xem Animal có phải là DogInterface không
	// Cú pháp: value, ok := interface.(Type)
	// - Nếu là DogInterface: ok = true, d = giá trị đã chuyển đổi
	// - Nếu không: ok = false, tránh panic
	if d, ok := a.(DogInterface); ok {
		fmt.Println(d.Hunting()) // Chỉ Dog mới có method Hunting()
	}

}

func main() {
	// Tạo Dog instance bằng constructor function New()
	myDog := dog.New("Buddy")
	MakeSound(myDog) // ✅ Dog implement Animal và DogInterface

	fmt.Println("====================")

	// Tạo Cat instance - có thể trả về error
	// Go convention: function trả về (value, error)
	myCat, err := cat.New("Whiskers")
	if err != nil {
		fmt.Println("Error creating cat:", err)
		return // Dừng nếu có lỗi
	}
	MakeSound(myCat) // ✅ Cat implement Animal (nhưng KHÔNG có Hunting)

	/*
	   TÓM TẮT CÁC KHÁI NIỆM:

	   1. INTERFACE: Định nghĩa tập hợp method (hợp đồng)
	      - Không có implementation, chỉ có signature
	      - Type tự động implement nếu có đủ method

	   2. METHOD vs FUNCTION:
	      - Method: func (receiver Type) MethodName() { }
	      - Function: func FunctionName() { }
	      - Interface CHỈ chấp nhận method, không chấp nhận function

	   3. EMBEDDED INTERFACE:
	      - Interface có thể nhúng interface khác
	      - Kế thừa tất cả method của interface cha

	   4. TYPE ASSERTION:
	      - Chuyển đổi interface về type cụ thể
	      - Dùng để kiểm tra type động và gọi method đặc thù
	*/
}
