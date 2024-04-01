package cli

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nkrumahthis/reminderBot/repository"
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	reminderStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)
const listHeight = 14

type model struct {
	list     list.Model
	choice   *repository.Reminder
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(repository.Reminder)
			if ok {
				m.choice = &i
			}
			return m, tea.Quit

		case "a", "A":
			
		}

	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != nil {
		return quitTextStyle.Render(fmt.Sprintf("%s? Sounds good to me.", m.choice))
	}
	if m.quitting {
		return quitTextStyle.Render("Goodbye")
	}
	return "\n" + m.list.View()
}

type reminderDelegate struct{}

func (d reminderDelegate) Height() int { return 1 }

func (d reminderDelegate) Spacing() int { return 0 }

func (d reminderDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d reminderDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(repository.Reminder)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.String())

	fn := reminderStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func Run() model {
	items := []list.Item{}

	reminders := repository.ListReminders()
	for _, reminder := range *reminders {
		items = append(items, reminder)
	}

	const defaultWidth = 20

	l := list.New(items, reminderDelegate{}, defaultWidth, listHeight)
	l.Title = "Which reminder do you want to do?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{list: l}

	return m
}