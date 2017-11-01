package main

import "fmt"

//Fyuse ...
type Fyuse struct {
	UID  string `json:"uid"`
	Path string `json:"path"`
}

//Book data
type Book struct {
	ID   string `json:"id"`
	Info *Info  `json:"volumeInfo"`
}

//Info ...
type Info struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Subtitle    string `json:"subtitle"`
}

func (b *Book) String() string {
	var str = "<nil>"
	if b.Info != nil {
		str = fmt.Sprintf("%s", b.Info.Title)

		if b.Info.Subtitle != "" {
			str = fmt.Sprintf("%s: %s", b.Info.Title, b.Info.Subtitle)
		}
	}

	return str
}
