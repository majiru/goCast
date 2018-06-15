package main

import (
	"flag"
	"fmt"
	"github.com/majiru/gocast/api"
	"os"
)

func main() {
	var port = flag.String("port", "8080", "Port for webserver")

	if len(os.Args) < 2 {
		printUsage()
	}

	switch os.Args[1] {
	case "client":
		gocast.Listen()
	case "server":
		if len(os.Args) < 3 {
			printUsage()
		} else {
			gocast.Serve(os.Args[2], os.Args[3], *port)
		}
	}

}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println(os.Args[0] + " client")
	fmt.Println(os.Args[0] + " -port=8080 server /path/to/media/dir playerIP")
	os.Exit(1)
}
