package main

import (
	// "fmt"
	"os"
	"os/exec"
	"strings"
	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
	// "syscall"
)

func main() {
	home := os.Getenv("HOME")
	cmd := exec.Command("fd", ".", home)
	out, err := cmd.Output()
	if err != nil {
		panic(err)
		return
	}
	// run fuzzyfinder on output
	files := strings.Split(string(out), "\n")
	files = files[:len(files)-1]

	idx, err := fuzzyfinder.Find(
		files,
		func(i int) string {
			return files[i]
		},
	)
	if err != nil {
		panic(err)
		return
	}
	file:=files[idx]
	// copy the file path to clipboard
	cmd = exec.Command("xclip", "-selection", "clipboard")
	cmd.Stdin = strings.NewReader(file)
	err = cmd.Run()
	if err != nil {
		panic(err)
		return
	}
	
}