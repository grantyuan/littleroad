// LittleRoad
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	tcpproxy "github.com/grantyuan/go-tcp-proxy"
)

type LittleRoadClientConfig struct {
	LocalAddr  string
	RemoteAddr string
}

func main() {
	//read config
	bytej, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	clientConfig := new(LittleRoadClientConfig)

	err = json.Unmarshal(bytej, clientConfig)
	if err != nil {
		fmt.Println("unable to read config file")
		os.Exit(2)
	}

	laddr, err := net.ResolveTCPAddr("tcp", clientConfig.LocalAddr)
	if err != nil {
		fmt.Println("unable to resove tcp addr" + clientConfig.LocalAddr)
		os.Exit(3)
	}

	raddr, err := net.ResolveTCPAddr("tcp", clientConfig.RemoteAddr)
	if err != nil {
		fmt.Println("unable to resove tcp addr" + clientConfig.RemoteAddr)
		os.Exit(4)
	}

	listener, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		fmt.Println("Failed to open local port to listen: %s", err)
		os.Exit(5)
	}
	connid := 0
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Failed to accept connection '%s'", err)
			continue
		}
		connid++
		fmt.Println("connection:", connid)
		var p = tcpproxy.New(conn, laddr, raddr)

		go p.Start()
	}

	//tcpproxy.n
	//tcpproxy.NewTLSUnwrapped()
}
