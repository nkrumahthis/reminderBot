package main

import (
	"flag"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	commands "github.com/nkrumahthis/reminderBot/commands"
	"github.com/nkrumahthis/reminderBot/db"
)


func main(){
	db.Init()
	
	// Define "create" and "list" subcommands
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	// Parse command-line arguments
	flag.Parse()

	// Check which subcommand is provided
	if flag.NArg() > 0 && flag.Arg(0) == "create" {
		commands.ParseCreateCmd()
	} else if flag.NArg() > 0 && flag.Arg(0) == "list" {
		listCmd.Parse(flag.Args()[1:])
		reminders := List()

		fmt.Println(reminders)
	} else if flag.NArg() > 0 && flag.Arg(0) == "delete" {
		fmt.Println("reminderBot", flag.Args()[0], "is not implemented yet")
		return
	} else if flag.NArg() > 0 && flag.Arg(0) == "edit" {
		fmt.Println("reminderBot", flag.Args()[0], "is not implemented yet")
		return
	} else if flag.NArg() > 0 && flag.Arg(0) == "snooze" {
		fmt.Println("reminderBot", flag.Args()[0], "is not implemented yet")
		return
	} else if flag.NArg() > 0 && flag.Arg(0) == "search" {
		fmt.Println("reminderBot", flag.Args()[0], "is not implemented yet")
		return
	} else if flag.NArg() > 0 && flag.Arg(0) == "config" {
		fmt.Println("reminderBot", flag.Args()[0], "is not implemented yet")
		return
	} else if flag.NArg() > 0 && flag.Arg(0) == "help" {
		fmt.Println("reminderBot", flag.Args()[0], "is not implemented yet")
		return
	} else {
		fmt.Println(flag.Arg(0), "is not a reminderBot command.")
		flag.PrintDefaults()
		return
	}
}
