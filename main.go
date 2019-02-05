package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Minute)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	input := make(chan string, 5)
	in := bufio.NewScanner(os.Stdin)
	go func() {
		for in.Scan() {
			input <- in.Text()
		}
	}()

	for {
		select {
		case t := <-ticker.C:
			log.Printf("Output time is %s", t.Format("2006-01-02 15:04:05"))
			fmt.Println("this is stdout")
			log.Println("this is stderr")
		case s := <-sigs:
			log.Printf("Interrupted by %s", s)
			return
		case str := <-input:
			fmt.Printf("Got input: %q\n", str)
		}
	}
}
