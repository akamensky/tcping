package main

import (
	"github.com/akamensky/argparse"
	"os"
	"fmt"
	"time"
	"net"
)

func main() {
	p := argparse.NewParser("", "A utility to make sure the remote TCP port is reachable and listening")

	server := p.String("s", "server", &argparse.Options{Required: true, Help: "Remote server to check"})
	port := p.Int("p", "port", &argparse.Options{Required: true, Help: "Remote port to check"})
	timeoutArg := p.Int("t", "timeout", &argparse.Options{Default: 500, Help: "Connection timeout in ms"})

	err := p.Parse(os.Args)
	if err != nil {
		fmt.Print(p.Usage(err))
		os.Exit(1)
	}

	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan interface{})

	var seqNumber uint64 = 0
	var network = fmt.Sprintf("%s:%d", *server, *port)
	var timeout = time.Duration(*timeoutArg) * time.Millisecond

	for ;;seqNumber++ {
		select {
		case <-ticker.C:
			tryPort(network, seqNumber, timeout)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func tryPort(network string, seq uint64, timeout time.Duration) {
	startTime := time.Now()
	conn, err := net.DialTimeout("tcp", network, timeout)
	endTime := time.Now()
	if err != nil {
		os.Stdout.Write([]byte(startTime.Format("[2006-01-02T15:04:05]:") + " connection failed\n"))
	} else {
		var t = float64(endTime.Sub(startTime)/time.Millisecond)
		os.Stdout.Write([]byte(startTime.Format("[2006-01-02T15:04:05]:") + fmt.Sprintf(" addr=%s seq=%d time=%4.2fms\n", conn.RemoteAddr().String(), seq, t)))
	}
}