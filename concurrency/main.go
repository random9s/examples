package main

import (
	"fmt"
	"time"
)

func main() {
	var fetcher = NewFetch()
	var formatter = NewFormat(new(Book))
	Connect(fetcher, formatter)

	var incinerator = NewFurnace()
	Connect(formatter, incinerator)

	fetcher.Get("https://www.googleapis.com/books/v1/volumes/3fZa9af1KtYC")
	fetcher.Get("https://www.googleapis.com/books/v1/volumes/I6BOBAAAQBAJ")
	fetcher.Get("https://www.googleapis.com/books/v1/volumes/RDyjvJbdVvQC")
	fetcher.Get("https://www.googleapis.com/books/v1/volumes/4t-sybVuoqoC")
	fetcher.Get("https://www.googleapis.com/books/v1/volumes/9f9uAQAAQBAJ")
	fetcher.Get("https://www.googleapis.com/books/v1/volumes/DZQg43mfFPsC")
	fetcher.Get("https://www.googleapis.com/books/v1/volumes/a2Q6U0b36rMC")
	fetcher.Get("https://www.googleapis.com/books/v1/volumes/UoN_r_NMR_EC")
	fetcher.Get("https://www.googleapis.com/books/v1/volumes/XXdyQgAACAAJ")
	fetcher.Get("https://www.googleapis.com/books/v1/volumes/Rl-F95_f0GoC")

	//pseudo wait
	for {
		time.Sleep(time.Second * 5)
	}

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
