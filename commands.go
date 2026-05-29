package main

import (
	"errors"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(cmd command) error
}

func (c *commands) run(cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("Command is not registered")
	}

	return f(cmd)
}

func (c *commands) register(name string, f func(cmd command) error) {
	c.registeredCommands[name] = f
}
