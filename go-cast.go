package main

import (
	"flag"
	"fmt"
	"github.com/majiru/gocast/api"
	"os"
)

func main() {
	var port = flag.String("port", "8080", "Port for webserver")
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		printUsage()
	}

	switch args[0] {
	case "client":
		gocast.Listen()
	case "server":
		if len(args) < 3 {
			printUsage()
		} else {
			gocast.Serve(args[1], args[2], *port)
		}
	}

}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println(os.Args[0] + " client")
	fmt.Println(os.Args[0] + " -port=8080 server /path/to/media/dir playerIP")
	os.Exit(1)
}
