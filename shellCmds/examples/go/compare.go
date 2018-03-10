package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	//SHELL COMMANDS ARE ACTUALLY PRETTY SLOW.
	//USING THE SAME EXAMPLE WE CAN CREATE TWO SIMILAR FUNCTIONS. ONE USES GO FUNCS AND THE OTHER CALLS A SHELL COMMAND
	//IF WE TIME BOTH, WE CAN SEE ON AVERAGE THAT THE WALKDIR FUNC IS ABOUT 85x	FASTER THAN CALLING THE SHELL COMMAND
	var dir = "tmp"

	var t1 = time.Now()
	n, err := walkDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Filewalk counted %d items in dir %s. Took: %s\n", n, dir, time.Since(t1))

	var t2 = time.Now()
	n, err = filesInDir1(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Shell command 1 counted %d items in dir %s. Took: %s\n", n, dir, time.Since(t2))

	var t3 = time.Now()
	n, err = filesInDir2(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Shell command 2 counted %d items in dir %s. Took: %s\n", n, dir, time.Since(t3))
}

func walkDir(dirpath string) (int64, error) {
	var count int64
	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if dirpath == path {
			return nil
		}

		count++
		return nil
	})
	return count, err
}

func filesInDir1(dirpath string) (int64, error) {
	ls := exec.Command("ls", dirpath)
	wc := exec.Command("wc", "-l")

	r, w, err := os.Pipe()
	if err != nil {
		return 0, err
	}

	var buff = bytes.Buffer{}
	ls.Stdout = w
	wc.Stdin = r
	wc.Stdout = &buff

	ls.Start()
	wc.Start()
	ls.Wait()
	w.Close()
	wc.Wait()

	b := buff.Bytes()
	b = b[:len(b)-1]
	str := strings.TrimSpace(string(b))
	return strconv.ParseInt(str, 10, 64)
}

func filesInDir2(dirpath string) (int64, error) {
	var sh = `ls ` + dirpath + ` | wc -l`

	b, err := exec.Command("bash", "-c", sh).Output()
	if err != nil {
		return 0, err
	}

	b = b[:len(b)-1]
	str := strings.TrimSpace(string(b))
	return strconv.ParseInt(str, 10, 64)
}
