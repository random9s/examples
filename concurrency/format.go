package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sync/atomic"
	"time"
)

//Format does some formatting
type Format struct {
	send    chan interface{}
	receive chan interface{}

	dataType interface{}
	n        int32
}

//NewFormat ...
func NewFormat(dt interface{}, n int32) *Format {
	f := &Format{
		dataType: dt,
		send:     make(chan interface{}),
		receive:  make(chan interface{}),
	}

	if n == -1 {
		f.n = n
	} else {
		atomic.AddInt32(&f.n, +n)
	}

	return f
}

//Receive takes in data
func (f *Format) Receive() {
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

//Process formats the data
func (f *Format) Process(v interface{}) {
	go func() {
		defer Close(f)

		ptr := reflect.ValueOf(f.dataType)
		if ptr.Kind() != reflect.Ptr || ptr.IsNil() {
			log.Fatal(fmt.Errorf("cannot reflect non-pointer type %#v", ptr))
		}

		if ptr.IsValid() {
			structType := reflect.TypeOf(f.dataType).Elem()
			newStruct := reflect.New(structType)
			var newStructIface = newStruct.Interface()
			var b = reflect.ValueOf(v).Bytes()

			err := json.Unmarshal(b, newStructIface)
			if err != nil {
				log.Fatal(err)
			}

			f.Send(newStructIface)
		}
	}()
}

//Send puts out data
func (f *Format) Send(v interface{}) {
	f.send <- v
}

//Out returns the sender
func (f *Format) Out() chan interface{} {
	return f.send
}

//In sets the receiver
func (f *Format) In(in chan interface{}) {
	f.receive = in
}

//Stop shuts formatter down
func (f *Format) Stop() {
	fmt.Println("Formatter shutting down...")
	close(f.send)
}

//Atomic ...
func (f *Format) Atomic() *int32 {
	return &f.n
}

//Wait ...
func (f *Format) Wait() {
	for f.n != 0 {
		time.Sleep(time.Microsecond)
	}
}
