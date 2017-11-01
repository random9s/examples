package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//Fetch gets data from some source
type Fetch struct {
	send chan interface{}
}

//NewFetch ...
func NewFetch() *Fetch {
	return &Fetch{
		send: make(chan interface{}),
	}
}

//Out returns sender
func (f *Fetch) Out() chan interface{} {
	return f.send
}

//Get pulls info from http source
func (f *Fetch) Get(url string) error {
	go func(url string) {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		//retry until 200, may cause problems
		if resp.StatusCode != http.StatusOK {
			fmt.Println("response:", resp.Status)
			time.Sleep(500 * time.Nanosecond)
			f.Get(url)
			return
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		f.Send(b)
	}(url)

	return nil
}

//Send moves data to the next part
func (f *Fetch) Send(v interface{}) {
	f.send <- v
}
