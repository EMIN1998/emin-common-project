package testinterface

import "fmt"

type Animal interface {
	Roar()
}

func NewAnimal(name string) Animal {
	switch name {
	case "Cat":
		cat := newCat()
		return &cat
	default:
		return nil
	}
}

type cat struct {
}

func newCat() cat {
	return cat{}
}

func (c *cat) Roar() {
	fmt.Print("I am Cat ...")
}

type Dog struct {
	Name string
}

func (d *Dog) Roar() {
	fmt.Printf("I am Dog :%s...", d.Name)
}
