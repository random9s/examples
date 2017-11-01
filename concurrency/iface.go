package main

import (
	"sync/atomic"
)

//Sender implements a send channel
type Sender interface {
	Send(interface{})
	Out() chan interface{}
}

//Receiver implements a receive channel
type Receiver interface {
	Receive()
	In(chan interface{})
}

//Processor assigns work to worker
type Processor interface {
	Process(interface{})
}

//Connect bridges two structs
func Connect(s Sender, r Receiver) {
	r.In(s.Out())
	r.Receive()
}

//Stopper is responsible for stopping a worker
type Stopper interface {
	Atomic() *int32
	Stop()
	Wait()
}

//Wait checks for one or more stoppers to complete
func Wait(s ...Stopper) {
	for _, st := range s {
		st.Wait()
	}
}

//Close checks for all work to be completed, then calls stop
func Close(s Stopper) {
	//If -1 is set, go forever
	var n = s.Atomic()

	if *n == -1 {
		return
	}

	if atomic.AddInt32(n, -1) == 0 {
		s.Stop()
	}
}
