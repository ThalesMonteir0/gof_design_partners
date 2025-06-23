package prototype

import "fmt"

type Node interface {
	Print(indentation string)
	Clone() Node
}

type Contact struct {
	Name string
	Cel  string
}

func NewContact(name, cel string) Node {
	return &Contact{
		name,
		cel,
	}
}

func (c Contact) Print(indentation string) {
	fmt.Println(indentation + c.Name + " - " + c.Cel)
}

func (c Contact) Clone() Node {
	return NewContact(c.Name+"_clone", c.Cel)
}
