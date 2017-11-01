package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

//Furnace burns data
type Furnace struct {
	receive chan interface{}
	n       int32
}

//NewFurnace ...
func NewFurnace(n int32) *Furnace {
	f := &Furnace{
		receive: make(chan interface{}),
	}

	if n == -1 {
		f.n = n
	} else {
		atomic.AddInt32(&f.n, +n)
	}

	return f
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

				f.Process(v)
			}
		}
	}()
}

//Process prints data
func (f *Furnace) Process(v interface{}) {
	go func() {
		defer Close(f)

		if fyu, ok := v.(*Fyuse); ok {
			fmt.Println("Incinerated fyuse", fyu.UID, " with path", fyu.Path)

			//Used for timing...
			FINISH <- fyu.UID
		}
	}()
}

//Stop ...
func (f *Furnace) Stop() {
	fmt.Println("Furnace shutting down...")
}

//Atomic ...
func (f *Furnace) Atomic() *int32 {
	return &f.n
}

//Wait ...
func (f *Furnace) Wait() {
	for f.n != 0 {
		time.Sleep(time.Microsecond)
	}
}
