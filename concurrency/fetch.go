package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

//Fetch gets data from some source
type Fetch struct {
	send    chan interface{}
	n       int32
	stopped bool
}

//NewFetch shuts down after n iterations
func NewFetch(n int32) *Fetch {
	f := &Fetch{
		send: make(chan interface{}),
	}

	if n == -1 {
		f.n = n
	} else {
		atomic.AddInt32(&f.n, +n)
	}

	return f
}

//Get pulls info from http source
func (f *Fetch) Get(url string) error {
	if f.stopped {
		return fmt.Errorf("system stopped")
	}

	go func(url string) {
		defer Close(f)

		resp, err := http.Get(url)
		defer resp.Body.Close()

		if err != nil {
			log.Fatal(err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			log.Fatal(resp.Status)
			return
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return
		}

		f.Send(b)
	}(url)

	return nil
}

//Send moves data to the next part
func (f *Fetch) Send(v interface{}) {
	f.send <- v
}

//Out returns sender
func (f *Fetch) Out() chan interface{} {
	return f.send
}

//Stop shuts fetcher down
func (f *Fetch) Stop() {
	fmt.Println("Fetcher shutting down...")
	f.stopped = true
	close(f.send)
}

//Atomic ...
func (f *Fetch) Atomic() *int32 {
	return &f.n
}

//Wait ...
func (f *Fetch) Wait() {
	for f.n != 0 {
		time.Sleep(time.Microsecond)
	}
}
