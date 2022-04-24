// package command provides a simple way to create cli applications in go.
package command

import (
	"fmt"
	"os"
)

type Command interface {
	Synopsis() string
	Help() string
	Run(args []string) error
}

type Commander struct {
	name     string
	commands map[string]Command
}

func New(name string) *Commander {
	cmder := &Commander{
		name:     name,
		commands: make(map[string]Command),
	}

	return cmder
}

func (c *Commander) Register(name string, cmd Command) {
	c.commands[name] = cmd
}

func (c *Commander) Run() error {
	if len(os.Args) > 1 {
		for name, cmd := range c.commands {
			if name == os.Args[1] {
				return cmd.Run(os.Args[2:])
			}
		}
	}
	if os.Args[1] == "help" {
		c.Usage()
	}
	return nil
}

func (c *Commander) Usage() {
	if len(os.Args) > 2 {
		cmd, ok := c.commands[os.Args[2]]
		if !ok {
			fmt.Printf("command %s does not exist", os.Args[2])
			return
		}
		fmt.Println(cmd.Help())
		return
	}

	fmt.Printf("  %s <command> [args]\n\n", c.name)
	for name, cmd := range c.commands {
		fmt.Printf("\t%s\t%s\n", name, cmd.Synopsis())
	}
}
