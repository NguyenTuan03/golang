package main

import "fmt"
type User struct {
	name string
	phone_number string
	email string
	address string
	age int
}
func hienthithongtin(user *User) {
	fmt.Println("Name: ", user.name)
	fmt.Println("Phone Numb`er: ", user.phone_number)
	fmt.Println("Email: ", user.email)
	fmt.Println("Address: ", user.address)
	fmt.Println("Age: ", user.age)
}
func main() {
   user := User{
	name: "John",
	phone_number: "0909090909",
	email: "john@gmail.com",
	address: "1234567890",
	age: 20,
   }
   hienthithongtin(&user)
   fmt.Println("================================================")
   anotherUser := User{
	name: "Jane",
	phone_number: "0909090909",
	email: "jane@gmail.com",
	address: "1234567890",
	age: 20,
   }
   hienthithongtin(&anotherUser)
   // Add receiver in struct func
   
}