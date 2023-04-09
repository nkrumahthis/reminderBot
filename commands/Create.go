package commands

import (
	"flag"
	"fmt"
	"time"
)

// Reminder represents a reminder with a task description, due date/time, and tags.
type Reminder struct {
	Description string
	DueTime     time.Time
	Tags        []string
}

func Create (description, dueTime, tags string) (*Reminder, error) {
	if description == "" {
		return nil, fmt.Errorf("Error: Task description is required")
	}

	if dueTime == "" {
		return nil, fmt.Errorf("Error: Due date/time is required.")
	}

	due, err := time.Parse("2006-01-02 15:04", dueTime)
	if err != nil {
		return nil, fmt.Errorf("Error: Invalid due date/time format. Please use YYYY-MM-DD HH:MM.")
	}

	// Parse tags
	tagsArr := []string{}
	if tags != "" {
		tagsArr = append(tagsArr, tags)
	}

	reminder := &Reminder{
		Description: description,
		DueTime: due,
		Tags: tagsArr,
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

	fmt.Println("flag.Narg", flag.NArg())

	if flag.NArg() > 0 && flag.Arg(0) == "create" {
		createCmd.Parse(flag.Args()[1:])
		reminder, err := Create(*description, *dueTime, *tags)
		if err != nil {
			fmt.Println(err)
			return
		}

		// print print reminder
		fmt.Println("Reminder created successfully:")
		fmt.Println("Description:", reminder.Description)
		fmt.Println("Due date/time:", reminder.DueTime)
		fmt.Println("Tags:", reminder.Tags)
	}
}
