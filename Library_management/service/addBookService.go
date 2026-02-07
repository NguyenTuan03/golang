package service

import (
	"fmt"
	structs "library-management/struct"

	"github.com/google/uuid"
)

func AddBookService(books map[uuid.UUID]structs.BookStruct) {
	fmt.Println("Adding new book...")
	var book structs.BookStruct

	fmt.Println("Enter book title:")
	fmt.Scanln(&book.Name)

	fmt.Println("Enter book author:")
	fmt.Scanln(&book.Author)

	if book.Name == "" || book.Author == "" {
		fmt.Println("Book title and author cannot be empty!")
		return
	}

	book.ID = uuid.New()
	book.Status = structs.Available
	books[book.ID] = book

	fmt.Println("Book added successfully!")
}
