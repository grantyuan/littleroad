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
	ServerAddr string
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

	//	serverAddrs, err := net.LookupHost("www.waterfrontech.com")
	//	if err != nil {
	//		fmt.Println("unable to resovle dns for server address", "www.waterfrontech.com")
	//		os.Exit(6)
	//	}
	var serverAdr string
	if clientConfig.ServerAddr != nil {
		serverAdr := clientConfig.ServerAddr
	} else {
		fmt.Println("unable get server address")
		os.Exit(6)
	}

	saddr, err := net.ResolveTCPAddr("tcp", serverAdr)
	if err != nil {
		fmt.Println("failed to resolve tcp address", serverAdr)
	}

	connid := 0
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Failed to accept connection '%s'", err)
			continue
		}
		connid++
		fmt.Println("connection:", connid, " connect to", raddr.String())
		var p = tcpproxy.New(conn, laddr, raddr)

		p.Nagles = true

		go p.Start()
	}

	//tcpproxy.n
	//tcpproxy.NewTLSUnwrapped()
}

func CreateNewProxy(conn *net.TCPConn, raddr *net.TCPAddr) error {
	//1. connect to server
	//2. tell server my name
	//3. server make a table, you name
	//4. client connect to server:8080,
	//5. client how to connect to the real server?.

	//TODO
}
