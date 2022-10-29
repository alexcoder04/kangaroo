package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alexcoder04/friendly/v2"
)

const SIGRTMIN = 34

var (
	flagSigNum   = flag.Int("signal", 1, "signal to listen on")
	flagInterval = flag.Int("interval", 0, "run command every so much seconds")
)

func ListenFor(signalNumber int, cmd []string) {
	if signalNumber == 0 {
		fmt.Println("signal number is 0, not listening")
		return
	}
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.Signal(SIGRTMIN+signalNumber))
	fmt.Printf("waiting for signal SIGRTMIN+%d to execute %s...\n", signalNumber, cmd[0])
	<-channel
	fmt.Printf("got signal, executing %s:\n", cmd[0])
	friendly.Run(cmd, "")
	ListenFor(signalNumber, cmd)
}

func ListenToQuit(signalNumber int) {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.Signal(signalNumber))
	<-channel
	fmt.Printf("caugth signal %d, exiting\n", signalNumber)
	os.Exit(0)
}

func main() {
	flag.Parse()

	for _, s := range []int{1, 2, 3, 15} {
		go ListenToQuit(s)
	}

	cmd := flag.CommandLine.Args()
	if len(cmd) == 0 {
		cmd = []string{"echo", "hello kangaroo"}
	}

	if *flagInterval == 0 {
		ListenFor(*flagSigNum, cmd)
		return
	}

	go ListenFor(*flagSigNum, cmd)
	fmt.Printf("executing %s in interval of %d seconds...\n", cmd, *flagInterval)
	for {
		time.Sleep(time.Duration(*flagInterval) * 1000000000)
		fmt.Printf("interval of %d seconds passed, executing %s:\n", *flagInterval, cmd)
		friendly.Run(cmd, "")
	}
}
