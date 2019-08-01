package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var interval = flag.String("interval", "1m", "Interval between output, should be parsable by time.ParseDuration")

func main() {
	flag.Parse()

	d, err := time.ParseDuration(*interval)
	if err != nil {
		log.Fatalf("Could not parse output interval: %v", err)
	}

	ticker := time.NewTicker(d)
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
