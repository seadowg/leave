package main

import (
	"fmt"
	"github.com/deckarep/gosx-notifier"
	"os"
	"strconv"
	"time"
)

func main() {
	command, minutesArg := parseArgs(os.Args)

	if command == "in" {
		minutes, err := strconv.Atoi(minutesArg)

		if err == nil {
			time.Sleep(time.Duration(minutes) * time.Minute)

			notification := gosxnotifier.NewNotification("Time to go!")
			notification.Title = "Leave"
			notification.Sound = gosxnotifier.Basso
			notification.Link = "http://www.youtube.com/watch?v=_W_szJ6M-kM"

			notification.Push()
		} else {
			fmt.Println("Time should be in minutes!")
		}
	}
}

func parseArgs(args []string) (string, string) {
	return args[1], args[2]
}
