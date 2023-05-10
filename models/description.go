package models

import (
	"DAB-SSH/helpers"
	"DAB-SSH/styling"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

type DescriptionPage struct {
	//dabTitle    string           // DAB written at top of page
	projectName string           // project name
	viewport    viewport.Model   // Viewport for scrolling - sets content upon creation
	help        help.Model       // Help bar at bottom of page
	keys        helpers.PPkeyMap // Key map for our help model
	minWidth    int              // Minimum Width so model won't break
}

// The about page DAB body
// var DABTitle, _ = os.ReadFile("content/DAB.md")
var BotBTitle, _ = os.ReadFile("content/descriptionpage/BotBTitle.md")
var BotBContent, _ = os.ReadFile("content/descriptionpage/BotB.md")

/*
  ______ .______       _______     ___   .___________. _______     .___  ___.   ______    _______   _______  __
 /      ||   _  \     |   ____|   /   \  |           ||   ____|    |   \/   |  /  __  \  |       \ |   ____||  |
|  ,----'|  |_)  |    |  |__     /  ^  \ `---|  |----`|  |__       |  \  /  | |  |  |  | |  .--.  ||  |__   |  |
|  |     |      /     |   __|   /  /_\  \    |  |     |   __|      |  |\/|  | |  |  |  | |  |  |  ||   __|  |  |
|  `----.|  |\  \----.|  |____ /  _____  \   |  |     |  |____     |  |  |  | |  `--'  | |  '--'  ||  |____ |  `----.
 \______|| _| `._____||_______/__/     \__\  |__|     |_______|    |__|  |__|  \______/  |_______/ |_______||_______|

*/

// Creates and gives our model values
func CreateDescriptionPage(projectAddress int, projectName string) DescriptionPage {

	// Renders content seperately from titles
	renderedContent, _ := glamour.Render(string(BotBContent), "dracula")

	// Create Viewport and sets content
	viewport := viewport.New(TerminalWidth-styling.Border.GetPaddingLeft(), TerminalHeight-15)
	viewport.SetContent(string(BotBTitle) + renderedContent)

	// Sets the help model and styling
	help := help.New()
	help.Styles.ShortKey = styling.APHelpBar
	help.Styles.FullKey = styling.APHelpBar

	// Return our created model
	return DescriptionPage{
		//dabTitle:    string(DABTitle),
		projectName: projectName,
		viewport:    viewport,
		help:        help,
		keys:        helpers.APkeys, // Change later
		minWidth:    85,             // Change later 65
	}
}

// Initializes our struct as a bubble tea model
func (d DescriptionPage) Init() tea.Cmd {
	return nil
}

/*
 __    __  .______    _______       ___   .___________. _______
|  |  |  | |   _  \  |       \     /   \  |           ||   ____|
|  |  |  | |  |_)  | |  .--.  |   /  ^  \ `---|  |----`|  |__
|  |  |  | |   ___/  |  |  |  |  /  /_\  \    |  |     |   __|
|  `--'  | |  |      |  '--'  | /  _____  \   |  |     |  |____
 \______/  | _|      |_______/ /__/     \__\  |__|     |_______|
*/
// Updates our model everytime a key event happens, mainly window resizes and key presses
func (d DescriptionPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Sets cmd as a tea command that can be easily changed later
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	// Sets msg as a switch for all events
	switch msg := msg.(type) {

	// Runs whenever the window is resized or first loaded
	case tea.WindowSizeMsg:
		// Sets the help model and main model width for sizing later
		d.help.Width = msg.Width - styling.HelpBar.GetPaddingLeft()

		// Sets terminal width and height
		TerminalWidth = msg.Width
		TerminalHeight = msg.Height

		// Viewport Size
		d.viewport.Width = msg.Width - styling.Border.GetPaddingLeft()
		d.viewport.Height = msg.Height - 15
	// All key presses
	case tea.KeyMsg:

		// Converts the press into a string
		switch msg.String() {

		// Back to project page
		case "esc":
			return CreateProjectPage(), tea.ClearScreen

		// Quit the program
		case "q":
			return d, tea.Quit

			// Switches between full help view
		case "?":
			d.help.ShowAll = !d.help.ShowAll

		}
	}
	// Handle keyboard and mouse events in the viewport
	d.viewport, cmd = d.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return d, tea.Batch(cmds...)
}

/*
____    ____  __   ___________    __    ____
\   \  /   / |  | |   ____\   \  /  \  /   /
 \   \/   /  |  | |  |__   \   \/    \/   /
  \      /   |  | |   __|   \            /
   \    /    |  | |  |____   \    /\    /
    \__/     |__| |_______|   \__/  \__/
*/

// Renders our model formatted to be viewed, then returns as a string
func (d DescriptionPage) View() string {

	// Our s string to build our model
	var s string

	// Size to return our model later
	var width, height int

	// Logic for setting terminal size to not break model
	if TerminalWidth <= d.minWidth {
		width = d.minWidth
	} else {
		width = TerminalWidth
	}
	height = TerminalHeight

	// Adds the help bar at the bottom
	fullHelpView := d.help.View(d.keys)

	// RENDERING OUR MODEL |*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|
	// |*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|

	// Adds DAB title
	//width = 0
	//s += d.dabTitle + "\n" // -- if Width != 0 then breaks shit

	s += styling.LightBlue.Render(strings.Repeat("━", TerminalWidth-styling.Border.GetPaddingLeft()))
	s += "\n\n"

	// Adds viewport and lower blue border
	s += styling.DPViewport.Render(d.viewport.View()) + "\n\n"
	s += styling.LightBlue.Render(strings.Repeat("━", TerminalWidth-styling.Border.GetPaddingLeft()))

	// Counts empty lines to put help model at bottom of terminal
	emptyLines := TerminalHeight - strings.Count(s, "\n") - 3
	if emptyLines < 0 {
		emptyLines = 0
	}

	// Add empty lines if there are any to bottom of terminal
	s += strings.Repeat("\n", emptyLines)

	// Render help bar in correct styling
	s += styling.HelpBar.Render(fullHelpView)

	return styling.Border.Width(width).Height(height).Render(s)
}
