package dog

type Dog struct {
	Name string `json:"name"`
}
func New(name string) *Dog {
	return &Dog{Name: name}
}
func (d *Dog) Speak() string {
	return "Bark!"
}
func (d *Dog) GetName() string {
	return d.Name
}
func (d *Dog) Hunting() string {
	return d.Name + " is hunting!"
}