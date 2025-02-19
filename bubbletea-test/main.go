package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Bubble Tea programs are comprised of a model that describes the application
// state and three simple methods on that model:
//
// * Init, a function that returns an initial command for the application to run.
// * Update, a function that handles incoming events and updates the model accordingly.
// * View, a function that renders the UI based on the data in the model.

// model will store our application state
type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing
	selected map[int]struct{} // which to-do items are selecteds
}

// NewModel returns our applications initial state
func NewModel() model {
	return model{
		// Our to-do list is a grocery list
		choices: []string{"Carrots", "Celery", "Cinnamon"},

		// We dont need to initialize the cursor. Default value of int is fine

		// A map which indicates which choices are selected. The keys refer of
		// the `choices` slice above
		selected: make(map[int]struct{}),
	}
}

// Init returns an initial command for our application
func (m model) Init() tea.Cmd {
	// That is no I/O as of now
	return nil
}

// Update updates the model due to an incomming I/O operation. msg triggers
// the update function
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// ? Is it a key press ?
	case tea.KeyMsg:
		// ? What is the actual key pressed ?
		switch msg.String() {
		// * These keys should exit the program
		case "ctrl+c", "q":
			return m, tea.Quit
		// * These keys should move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		// * These keys should move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		// * The "enter" key and the sapcebar toggle the selected state for the item
		// * that the cursor is pointing
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	return m, nil
}

// View renders the program's UI after every Update
func (m model) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for ii, choice := range m.choices {
		// ? Is the cursor pointing at this choice ?
		cursor := " "
		if m.cursor == ii {
			cursor = ">"
		}

		// ? Is this choice selected ?
		checked := " "
		if _, ok := m.selected[ii]; ok {
			checked = "x"
		}

		// ? Render the row.
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// ? The footer
	s += "\nPress q to quit. \n"

	// ? Send the UI for rendering
	return s
}

func main() {
	p := tea.NewProgram(NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
