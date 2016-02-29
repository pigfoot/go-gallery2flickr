package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var (
	appName         = "Gallery2Flickr"
	appUsage        = "Migrate gallery to flickr"
	appVersion      = "0.0.1alpha"
	appAuthor       = "pigfoot"
	appEmail        = "pigfoot@gmail.com"
	appCommands     = appMainCommands
	appMainCommands = []cli.Command{
		{
			Name:    "scan",
			Aliases: []string{"s"},
			Usage:   "Scanning Path",
			Action:  scanPath,
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appUsage
	app.Version = appVersion
	app.Author = appAuthor
	app.Email = appEmail
	app.Commands = appCommands

	app.Run(os.Args)
}

func scanPath(c *cli.Context) {
	if len(c.Args()) == 0 {
		fmt.Printf("please pass scanning path.\n")
		return
	}
	path := c.Args()[0]
	fmt.Printf("%s\n", path)
}
