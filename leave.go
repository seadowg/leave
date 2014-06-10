package main

import (
	"github.com/deckarep/gosx-notifier"
	"os"
	"strconv"
	"time"
)

func main() {
	command, minutesArg := parseArgs(os.Args)

	if command == "in" {
		if minutes, err := strconv.Atoi(minutesArg); err == nil {
			time.Sleep(time.Duration(minutes) * time.Minute)

			notification := gosxnotifier.NewNotification("Time to go!")
			notification.Push()
		}
	}
}

func parseArgs(args []string) (string, string) {
	return args[1], args[2]
}
