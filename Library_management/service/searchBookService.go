package service

import (
	"fmt"
	"library-management/struct"

	"github.com/google/uuid"
)

func SearchBookService(books map[uuid.UUID]structs.BookStruct) {
	fmt.Println("Searching book...")
	fmt.Println("Enter book name:")
	var name string
	fmt.Scanln(&name)

	for _, book := range books {
		if book.Name == name {
			fmt.Println(book)
		}
	}
}