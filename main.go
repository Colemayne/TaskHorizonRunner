package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Task Horizon Runner"
	app.Usage = "Let's you configure & interact with Task Horizon"

	// We'll be using the same flag for all our commands
	// so we'll define it up here
	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "message",
			Value: "Default message",
		},
	}

	// the commands
	app.Commands = []*cli.Command{
		{
			Name:  "th",
			Usage: "Just Prints TH",
			Flags: myFlags,
			// the action, or code that will be executed when
			Action: func(c *cli.Context) error {
				fmt.Println("TH")
				return nil
			},
		},
	}

	// starts the application
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
