package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
)

const switchSign = ";;"
const endSign = "\r\n\r\n"

type file struct {
	Filename string
	Link     string
}

type page struct {
	Files []file
}

var (
	fileDir  string
	playerIP = "localhost"
	address  string
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: " + os.Args[0] + " dir playerIP")
		os.Exit(1)
	}

	fileDir = os.Args[1]
	if fileDir[len(fileDir)-1] == '/' {
		fileDir = fileDir[:len(fileDir)-2]
	}
	playerIP = os.Args[2]
	address, _ = os.Hostname()

	http.HandleFunc("/", rootHandler)
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir(fileDir)+"/")))
	http.HandleFunc("/command/", commandHandler)
	http.HandleFunc("/dir/", dirHandler)
	log.Fatal(http.ListenAndServe(":80", nil))

}

func writeCommand(commands []string) {
	conn, err := net.Dial("tcp", playerIP+":8000")

	defer conn.Close()

	if err != nil {
		return
	}

	var buffer string
	for _, command := range commands {
		buffer += command + switchSign
	}
	buffer += endSign
	conn.Write([]byte(buffer))
}

func walkDir(path string, files *[]file) {
	current, err := ioutil.ReadDir(fileDir + path)

	if err != nil {
		return
	}

	for _, f := range current {
		v := url.Values{}
		var link string
		v.Add("File", url.QueryEscape(path+f.Name()))

		if f.IsDir() {
			link = "/dir/?" + v.Encode() + "/"
		} else {
			v.Add("Command", "Open")
			link = "/command/?" + v.Encode()
		}
		*files = append(*files, file{f.Name(), link})
	}

}

func commandHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var commands []string
	switch q.Get("Command") {
	case "Open":
		commands = append(commands, "open")
		path, _ := url.QueryUnescape(q.Get("File"))
		commands = append(commands, "http://"+address+"/files"+path)
	case "Play":
		commands = append(commands, "play")
	case "Pause":
		commands = append(commands, "pause")
	case "Next":
		commands = append(commands, "next")
	}
	writeCommand(commands)
	w.Write([]byte("Sent Command!!"))
}

func dirHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var files []file
	path, _ := url.QueryUnescape(q.Get("File"))
	walkDir(path, &files)
	t, _ := template.ParseFiles("main.tmpl")
	t.Execute(w, page{files})
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	var files []file
	walkDir("/", &files)

	t, _ := template.ParseFiles("main.tmpl")

	t.Execute(w, page{files})
}
