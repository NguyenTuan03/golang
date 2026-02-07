package service

import (
	"fmt"

	"library-management/struct"

	"github.com/google/uuid"
)

func AddPersonService(persons map[uuid.UUID]structs.PersonStruct) {
	fmt.Println("Adding new person...")
	var person structs.PersonStruct

	fmt.Println("Enter person name:")
	fmt.Scanln(&person.Name)

	fmt.Println("Enter person email:")
	fmt.Scanln(&person.Email)

	if person.Name == "" || person.Email == "" {
		fmt.Println("Person name and email cannot be empty!")
		return
	}

	person.ID = uuid.New()
	persons[person.ID] = person

	fmt.Println("Person added successfully!")
}
