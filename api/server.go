package gocast

import (
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
)

type file struct {
	Filename string
	Link     string
}

type page struct {
	Files []file
}

var (
	fileDir    string
	playerIP   string
	playerPort string
	address    string
)

//Serve creates and initalized the gocast server//
func Serve(directory, playerAddress, port string) {
	fileDir = directory
	if fileDir[len(fileDir)-1] == '/' {
		fileDir = fileDir[:len(fileDir)-2]
	}
	playerIP = playerAddress
	address, _ = os.Hostname()
	playerPort = port

	http.HandleFunc("/", rootHandler)
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir(fileDir)+"/")))
	http.HandleFunc("/command/", commandHandler)
	http.HandleFunc("/dir/", dirHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))

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
		commands = append(commands, "http://"+address+":"+playerPort+"/files"+path)
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
	t, _ := template.New("mainPage").Parse(mainPage)
	t.Execute(w, page{files})
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	var files []file
	walkDir("/", &files)

	t, _ := template.New("mainPage").Parse(mainPage)
	t.Execute(w, page{files})
}
