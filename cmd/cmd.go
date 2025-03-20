package cmd

import "fmt"

type Command interface {
	Execute() error
}

type LoopCmd struct {
	Parent Command
}

func (c *LoopCmd) Execute() error {
	for {
		if err := c.Parent.Execute(); err != nil {
			return err
		}
	}
}

func Wrap(cmd Command) Command {
	return &StatusWrapperCmd{Parent: cmd}
}

type StatusWrapperCmd struct {
	Parent Command
}

func (s *StatusWrapperCmd) Execute() error {
	if err := s.Parent.Execute(); err != nil {
		fmt.Printf("Команда завершилась ошибкой: %s", err)
	} else {
		fmt.Println("Успешно")
	}
	return nil
}
