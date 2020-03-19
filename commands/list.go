package commands

import (
	"fmt"

	"../common"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type ListCommand struct {
	context *cli.Context
}

func (c *ListCommand) Execute(context *cli.Context) error {
	fmt.Println("hello from the list command.")
	return nil
}

func RunFunction(c *cli.Context) error {
	err := &ListCommand{}
	switch c.Command.Name {
	case "all":
		fmt.Println("thats all")
	default:
		fmt.Println("default one")
	}

	if err != nil {
		logrus.Fatal(err)
	}

	return nil
}

func init() {
	common.RegisterCommand2("list", "Just prints right now", &ListCommand{})
}
