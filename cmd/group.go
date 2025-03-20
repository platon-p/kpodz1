package cmd

import (
	"fmt"
)

type NamedCommand interface {
	Command
	fmt.Stringer
}

func Named(cmd Command, name string) *NamedCommandImpl {
	return &NamedCommandImpl{
		Cmd:  cmd,
		Name: name,
	}
}

type NamedCommandImpl struct {
	Cmd  Command
	Name string
}

func (n *NamedCommandImpl) Execute() error {
	return n.Cmd.Execute()
}

func (n *NamedCommandImpl) String() string {
	return n.Name
}

type GroupCmd struct {
	Commands  []NamedCommand
	ExitLabel string
}

func (c *GroupCmd) Execute() error {
	cmd, err := c.selector()
	if err != nil {
		return err
	}
	return cmd.Execute()
}

func (c *GroupCmd) selector() (Command, error) {
	for i := range c.Commands {
		fmt.Printf("%v. %s\n", i+1, c.Commands[i])
	}
	fmt.Println()
	fmt.Print("Выберите номер команды: ")
	var id uint
	if _, err := fmt.Scan(&id); err != nil {
		return nil, err
	}
	if id <= 0 || id > uint(len(c.Commands)) {
		return nil, fmt.Errorf("out of range")
	}
	return c.Commands[id-1], nil
}

type SimpleCmd struct{ fun func() error }

func NewSimpleCmd(fun func() error) *SimpleCmd {
	return &SimpleCmd{fun: fun}
}

func (e *SimpleCmd) Execute() error {
	return e.fun()
}
