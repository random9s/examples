package main

import (
	"fmt"
)

//Furnace burns data
type Furnace struct {
	receive chan interface{}
}

//NewFurnace ...
func NewFurnace() *Furnace {
	return &Furnace{
		receive: make(chan interface{}),
	}
}

//In sets receiver
func (f *Furnace) In(in chan interface{}) {
	f.receive = in
}

//Receive takes in data
func (f *Furnace) Receive() {
	go func() {
		for {
			select {
			case v := <-f.receive:
				if book, ok := v.(*Book); ok {
					fmt.Println("Incinerated book", book)
				}
			}
		}
	}()
}
