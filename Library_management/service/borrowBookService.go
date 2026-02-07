package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	structs "library-management/struct"

	"github.com/google/uuid"
)

func BorrowBookService(books map[uuid.UUID]structs.BookStruct, persons map[uuid.UUID]structs.PersonStruct) {
	fmt.Println("Borrowing book...")
	var bookID uuid.UUID
	var personID uuid.UUID
	reader := bufio.NewReader(os.Stdin)

	// Clean unexpected buffer (optional but safe)
	// reader.ReadString('\n')

	fmt.Println("Enter book ID:")
	idStr, _ := reader.ReadString('\n')
	bookID, _ = uuid.Parse(strings.TrimSpace(idStr))

	fmt.Println("Enter person ID:")
	idStr, _ = reader.ReadString('\n')
	personID, _ = uuid.Parse(strings.TrimSpace(idStr))

	if _, ok := books[bookID]; !ok {
		fmt.Println("Book not found!")
		return
	}

	if _, ok := persons[personID]; !ok {
		fmt.Println("Person not found!")
		return
	}

	// 1. Lay sach ra
	book := books[bookID]
	person := persons[personID]

	if book.Status == structs.Borrowed {
		fmt.Println("Book is already borrowed!")
		return
	}

	// 2. Sua status
	book.Status = structs.Borrowed

	// 3. Gan nguoc lai vao map
	books[bookID] = book
	person.BorrowedBooks = append(person.BorrowedBooks, bookID)
	persons[personID] = person

	fmt.Println("Book borrowed successfully!")
}
