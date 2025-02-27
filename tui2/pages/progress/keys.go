package progress

import (
	"github.com/charmbracelet/bubbles/key"

	"github.com/streamingfast/substreams/tui2/keymap"
)

func (p *Progress) ShortHelp() []key.Binding {
	return []key.Binding{
		keymap.MainNavigation,
		keymap.UpDown,
		keymap.ToggleProgressDisplayMode,
		keymap.RestartStream,
		keymap.Build,
		keymap.Help,
		keymap.Quit,
	}
}

func (p *Progress) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			keymap.MainNavigation,
		},
		{
			keymap.UpDown,
			keymap.UpDownPage,
		},
		{
			keymap.ToggleProgressDisplayMode,
		},
		{
			keymap.RestartStream,
		},
		{
			keymap.Build,
		},
		{
			keymap.Help,
			keymap.Quit,
		},
	}
}
