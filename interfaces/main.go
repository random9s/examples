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

/*
 *
 *  Interfaces have static checking at compile time. For example:
 *      Which(0) produces error: cannot use 0 (type int) as type Speaker in argument to Which: int does not implement Speaker (missing Speak method)
 *
 *	Interfaces allow for dynamic checking at runtime!
 *      var s Speaker; val, ok := s.(*Cat); <-- Allows us to convert the interface to a concrete type
 *      s.(int) <-- produces error: impossible type switch case: s (type Speaker) cannot have dynamic type int (missing Speak method)
 *
 */

//Which checks the speaker for the underlying type
func Which(s Speaker) {
	switch s.(type) {
	case *Cat:
		fmt.Println("Underlying type is cat")
	case *Dog:
		fmt.Println("Underlying type is dog")
	case *Gorilla:
		fmt.Println("Underlying type is gorilla")
	default:
		fmt.Println("Animal not found")
	}
}

func main() {
	var animals = make([]Speaker, 0)
	animals = append(animals, new(Cat))
	animals = append(animals, new(Dog))
	animals = append(animals, new(Gorilla))
	animals = append(animals, new(Dog))
	animals = append(animals, new(Gorilla))
	animals = append(animals, new(Cat))

	for _, animal := range animals {
		Which(animal)
		animal.Speak()
	}
}
