// package command provides a simple way to create cli applications in go.
package command

import (
	"fmt"
	"os"
)

type Command interface {
	Synopsis() string
	Help() string
	Run() error
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
		for k, cmd := range c.commands {
			if k == os.Args[1] {
				if len(os.Args) > 2 {
					if os.Args[2] == "-h" || os.Args[2] == "--help" {
						fmt.Println(cmd.Help())
						return nil
					}
				}
				return cmd.Run()
			}
		}
	}
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		c.Usage()
	}
	return nil
}

func (c *Commander) Usage() {
	fmt.Printf("  %s <command> [args]\n\n", c.name)
	for _, cmd := range c.commands {
		fmt.Printf("\t%s", cmd.Synopsis())
	}
}
