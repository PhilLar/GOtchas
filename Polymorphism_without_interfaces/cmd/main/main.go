package main

import (
	"fmt"
)

type Animal struct {
	makeNoiseFn func(*Animal) string
	name        string
	legs        int
}

func (a *Animal) makeNoise() string {
	return a.makeNoiseFn(a)
}

func NewDog(name string) *Animal {
	return &Animal{
		makeNoiseFn: func(a *Animal) string {
			return a.name + " says woof!"
		},
		legs: 4,
		name: name,
	}
}

func NewDuck(name string) *Animal {
	return &Animal{
		makeNoiseFn: func(a *Animal) string {
			return a.name + " says quack!"
		},
		legs: 2,
		name: name,
	}
}

// CODE ABOVE SHOWS HOW YOU CAN DO IT
// CODE BELOW SHOWS HOW YOU SHOULD DO IT

//type Animal interface {
//	makeNoise() string
//}
//
//type Dog struct {
//	name string
//	legs int
//}
//
//func (d *Dog) makeNoise() string {
//	return d.name + " says woof!"
//}
//
//type Duck struct {
//	name string
//	legs int
//}
//
//func (d *Duck) makeNoise() string {
//	return d.name + " says quack!"
//}
//
//func NewDog(name string) Animal {
//	return &Dog{
//		legs: 4,
//		name: name,
//	}
//}
//
//func NewDuck(name string) Animal {
//	return &Duck{
//		legs: 4,
//		name: name,
//	}
//}

func main() {
	var dog, duck *Animal

	dog = NewDog("fido")
	duck = NewDuck("donald")

	fmt.Println(dog.makeNoise())
	// fido says woof!
	fmt.Println(duck.makeNoise())
	// donald says quack!
}
