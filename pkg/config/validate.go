package config

import "fmt"

func (c Config) Validate() []string {
	var errors []string

	for index, host := range c.JumpHosts {
		errors = append(errors, validateJumpHost(host, fmt.Sprintf("jumphosts.%d", index))...)
	}

	for index, group := range c.Groups {
		errors = append(errors, validateGroup(group, fmt.Sprintf("groups.%d", index))...)
	}

	return errors
}

func validateJumpHost(host Host, path string) []string {
	var errors []string

	if host.Name == "" {
		errors = append(errors, fmt.Sprintf("Jump Host (%s) is missing 'name'", path))
	}

	if host.Host == "" {
		errors = append(errors, fmt.Sprintf("Jump Host (%s) is missing 'host'", path))
	}

	return errors
}

func validateGroup(group Group, path string) []string {
	var errors []string

	if group.Name == "" {
		errors = append(errors, fmt.Sprintf("Group (%s) is missing 'name'", path))
	}

	if len(group.Groups) == 0 && len(group.Hosts) == 0 {
		errors = append(errors, fmt.Sprintf("Group (%s) must have either 'groups' or 'hosts'", path))
	}

	for index, host := range group.Hosts {
		errors = append(errors, validateHost(host, fmt.Sprintf("%s.hosts.%d", path, index))...)
	}

	for index, subGroup := range group.Groups {
		errors = append(errors, validateGroup(subGroup, fmt.Sprintf("%s.groups.%d", path, index))...)
	}

	return errors
}

func validateHost(host Host, path string) []string {
	var errors []string

	if host.Name == "" {
		errors = append(errors, fmt.Sprintf("Host (%s) is missing 'name'", path))
	}

	if host.Host == "" {
		errors = append(errors, fmt.Sprintf("Host (%s) is missing 'host'", path))
	}

	if host.JumpHost != "" && Get().FindJumpHostByName(host.JumpHost) == nil {
		errors = append(errors, fmt.Sprintf("Host (%s) specified jump host (%s) does not exist", path, host.JumpHost))
	}

	return errors
}
