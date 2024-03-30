package main

import (
	"flag"
	"fmt"

	"github.com/nkrumahthis/reminderBot/commands"
	"github.com/nkrumahthis/reminderBot/db"
)


func main(){
	db.Init()

	// Parse command-line arguments
	flag.Parse()

	// Check which subcommand is provided
	if flag.NArg() > 0 && flag.Arg(0) == "create" {
		commands.ParseCreateCmd()
		return
	} else if flag.NArg() > 0 && flag.Arg(0) == "list" {
		commands.ParseListCmd()
		return
	} else if flag.NArg() > 0 && flag.Arg(0) == "delete" {
		commands.ParseDeleteCmd()
		return
	} else if flag.NArg() > 0 && flag.Arg(0) == "edit" {
		fmt.Println("reminderBot Edit is not implemented yet")
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
