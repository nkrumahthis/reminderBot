package commands

import (
	"flag"
	"fmt"

	"github.com/nkrumahthis/reminderBot/repository"
)

func Edit(id, description, dueTime, tags string)(*repository.Reminder, error) {
	
}

func ParseEditCmd() {
	editCmd := flag.NewFlagSet("edit", flag.ExitOnError)

	description := editCmd.String("description", "", "Task description")
	dueTime := editCmd.String("due", "", "Due date/time (YYYY-MM-DD HH:MM)")
	tags := editCmd.String("tags", "", "Tags separated by commas")

	flag.Parse()

	if flag.NArg() > 0 && flag.Arg(0) == "edit" {
		editCmd.Parse(flag.Args()[1:])
		reminder, err := Edit(id, *description, *dueTime, *tags)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
