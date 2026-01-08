package cat

import (
	"errors"
	"strings"
)

type Cat struct {
	Name string `json:"name"`
}
func New(name string) (*Cat, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return  nil, errors.New("Name cannot be empty");
	}
	if len(name) > 50 {
		return nil, errors.New("Name cannot be longer than 50 characters")
	}
	return &Cat{Name: name}, nil
}
func (c *Cat) Speak() string {
	return "Meow!"
}
func (c *Cat) GetName() string {
	return c.Name
}