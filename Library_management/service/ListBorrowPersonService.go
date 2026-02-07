package service

import (
	"fmt"
	structs "library-management/struct"


	"github.com/google/uuid"
)

func ListBorrowPersonService(books map[uuid.UUID]structs.BookStruct, persons map[uuid.UUID]structs.PersonStruct) {
	for _, p := range persons {
		if p.BorrowedBooks == nil {
			fmt.Println("No books borrowed by", p.Name)
			continue
		}
		fmt.Print("Person:", p.Name, " is borrowing books: ")
		for _, bookID := range p.BorrowedBooks {
			fmt.Print(books[bookID].Name, " status is ", books[bookID].Status)
		}
		fmt.Println()
	}
}
//ea3ee9b7-9303-437b-8b00-16b296a50e33
//bfd2ce6a-6f56-413f-9863-8423d999cab0