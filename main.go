package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-ping/ping"
)

func main() {
	var ips []string
	f, err := os.Open("ips.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		ips = append(ips, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Concurrent pings

	c := make(chan string)
	fmt.Print("\nConcurrent pings... \n\n")

	startC := time.Now()
	for _, ip := range ips {
		go concurrentPing(ip, c)
	}
	for range ips {
		fmt.Println(<-c)
	}
	elapsedC := time.Since(startC)

	// Sequential pings

	fmt.Print("\nSequential pings... \n\n")

	startS := time.Now()
	for _, ip := range ips {
		sequentialPing(ip)
	}
	elapsedS := time.Since(startS)

	fmt.Println("\nConcurrent done in => ", elapsedC.Seconds())
	fmt.Println("Sequential done in => ", elapsedS.Seconds())
}

func concurrentPing(ip string, c chan string) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	pinger.Count = 5

	pinger.OnFinish = func(stats *ping.Statistics) {
		msg := "PING " + pinger.Addr() + "(" + pinger.IPAddr().String() + ") => " + "avg = " + stats.AvgRtt.String()
		c <- msg
	}

	pinger.Run()
}

func sequentialPing(ip string) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	pinger.Count = 5

	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Println("PING " + pinger.Addr() + "(" + pinger.IPAddr().String() + ") => " + "avg = " + stats.AvgRtt.String())
	}

	pinger.Run()
}
