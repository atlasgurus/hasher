package hasher

import (
	"fmt"
	"testing"
)

// Person struct embedding the Hasher.
type Person struct {
	Name   string
	Age    int `hash:"-"`
	Parent *Person
	Hash   Hash
}

// NewPerson creates a new Person with a Hasher.
func NewPerson(name string, age int, parent *Person) *Person {
	p := &Person{Name: name, Age: age, Parent: parent}
	p.Hash = ComputeHash(p)
	return p
}

func TestComputeHash(t *testing.T) {
	parent := NewPerson("Jane Doe", 50, nil)
	person := NewPerson("John Doe", 30, parent)

	parentHash := ComputeHash(parent)
	personHash := ComputeHash(person)

	if fmt.Sprintf("%x", parentHash) != fmt.Sprintf("%x", parent.Hash) {
		t.Errorf("Unexpected value for parentHash")
	}
	if fmt.Sprintf("%x", personHash) != fmt.Sprintf("%x", person.Hash) {
		t.Errorf("Unexpected value for personHash")
	}
	if fmt.Sprintf("%x", parentHash) != "2d6101098438c11b4eaacc32d13b26d0b7ef1670037fcaf5742541b9f0d375f6" {
		t.Errorf("Unexpected value for parentHash")
	}
	if fmt.Sprintf("%x", personHash) != "b0f306193dddc919d3d1c4592ff5e3d9d20f5d2d804ad1b3f743eaa5c62df5d5" {
		t.Errorf("Unexpected value for personHash")
	}
}
