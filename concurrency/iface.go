package main

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
