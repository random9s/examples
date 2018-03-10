package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	var dir = "tmp"
	n, err := filesInDirShouldFail(dir)
	if err != nil {
		fmt.Println("First time failed with err:", err)
	}

	n, err = filesInDirShouldSucceed(dir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Second time worked with output:", n)

	n, err = filesInDirCleanerOption(dir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Cleaner option worked with output:", n)
}

func filesInDirShouldFail(dirpath string) (int64, error) {
	//ONLY THE FIRST INPUT IN EXEC COMMAND IS TREATED AS A PROCESS. THE REST OF THE ARGUMENTS ARE TAKEN AS ARGUMENTS FOR THE PROCESS
	cmd := exec.Command("ls", dirpath, "|", "wc", "-l")
	var errBuff = bytes.Buffer{}
	var outBuff = bytes.Buffer{}
	cmd.Stderr = &errBuff
	cmd.Stdout = &outBuff

	cmd.Run()
	if errBuff.Len() > 0 {
		//THE ERROR WILL HAVE TO DO WITH |, wc, and -l not being directories.  Because they are being interpreted as arguments for ls
		return 0, errors.New(errBuff.String())
	}

	//It may be worth mentioning that the output, if we are expecting output, will not always be formatted the way we expect, so be sure to check for correct formats when converting to another type
	var b = outBuff.Bytes()
	b = b[:len(b)-1] //Remove newline string
	str := strings.TrimSpace(string(b))
	return strconv.ParseInt(str, 10, 64)
}

//This looks pretty annoying, and the more options you have the worse it gets. Why is this better or worse than just allowing a full string command?
func filesInDirShouldSucceed(dirpath string) (int64, error) {
	ls := exec.Command("ls", dirpath)
	wc := exec.Command("wc", "-l")

	//When you call os.Pipe you get two copies of a file descriptor with specific permissions
	//w (has write permissions) r (has read permissions)
	//These file descriptors are then given to the commands Stdout and Stdin, respectively.
	//The command which is reading from the file descriptor will continue to do so until the file emits an EOF
	//This is why we need to call close immediately after the ls.Wait, if we do not call Close, the wc.Wait will
	//wait indefinitely.
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

	//It may be worth mentioning that the output, if we are expecting output, will not always be formatted the way we expect, so be sure to check for correct formats when converting to another type
	b := buff.Bytes()
	b = b[:len(b)-1] //Remove newline string
	str := strings.TrimSpace(string(b))
	return strconv.ParseInt(str, 10, 64)
}

func filesInDirCleanerOption(dirpath string) (int64, error) {
	//You can either write a bash script or just write the line

	var sh = `ls ` + dirpath + ` | wc -l`

	//Output handles all Start, wait, and returns Stdin, which allows for cleaner usage when complex bash commands are required
	b, err := exec.Command("bash", "-c", sh).Output()
	if err != nil {
		return 0, err
	}

	b = b[:len(b)-1] //Remove newline string
	str := strings.TrimSpace(string(b))
	return strconv.ParseInt(str, 10, 64)
}
