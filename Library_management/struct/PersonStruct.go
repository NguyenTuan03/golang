package structs

import "github.com/google/uuid"

type PersonStruct struct {
	ID    uuid.UUID
	Name  string
	Email string
	BorrowedBooks []uuid.UUID
}
