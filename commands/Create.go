package commands

import (
	"flag"
	"fmt"

	"github.com/nkrumahthis/reminderBot/repository"
)

func Create (description, dueTime, tags string) (*repository.Reminder, error) {
	if description == "" {
		return nil, fmt.Errorf("error: Task description is required")
	}

	if dueTime == "" {
		return nil, fmt.Errorf("error: Due date/time is required")
	}

	reminder := &repository.Reminder{
		Description: description,
		DueTime: dueTime,
		Tags: tags,
	}

	_, err := repository.CreateReminder(reminder)
	if err != nil {
		panic(err)
	}

	return reminder, nil
}

func ParseCreateCmd() {
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)

	// define create subcommand flagsß
	description := createCmd.String("description", "", "Task description")
	dueTime := createCmd.String("due", "", "Due date/time (YYYY-MM-DD HH:MM)")
	tags := createCmd.String("tags", "", "Tags separated by commas")

	flag.Parse()

	if flag.NArg() > 0 && flag.Arg(0) == "create" {
		createCmd.Parse(flag.Args()[1:])
		reminder, err := Create(*description, *dueTime, *tags)
		if err != nil {
			fmt.Println(err)
			return
		}

		// print print reminder
		fmt.Println("Reminder created successfully:")
		fmt.Println("Id:", reminder.Id)
		fmt.Println("Description:", reminder.Description)
		fmt.Println("Due date/time:", reminder.DueTime)
		fmt.Println("Tags:", reminder.Tags)


	}
}
