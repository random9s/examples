package main

import "fmt"

//Speaker implements the Speak func
type Speaker interface {
	Speak()
}

//Cat is a cat
type Cat struct{}

//Speak ...
func (c *Cat) Speak() { fmt.Println("Cat says meow or whatever") }

//Dog is a dog
type Dog struct{}

//Speak ...
func (d *Dog) Speak() { fmt.Println("Dog says woof, maybe") }

//Gorilla is a gorilla
type Gorilla struct{}

//Speak ...
func (g *Gorilla) Speak() { fmt.Println("Gorilla says arghhhgfhghgghghghghgh!") }

func main() {
	var animals = make([]Speaker, 0)
	animals = append(animals, new(Cat))
	animals = append(animals, new(Dog))
	animals = append(animals, new(Gorilla))
	animals = append(animals, new(Dog))
	animals = append(animals, new(Gorilla))
	animals = append(animals, new(Cat))

	for _, animal := range animals {
		animal.Speak()
	}
}
