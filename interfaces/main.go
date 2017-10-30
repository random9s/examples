package main

import (
	"fmt"
	"reflect"
)

/*
 * Outline
 * ------------------------------------------------------------------------------------------------
 * 1. What is an Interface
 *     - Interface-based programming (https://en.wikipedia.org/wiki/Interface-based_programming)
 *     - Art of Unix Programming
 *	   - Languages that implement interfaces and their purposes
 * 2  Golang Specific Interface Type
 *     - How this implementation differs from more (dynamic) languages (PHP, Ruby, Obj-C)
 *	   - Implicit implementation (You don't need to declare intent for interfaces)
 *     - val, ok ~vs~ val
 * 	   - Static - checked at compile time
 *	   - Dynamic - checked when asked for at runtime
 * 3. Creating your own
 *	   - Naming conventions
 *     - Thinking about the most abstract use
 *     - When to use them
 * 4. Description of Popular Golang Interfaces
 *	   - Reader/Writer/Closer/Seeker (Useful for reading from and writing to anything)
 *     - Encoding (useful for converting arbitrary data into Golang type)
 *	   - Context (useful for any interaction between a client server)
 *     - Sort (useful for creating a quick sort function for any data type)
 * 5. Reflection
 *	   - Basics (types, values, kind)
 *	   - Meta Tags
 *	   - Marshaling data from one obj to another
 */

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
 *  Also, worth noting: you can convert the underlying type as well as check for a successful conversion by doing val, ok := s.(*Cat) instead of val := s.(*Cat)
 *  This will cause the the ok value to be populated with a value indicating if the value contains the specified type instead of throwing an error at runtime.
 *
 */

//Which checks the speaker for the underlying type
func Which(s Speaker) string {
	var val string
	switch s.(type) {
	case *Cat:
		val = "Cat"
	case *Dog:
		val = "Dog"
	case *Gorilla:
		val = "Gorilla"
	default:
		val = "UFO"
	}
	return val
}

//Equals checks the underlying type to see if a and b are equal
func Equals(a, b Speaker) bool {
	return reflect.DeepEqual(a, b)
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
		fmt.Println("Underlying type is:", Which(animal))
		animal.Speak()
		fmt.Println()
	}

	fmt.Println("What happens if we pass in nil as type Speaker?\nUnderlying type is:", Which(nil))
	fmt.Println()
	fmt.Println("What about checking equality of two interfaces?\nCat == Gorilla:", Equals(new(Cat), new(Gorilla)))
}
