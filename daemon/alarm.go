package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/0xAX/notificator"
)

type Reminder struct {
	Description string
	DueTime     string
}

var reminders = []Reminder{
	{Description: "Reminder 1", DueTime: "08:00"},
	{Description: "Reminder 2", DueTime: "12:30"},
	{Description: "Reminder 3", DueTime: "22:47"},

	// Add more reminders here
}

func main() {
	fmt.Println("Reminder Daemon Started")

	// Create a channel to receive OS signals
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Initialize the notifier
	notify := notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "Reminder Daemon",
	})

	// Run the daemon loop
	for {
		now := time.Now().Format("15:04") // Format the current time as "hh:mm"

		// Check if any reminders are due
		for _, reminder := range reminders {
			if reminder.DueTime == now {
				fmt.Println("matches")
				message := fmt.Sprintf("Reminder: %s", reminder.Description)
				notify.Push("Reminder", message, "", notificator.UR_NORMAL)
			}
		}

		// Sleep for 1 minute before checking again
		time.Sleep(time.Second * 5)

		// Check for termination signal
		select {
		case <-sig:
			fmt.Println("Termination signal received. Exiting...")
			return
		default:
			// Continue looping
		}
	}
}
