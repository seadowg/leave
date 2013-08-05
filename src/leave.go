package main

import (
    "fmt"
    "os"
)

func main() {
    command, time := parseArgs(os.Args);

    if command == "in" {
        fmt.Println("Leaving in", time, "minutes...")
    } else {
        fmt.Println("Leaving at", time, "...")
    }
}

func parseArgs(args []string) (string, string) {
    return args[1], args[2]
}
