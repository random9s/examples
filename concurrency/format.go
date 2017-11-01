package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

//Format does some formatting
type Format struct {
	dataType interface{}
	send     chan interface{}
	receive  chan interface{}
}

//NewFormat ...
func NewFormat(dt interface{}) *Format {
	return &Format{
		dataType: dt,
		send:     make(chan interface{}),
		receive:  make(chan interface{}),
	}
}

//Out returns the sender
func (f *Format) Out() chan interface{} {
	return f.send
}

//In sets the receiver
func (f *Format) In(in chan interface{}) {
	f.receive = in
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
