package repository

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nkrumahthis/reminderBot/db"
)

// Reminder represents a reminder with a task description, due date/time, and tags.
type Reminder struct {
	Id          int64
	Description string
	DueTime     string
	Tags        string
}

func (r *Reminder) String() string {
	due, err := time.Parse(time.RFC3339, r.DueTime)
	if err != nil {
		panic(err)
	}

	// Convert int64 to string
	id := strconv.FormatInt(r.Id, 10)

	return "[" + id + "] " + due.Format(time.Kitchen) + " |-> " + r.Description + r.Tags
}

func validate(description string, dueTime string, tags string) error {
	if description == "" {
		return fmt.Errorf("error: --description is required")
	}

	_, err := time.Parse("2006-01-02 15:04", dueTime)
	if err != nil {
		fmt.Println(dueTime)
		return fmt.Errorf("error: Invalid due date/time format. Please use YYYY-MM-DD HH:MM")
	}

	return nil
}

func CreateReminder(reminder *Reminder) (*Reminder, error) {

	err := validate(reminder.Description, reminder.DueTime, reminder.Tags)
	if err != nil {
		panic(err)
	}

	result, err := db.DB.Exec(
		"INSERT INTO reminders (description, due, tags) VALUES ($1, $2, $3)",
		reminder.Description,
		reminder.DueTime,
		reminder.Tags,
	)

	lastInsertID, err := result.LastInsertId()
	reminder.Id = lastInsertID

	return reminder, err
}

func ListReminders() *[]Reminder {
	reminders := []Reminder{}

	rows, err := db.DB.Query("SELECT * FROM reminders;")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int64
		var description string
		var due string
		var tags string

		err = rows.Scan(&id, &description, &due, &tags)
		reminders = append(reminders, Reminder{id, description, due, tags})
		if err != nil {
			panic(err)
		}
	}
	return &reminders
}

func GetReminder(id int64) (*Reminder, error) {
	row, err := db.DB.Query("SELECT * FROM reminder WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	var reminder Reminder
	err = row.Scan(&reminder.Id, &reminder.Description, &reminder.DueTime, &reminder.Tags)
	if err != nil {
		return nil, err
	}

	return &reminder, nil
}

func EditReminder(reminder *Reminder) (*Reminder, error) {
	err := validate(reminder.Description, reminder.DueTime, reminder.Tags)
	if err != nil {
		return nil, err
	}

	_, err = db.DB.Exec(
		"UPDATE reminders SET description = $1, due = $2, tags = $3 WHERE id = $4",
		reminder.Description,
		reminder.DueTime,
		reminder.Tags,
		reminder.Id,
	)

	if err != nil {
		return nil, err
	}

	return reminder, nil
}

func DeleteReminder(id int64) error {
	_, err := db.DB.Exec(
		"DELETE FROM reminders WHERE id = $1",
		id,
	)

	if err != nil {
		panic(err)
	}

	return err
}
