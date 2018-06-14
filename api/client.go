package gocast

import (
	"bufio"
	"fmt"
	"github.com/blang/mpv"
	"log"
	"net"
	"strings"
)

var mpvc *mpv.Client

//Listen starts a new client gocast client listener//
func Listen() {
	mpvc = mpv.NewClient(mpv.NewIPCClient("/tmp/mpvsocket"))

	ln, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleConn(conn)

	}

}

func handleConn(conn net.Conn) {
	var r = bufio.NewReader(conn)

	buffer := make([]byte, 512)
	for {
		n, err := r.Read(buffer)
		input := string(buffer[:n])
		fmt.Println(input)
		if err != nil {
			break
		}
		if strings.HasSuffix(input, endSign) {
			parseInput(strings.Split(input, endSign)[0])
		}
	}

}

func parseInput(inputRaw string) {
	input := strings.Split(inputRaw, switchSign)
	switch input[0] {
	case "open":
		fmt.Println("Loading file...")
		//mpvc.Loadfile("http://"+input[1]+"/"+input[2], mpv.LoadFileModeAppend)
		mpvc.Loadfile(input[1], mpv.LoadFileModeAppendPlay)
		mpvc.SetPause(false)
	case "play":
		mpvc.SetPause(false)
		mpvc.SetFullscreen(true)
	case "pause":
		mpvc.SetPause(true)
	case "next":
		mpvc.PlaylistNext()
	}
}
