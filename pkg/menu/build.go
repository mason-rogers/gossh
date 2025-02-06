package menu

import "github.com/mason-rogers/gossh/pkg/config"

type MenuItem struct {
	Name     string
	IsGroup  bool
	Host     *config.Host
	Children []MenuItem
}

func buildMenuItems(groups []config.Group, prefix string) []MenuItem {
	var menu []MenuItem
	for _, group := range groups {
		groupName := group.Name
		if prefix != "" {
			groupName = prefix + "/" + groupName
		}

		// Process child groups first
		var subItems []MenuItem
		if len(group.Groups) > 0 {
			subItems = buildMenuItems(group.Groups, groupName)
		}

		// Only add this group if it has direct hosts
		if len(group.Hosts) > 0 {
			item := MenuItem{
				Name:    groupName,
				IsGroup: true,
			}

			// Add hosts from this group
			for _, host := range group.Hosts {
				item.Children = append(item.Children, MenuItem{
					Name: groupName + "/" + host.Name,
					Host: &host,
				})
			}

			menu = append(menu, item)
			menu = append(menu, item.Children...)
		}

		// Add subgroup items
		menu = append(menu, subItems...)
	}

	return menu
}
