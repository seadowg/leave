package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	command, time := parseArgs(os.Args)

	if command == "in" {
		if minutes, err := strconv.Atoi(time); err == nil {
			fmt.Println("Leaving in", minutes, "minutes...")
		}
	}
}

func parseArgs(args []string) (string, string) {
	return args[1], args[2]
}
