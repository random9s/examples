package main

import (
	"flag"
	"fmt"
	"runtime"
	"strings"
	"sync/atomic"
	"time"
)

/* Things covered:
 * - interfaces
 * - reflection
 * - channels
 * - concurrency
 * - parallelism
 */

/* Missing things:
 * - Error handling
 */

//SLOWMOBIUS ...
var (
	SLOWMOBIUS = make(chan *Track)
	FINISH     = make(chan string)
	TRACKER    = make(map[string]*Track)
	TOTDUR     time.Duration

	FYU = []string{
		"https://fyu.se/embed/7mgtv42byx",
		"https://fyu.se/embed/4c326r25cz",
		"https://fyu.se/embed/82wet0pjcy",
		"https://fyu.se/embed/fduon79iw9",
		"https://fyu.se/embed/g6fm2y7fi0",
		"https://fyu.se/embed/5vqu2es4m2",
		"https://fyu.se/embed/44n6iomk5i",
		"https://fyu.se/embed/nmfevehdmi",
		"https://fyu.se/embed/oarvq776pv",
		"https://fyu.se/embed/l1d3pa1678",
		"https://fyu.se/embed/upxiojwsyy",
		"https://fyu.se/embed/hnj2x93lte",
		"https://fyu.se/embed/eejvb7brwl",
		"https://fyu.se/embed/hba8og48bo",
		"https://fyu.se/embed/lwi46rfkuj",
		"https://fyu.se/embed/zklfkkxyc5",
		"https://fyu.se/embed/xkpj1t41ls",
		"https://fyu.se/embed/99mkv9d0jk",
		"https://fyu.se/embed/pgza17fviv",
		"https://fyu.se/embed/xevaodtx37",
		"https://fyu.se/embed/vmjr1mi8tg",
		"https://fyu.se/embed/s5ynvypx06",
		"https://fyu.se/embed/sbifgibcj9",
		"https://fyu.se/embed/1xaz279fal",
		"https://fyu.se/embed/4tx0np7dbb",
		"https://fyu.se/embed/ir8k8sr19a",
		"https://fyu.se/embed/fgzt12a3ke",
		"https://fyu.se/embed/qpumc2ut9x",
	}

	//BOOKS ...
	BOOKS = []string{
		"https://www.googleapis.com/books/v1/volumes/3fZa9af1KtYC",
		"https://www.googleapis.com/books/v1/volumes/I6BOBAAAQBAJ",
		"https://www.googleapis.com/books/v1/volumes/RDyjvJbdVvQC",
		"https://www.googleapis.com/books/v1/volumes/4t-sybVuoqoC",
		"https://www.googleapis.com/books/v1/volumes/9f9uAQAAQBAJ",
		"https://www.googleapis.com/books/v1/volumes/DZQg43mfFPsC",
		"https://www.googleapis.com/books/v1/volumes/a2Q6U0b36rMC",
		"https://www.googleapis.com/books/v1/volumes/UoN_r_NMR_EC",
		"https://www.googleapis.com/books/v1/volumes/XXdyQgAACAAJ",
		"https://www.googleapis.com/books/v1/volumes/Rl-F95_f0GoC",
	}
)

//Track ...
type Track struct {
	UID string
	Now time.Time
}

//NewTracker ...
func NewTracker(uid string) *Track {
	return &Track{
		UID: uid,
		Now: time.Now(),
	}
}

func timer() {
	go func() {
		var i int64
		for {
			select {
			case t := <-SLOWMOBIUS:
				TRACKER[t.UID] = t
			case uid := <-FINISH:
				if t, ok := TRACKER[uid]; ok {
					var dur = time.Since(t.Now)
					fmt.Println("Proc took", dur, "to complete")
					i++

					TOTDUR += dur
					avg := fmt.Sprintf("%dns", (TOTDUR.Nanoseconds() / i))
					dur, _ = time.ParseDuration(avg)
					fmt.Println("Avg duration of download:", dur, "Download count:", i, "\n")
				}
			}
		}
	}()
}

func main() {
	var parallel bool
	flag.BoolVar(&parallel, "parallel", false, "Switch to parallel mode")
	flag.Parse()

	//Start timer
	timer()

	//Run in parallel
	if parallel {
		parallelMain()
		return
	}

	var n = int32(len(FYU))
	fetch, format, burn := system(n)
	for _, fyu := range FYU {
		SLOWMOBIUS <- NewTracker(strings.TrimPrefix(fyu, "https://fyu.se/embed/"))
		fetch.Get(fyu)
	}

	Wait(fetch, format, burn)
}

func parallelMain() {
	//Find core, work block size, and remainder that will be spread across cores arbitrarily
	var cores = runtime.NumCPU()
	var blocksize = getWorkLoad(cores, len(FYU))
	var remainder = len(FYU) % cores
	var coreDone = make(chan bool)
	var n int32

	var assigned = 0
	for i := 0; i < cores; i++ {
		atomic.AddInt32(&n, +1)
		//Get section of work for core i
		var work []string
		var offset = assigned + blocksize

		//Assign work with additional offset, if needed.
		work = FYU[assigned:offset]
		if remainder > 0 {
			work = FYU[assigned : offset+1]

			remainder--
			assigned++
		}
		assigned += blocksize

		//Async start proc that creates fetch system and fetches all work
		go func(w []string) {
			var n = int32(len(w))
			fetch, format, burn := system(n)
			for _, url := range w {
				SLOWMOBIUS <- NewTracker(strings.TrimPrefix(url, "https://fyu.se/embed/"))
				fetch.Get(url)
			}

			Wait(fetch, format, burn)
			coreDone <- true
		}(work)
	}

	for range coreDone {
		if atomic.AddInt32(&n, -1) == 0 {
			break
		}
	}
}

func getWorkLoad(cores, worklen int) int {
	if cores < worklen {
		return worklen / cores
	}

	return -1
}

func altParallelMain() {
	//Find core
	var cores = runtime.NumCPU()

	ch := make(chan string)
	for i := 0; i < cores; i++ {
		go run(i, ch)
	}

	for _, fyu := range FYU {
		ch <- fyu
	}

	//Because we don't know what went where, we cannot use our wait method
	for {
		time.Sleep(time.Minute)
	}

}

func run(id int, ch chan string) {
	fetch, _, _ := system(-1)
	for {
		select {
		case url := <-ch:
			SLOWMOBIUS <- NewTracker(strings.TrimPrefix(url, "https://fyu.se/embed/"))
			fetch.Get(url)
		}
	}
}

func system(n int32) (*Fetch, *Format, *Furnace) {
	var fetcher = NewFetch(n)
	var formatter = NewFormat(new(Fyuse), n)
	Connect(fetcher, formatter)

	var incinerator = NewFurnace(n)
	Connect(formatter, incinerator)

	return fetcher, formatter, incinerator
}
