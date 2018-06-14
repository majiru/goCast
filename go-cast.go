package main

import (
	"fmt"
	"github.com/majiru/gocast/api"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
	}

	switch os.Args[1] {
	case "client":
		gocast.Listen()
	case "serve":
		if len(os.Args) < 3 {
			gocast.Serve(os.Args[1], os.Args[2])
		} else {
			printUsage()
		}
	}

}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println(os.Args[0] + " client")
	fmt.Println(os.Args[0] + " server /path/to/media/dir playerIP")
	os.Exit(1)
}
