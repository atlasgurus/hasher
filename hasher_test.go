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
}

// NewPerson creates a new Person with a Hasher.
func NewPerson(name string, age int, parent *Person) *Person {
	p := &Person{Name: name, Age: age, Parent: parent}
	return p
}

func TestComputeHash(t *testing.T) {
	parent := NewPerson("Jane Doe", 50, nil)
	person := NewPerson("John Doe", 30, parent)

	parentHash := ComputeHash(parent)
	personHash := ComputeHash(person)

	if fmt.Sprintf("%x", parentHash) != "01332c876518a793b7c1b8dfaf6d4b404ff5db09b21c6627ca59710cc24f696a" {
		t.Errorf("Unexpected value for parentHash")
	}
	if fmt.Sprintf("%x", personHash) != "2c34e78797862b4012fea78fb69a81206fc236786aea7932ce36df43788cfde5" {
		t.Errorf("Unexpected value for personHash")
	}
}
