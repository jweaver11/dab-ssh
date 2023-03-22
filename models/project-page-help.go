package models

import (
	"github.com/charmbracelet/bubbles/key"
)

// Sets a keymap struct to store the controls and key bind variables
// So they can be called on later for the help view
type PPkeyMap struct {
	Advance key.Binding
	Quit    key.Binding
}

// Built in function from the help package that shows our mini help view at the bottom of our active model
// It is part of the key.Map interface
func (k PPkeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Advance, k.Quit}
}

// Built in function from the help package that shows our full help view at the bottom of our active model
// It is part of the key.Map interface
func (k PPkeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Advance, k.Quit},
	}
}

// Sets keys as our object using our keyMap struct from above
var PPkeys = WPkeyMap{
	Advance: key.NewBinding(
		key.WithKeys(""),
		key.WithHelp("Press any key", "to continue"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
}
