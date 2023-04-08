package main

import (
	"flag"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Reminder represents a reminder with a task description, due date/time, and tags.
type Reminder struct {
	Description string
	DueTime     time.Time
	Tags        []string
}

func main(){
	description := flag.String("description", "", "Task description")
	dueTime := flag.String("due", "", "Due date/time (YYYY-MM-DD HH:MM)")
	tags := flag.String("tags", "", "Tags separated by commas")

	flag.Parse()

	if *description == "" {
		fmt.Println("Error: Task description is required")
		flag.PrintDefaults()
		return
	}

	if *dueTime == "" {
		fmt.Println("Error: Due date/time is required.")
		flag.PrintDefaults()
		return
	}

	due, err := time.Parse("2006-01-02 15:04", *dueTime)
	if err != nil {
		fmt.Println("Error: Invalid due date/time format. Please use YYYY-MM-DD HH:MM.")
		return
	}

	// Parse tags
	tagsArr := []string{}
	if *tags != "" {
		tagsArr = append(tagsArr, *tags)
	}

	reminder := Reminder{
		Description: *description,
		DueTime: due,
		Tags: tagsArr,
	}

	fmt.Println("Reminder created successfully:")
	fmt.Println("Description:", reminder.Description)
	fmt.Println("DueTime:", reminder.DueTime)
	fmt.Println("Tags:", reminder.Tags)
}