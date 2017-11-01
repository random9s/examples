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
			case v, ok := <-f.receive:
				if !ok {
					return
				}

				if fyu, ok := v.(*Fyuse); ok {
					fmt.Println("Incinerated fyuse", fyu.UID, " with path", fyu.Path)
					FINISH <- fyu.UID
				}
				/*
					if book, ok := v.(*Book); ok {
						fmt.Println("Incinerated book", book)
					}
				*/
			}
		}
	}()
}
