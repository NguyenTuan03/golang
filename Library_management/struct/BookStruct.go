package structs

import "github.com/google/uuid"

type BookStatus int

const (
	Available BookStatus = iota
	Borrowed
)

func (s BookStatus) String() string {
	switch s {
	case Available:
		return "Available"
	case Borrowed:
		return "Borrowed"
	default:
		return "Unknown"
	}
}

type BookStruct struct {
	ID     uuid.UUID
	Name   string
	Author string
	Status BookStatus
}
