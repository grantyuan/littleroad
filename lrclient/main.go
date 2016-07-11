// lrclient project main.go
package main

import (
	"fmt"
	"net"
	"os"

	tcpproxy "github.com/grantyuan/go-tcp-proxy"
)

func main() {
	addrs, err := net.LookupHost("www.waterfrontech.com:8081")
	if err != nil {
		fmt.Println("unable resolve server")
		os.Exit(1)

	}

	laddr, err := net.ResolveTCPAddr("tcp", ":8081")
	raddr, err := net.ResolveTCPAddr("tcp", addrs[0])
	if err != nil {
		fmt.Println("unalbe resove TCP addr")
		os.Exit(1)
	}

	for {
		conn, err := net.DialTCP("tcp", laddr, raddr)
		if err != nil {
			fmt.Println("unable to connect to server")

		}

		p := tcpproxy.New(conn, laddr, raddr)
		p.Log = tcpproxy.ColorLogger{
			Verbose:     true,
			VeryVerbose: true,
			Prefix:      fmt.Sprintf("Connection #%03d ", connid),
			Color:       true,
		}
		p.Nagles = true
		go p.Start()
	}

	defer conn.Close()

}
