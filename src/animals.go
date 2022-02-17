package main

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) WhoAmI() {
	fmt.Printf("Hi, I am a %s!\n", a.Name)
}

type Cat struct {
	Animal
}

func (c *Cat) SayHi() {
	fmt.Println("Meow~")
}

type Dog struct {
	Animal
}

func (d *Dog) SayHi() {
	fmt.Println("Woof~")
}

func NewDog() *Dog {
	return &Dog{
		Animal{Name: "dog"},
	}
}

func main() {
	c := &Cat{
		Animal: Animal{
			Name: "cat",
		},
	}
	c.WhoAmI()
	c.SayHi()

	d := NewDog()
	d.WhoAmI()
	d.SayHi()

}
