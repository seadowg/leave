package main

import (
	"fmt"
	"github.com/deckarep/gosx-notifier"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		help()
	} else {
		command, minutesArg := os.Args[1], os.Args[2]

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
				help()
			}
		} else {
			help()
		}
	}
}

func help() {
	fmt.Println("Format: 'leave in x' where 'x' is minutes")
}
