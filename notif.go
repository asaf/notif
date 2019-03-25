package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Verify that a sub-command has been provided
	// os.Arg[0] is the main command
	// os.Arg[1] will be the sub-command
	if len(os.Args) < 2 {
		fmt.Println("[smtp] sub-command is required")
		os.Exit(1)
	}

	// Switch on the sub-command
	// os.Args[2:] will be all arguments starting after the sub-command at os.Args[1]
	switch os.Args[1] {
	case "smtp":
		if err := Smtp(os.Args[2:]); err != nil {
			panic(err)
		}
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
