package reminder

import (
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

func List() []Reminder {
	reminders := []Reminder{}
	return reminders
}