package menu

import (
	"github.com/manifoldco/promptui"
	"github.com/mason-rogers/gossh/pkg/config"
	"github.com/pkg/errors"
)

func PromptForHost() (config.Host, error) {
	menuItems := buildMenuItems(config.Get().Groups, "")

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "→ {{ if .IsGroup }}📁{{ else }}   🖥{{ end }} {{ .Name | cyan }}",
		Inactive: "  {{ if .IsGroup }}📁{{ else }}   🖥{{ end }} {{ .Name }}",
		Selected: "→ {{ .Name | green }}",
	}

	prompt := promptui.Select{
		Label:     "Select host",
		Items:     menuItems,
		Templates: templates,
		Size:      15,
	}

	idx, _, err := prompt.Run()
	if err != nil {
		return config.Host{}, err
	}

	selected := menuItems[idx]
	if selected.IsGroup {
		return config.Host{}, errors.New("Please select a host, not a group.")
	}

	return *selected.Host, nil
}
