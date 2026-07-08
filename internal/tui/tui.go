package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	"github.com/inflame-ue/pastiche/internal/config"
)

type model struct {
	config        *config.Config
	focus         int
	width, height int
	modeIdx       int
	modes         []string
	formatterIdx  int
	buttonIdx     int
}

func InitialModel(config *config.Config) model {
	modeIdx := 0
	for i, m := range []string{"hotkey", "autowatch", "both"} {
		if m == config.Trigger.Mode {
			modeIdx = i
			break
		}
	}
	return model{
		config:       config,
		focus:        0,
		modeIdx:      modeIdx,
		modes:        []string{"hotkey", "autowatch", "both"},
		formatterIdx: 0,
		buttonIdx:    0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			m.focus = (m.focus + 1) % 5
		case "shift+tab":
			m.focus = (m.focus - 1 + 5) % 5
		case "enter":
			if m.focus == 4 {
				switch m.buttonIdx {
				case 0:
					config.Save(m.config)
					return m, tea.Quit
				case 1:
					return m, tea.Quit
				}
			}
		default:
			switch m.focus {
			case 0:
				if msg.String() == "up" || msg.String() == "left" {
					m.modeIdx = (m.modeIdx - 1 + 3) % 3
				} else if msg.String() == "down" || msg.String() == "right" {
					m.modeIdx = (m.modeIdx + 1) % 3
				}
				m.config.Trigger.Mode = m.modes[m.modeIdx]
			case 1:
				code := hotkeyCode(msg.String())
				if code != 0 {
					m.config.Hotkey.Key = code
				}
			case 2:
				if msg.String() == "up" || msg.String() == "right" {
					if m.config.Heuristic.Value < 5 {
						m.config.Heuristic.Value++
					}
				} else if msg.String() == "down" || msg.String() == "left" {
					if m.config.Heuristic.Value > 1 {
						m.config.Heuristic.Value--
					}
				}
			case 3:
				order := m.config.Formatters.Order
				if len(order) == 0 {
					break
				}
				switch msg.String() {
				case "up":
					if m.formatterIdx > 0 {
						order[m.formatterIdx], order[m.formatterIdx-1] = order[m.formatterIdx-1], order[m.formatterIdx]
						m.formatterIdx--
					}
				case "down":
					if m.formatterIdx < len(order)-1 {
						order[m.formatterIdx], order[m.formatterIdx+1] = order[m.formatterIdx+1], order[m.formatterIdx]
						m.formatterIdx++
					}
				}
			case 4:
				if msg.String() == "left" {
					m.buttonIdx = (m.buttonIdx - 1 + 2) % 2
				} else if msg.String() == "right" {
					m.buttonIdx = (m.buttonIdx + 1) % 2
				}
			}
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	const indent = "    "

	bold := lipgloss.NewStyle().Bold(true).Render
	active := lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Render
	accent := lipgloss.NewStyle().Foreground(lipgloss.Color("14")).Render
	box := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0, 2).Width(70)
	focus := func(idx int) string {
		if m.focus == idx {
			return accent("▎")
		}
		return " "
	}

	var b strings.Builder

	fmt.Fprintln(&b, "pastiche configure")
	fmt.Fprintln(&b)

	{
		var inner strings.Builder
		fmt.Fprintln(&inner, focus(0)+bold("Trigger Mode"))
		fmt.Fprint(&inner, indent)
		for i, mode := range []string{"Hotkey", "Autowatch", "Both"} {
			if i > 0 {
				fmt.Fprint(&inner, indent)
			}
			dot := "○"
			if m.modes[i] == m.config.Trigger.Mode {
				dot = active("●")
			}
			fmt.Fprintf(&inner, "%s %s", dot, mode)
		}
		fmt.Fprintln(&b, box.Render(inner.String()))
	}

	fmt.Fprintln(&b)

	{
		var inner strings.Builder
		fmt.Fprintln(&inner, focus(1)+bold("Hotkey Binding"))
		kn := "Press any key to set..."
		if k := keyName(m.config.Hotkey.Key); k != "" {
			kn = k
		}
		fmt.Fprintf(&inner, "%s[ %s ]   Current: %s", indent, kn, kn)
		fmt.Fprintln(&b, box.Render(inner.String()))
	}

	fmt.Fprintln(&b)

	{
		var inner strings.Builder
		fmt.Fprintln(&inner, focus(2)+bold("Heuristic Sensitivity"))
		fmt.Fprint(&inner, indent+"Loose ")
		for i := 1; i <= 5; i++ {
			if i <= m.config.Heuristic.Value {
				fmt.Fprint(&inner, active("●"))
			} else {
				fmt.Fprint(&inner, "○")
			}
		}
		fmt.Fprintln(&inner, " Strict")
		fmt.Fprintln(&b, box.Render(inner.String()))
	}

	fmt.Fprintln(&b)

	{
		var inner strings.Builder
		fmt.Fprintln(&inner, focus(3)+bold("Formatter Order"))
		for i, f := range m.config.Formatters.Order {
			up, down := "  ", "  "
			if i > 0 {
				up = " ▲"
			}
			if i < len(m.config.Formatters.Order)-1 {
				down = " ▼"
			}
			fmt.Fprintf(&inner, "%s%-20s%s%s", indent, f, up, down)
			if i < len(m.config.Formatters.Order)-1 {
				fmt.Fprintln(&inner)
			}
		}
		fmt.Fprintln(&b, box.Render(inner.String()))
	}

	fmt.Fprintln(&b)

	{
		buttons := []string{"Save", "Cancel"}
		fmt.Fprint(&b, "  ")
		for i, label := range buttons {
			if i > 0 {
				fmt.Fprint(&b, "  ")
			}
			if m.focus == 4 && m.buttonIdx == i {
				fmt.Fprint(&b, accent("[ "+label+" ]"))
			} else {
				fmt.Fprint(&b, "[ "+label+" ]")
			}
		}
	}

	rendered := b.String()
	if m.width > 0 {
		rendered = lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, rendered)
	}

	v := tea.NewView(rendered)
	v.AltScreen = true
	return v
}
