package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	cli "github.com/nkrumahthis/reminderBot/cli/viewmodels"
	"github.com/nkrumahthis/reminderBot/db"
)

func main() {
	db.Init()

	homeModel := cli.Run()

	if _, err := tea.NewProgram(homeModel).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
