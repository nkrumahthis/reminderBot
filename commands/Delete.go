package commands

import (
	"flag"
	"fmt"

	"github.com/nkrumahthis/reminderBot/repository"
)

func Delete(id int64) error {
	return repository.DeleteReminder(id)
}

func ParseDeleteCmd() {
	// Define a command-line flag for an integer value
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	// define delete subcommand flags
	id := deleteCmd.Int64("id", -1, "Task id")

	flag.Parse()
	if flag.Arg(0) == "delete" {
		deleteCmd.Parse(flag.Args()[1:])
		err := Delete(*id)
		if err != nil {
			panic(err)
		}
		fmt.Println("Reminder deleted successfully:")

	}

}
