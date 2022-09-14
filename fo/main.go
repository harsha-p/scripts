package main

import (
	// "fmt"
	"os"
	"os/exec"
	"strings"
	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
	"syscall"
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
	// open file in zathura as background process
	cmd = exec.Command("xdg-open", file)
	// start and do not wait for process to finish
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Start()
}