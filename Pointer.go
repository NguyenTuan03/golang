package main 
import "fmt"
func main() {
	var a int = 10
	var b *int = &a // b là con trỏ trỏ đến a
	fmt.Println("Trước khi gán: ", &a, b)
	*b = 20 // Gán giá trị mới cho a thông qua con trỏ b
	fmt.Println("Sau khi gán: ", a, *b)
	//=============================================================================
	//==============================================================================
	arr := []int{1, 2, 3}
	newArr := arr // newArr là một bản sao của arr
	newArr[0] = 10
	// Thay đổi giá trị của phần tử thứ hai trong mảng thông qua con trỏ	
	fmt.Println("Mảng arr1 và arr: ", newArr, arr)
}