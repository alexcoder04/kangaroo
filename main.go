package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

const SIGRTMIN = 34

var (
	flagSigNum = flag.Int("signal", 1, "signal to listen on")
)

func ListenToQuit(signalNumber int) {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.Signal(signalNumber))
	select {
	case <-channel:
		fmt.Printf("caugth signal %d, exiting\n", signalNumber)
		os.Exit(0)
	}
}

func ListenFor(signalNumber int, cmd string, args []string) {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.Signal(SIGRTMIN+signalNumber))
	fmt.Printf("waiting for signal SIGRTMIN+%d to execute %s...\n", signalNumber, cmd)
	select {
	case <-channel:
		fmt.Printf("got signal, executing %s:\n", cmd)
		command := exec.Command(cmd, args...)

		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		command.Stdout = mw
		command.Stderr = mw

		command.Run()

		ListenFor(signalNumber, cmd, args)
	}
}

func main() {
	flag.Parse()

	go ListenToQuit(1)
	go ListenToQuit(2)
	go ListenToQuit(3)
	go ListenToQuit(15)

	var cmd string
	args := flag.CommandLine.Args()
	switch len(args) {
	case 0:
		cmd = "echo"
		args = []string{"hello world"}
	case 1:
		cmd = args[0]
		args = []string{}
	default:
		cmd = args[0]
		args = args[1:]
	}

	ListenFor(*flagSigNum, cmd, args)
}
