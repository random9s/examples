package main

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	//User input to a shell command
	//Let's say we ask a user for a directory and we'll return the number of objects in that directory
	var cmd = "ls tmp; rm -rf target" //harmful attack
	_, err := dontHurtMe(cmd)
	if err != nil {
		log.Fatal(err)
	}
}

func dontHurtMe(cmdstr string) (interface{}, error) {
	fields := strings.Fields(cmdstr)
	bin := fields[0]
	args := fields[1:]

	cmd := exec.Command(bin, args...)
	var errBuff = bytes.Buffer{}
	var outBuff = bytes.Buffer{}
	cmd.Stderr = &errBuff
	cmd.Stdout = &outBuff

	cmd.Run()
	if errBuff.Len() > 0 {
		//The user cannot input this sort of command
		return nil, errors.New(errBuff.String())
	}

	var b = outBuff.Bytes()
	b = b[:len(b)-1]
	str := strings.TrimSpace(string(b))
	return strconv.ParseInt(str, 10, 64)
}
