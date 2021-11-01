package main

import (
	"io"
	"log"
	"net"
)

func (r *Redirection) Listen() {
	ln, err := net.Listen(r.Network, r.Source)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		// ToDo: add killswitch here
		sourceConnection, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go handleRequest(r.Network, sourceConnection, r.Destination)
	}
}

type Redirection struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Network     string `json:"network"`
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
