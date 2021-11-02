package models

import (
	"io"
	"log"
	"net"
)

type PortForward struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Network     string `json:"network"`
}

func (pf *PortForward) Listen() {
	ln, err := net.Listen(pf.Network, pf.Source)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		sourceConnection, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go handleRequest(pf.Network, sourceConnection, pf.Destination)
	}
}

func handleRequest(network string, sourceConnection net.Conn, destinationAddress string) {
	destinationConnection, err := net.Dial(network, destinationAddress)
	if err != nil {
		panic(err)
	}

	go copyIO(sourceConnection, destinationConnection)
	go copyIO(destinationConnection, sourceConnection)
}

func copyIO(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	io.Copy(src, dest)
}
