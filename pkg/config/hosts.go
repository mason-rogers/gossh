package config

import "strings"

func (c Config) FindJumpHostByName(name string) *Host {
	for _, jh := range c.JumpHosts {
		if jh.Name == name {
			return &jh
		}
	}
	return nil
}

func (c Config) FindHostByPath(path string) *Host {
	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		return nil
	}

	currentGroups := c.Groups
	hostName := parts[len(parts)-1]

	// Navigate through groups
	for _, groupName := range parts[:len(parts)-1] {
		found := false
		for _, group := range currentGroups {
			if group.Name == groupName {
				currentGroups = group.Groups
				// Check hosts at this level
				for _, host := range group.Hosts {
					if host.Name == hostName {
						return &host
					}
				}
				found = true
				break
			}
		}
		if !found {
			return nil
		}
	}

	return nil
}
