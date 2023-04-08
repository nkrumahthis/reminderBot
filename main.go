package main

import (
	"flag"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)


func main(){
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	description := createCmd.String("description", "", "Task description")
	dueTime := createCmd.String("due", "", "Due date/time (YYYY-MM-DD HH:MM)")
	tags := createCmd.String("tags", "", "Tags separated by commas")

	flag.Parse()

	// Check if subcommand "create" is provided
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
	} else if flag.NArg() > 0 && flag.Arg(0) == "list" {
		listCmd.Parse(flag.Args()[1:])
		reminders := List()

		fmt.Println(reminders)
	} else {
		fmt.Println(flag.Arg(0), "is not a reminderBot command.")
		flag.PrintDefaults()
		return
	}
}