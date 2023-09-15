package commands

import (
	"flag"
	"fmt"

	"github.com/nkrumahthis/reminderBot/repository"
)

func ParseListCmd() {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	listCmd.Parse(flag.Args()[1:])
	reminders := repository.ListReminders()

	for _, reminder := range reminders {
		fmt.Println(reminder.String())
	}
}
