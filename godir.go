package main

import (
	"fmt"
	"os"
	"path/filepath"

	putil "github.com/Masterminds/godir/pathutil"
	"github.com/codegangsta/cli"
)

var version = "0.0.1"

var Description = `godir - The Go Path Tool

This tool is a utility for working with paths that are significant for Go.

Unless otherwise noted, the tools here are designed to ignore directories that
Go treats as special, including prefixed directories ("_scripts/"), test
directories ("testdata/") and the vendor directory ("vendor/").

EXAMPLES:

Run "go test" without running the tests in vendor:

	$ go test $(godir pkgs)

Given a case where $GOPATH has more than one path, show which path the current
code lives on.

	$ GOPATH=~/Code/Go:~/Code/tmp godir gopath
	/Users/mbutcher/Code/Go
`

func main() {
	app := cli.NewApp()
	app.Name = "godir"
	app.Usage = Description
	app.Version = version
	app.Action = func(c *cli.Context) {
		sp := putil.Subpaths(wdir(), true)
		for _, p := range sp {
			fmt.Println(putil.Name(p))
		}
	}
	commands(app)
	app.Run(os.Args)
}

func commands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:      "name",
			Usage:     "Print the name of this package, relative to $GOPATH.",
			ArgsUsage: "[PATH]",
			Action: func(c *cli.Context) {
				fmt.Println(putil.Name(argOrWdir(c)))
			},
		},
		{
			Name:      "gopath",
			Usage:     "Print the path in $GOPATH that the given package was found on.\n\tThis is for cases where $GOPATH has numerous paths.",
			ArgsUsage: "[PATH]",
			Action: func(c *cli.Context) {
				fmt.Println(putil.WhichGopath(argOrWdir(c)))
			},
		},
		{
			Name:  "pkgs",
			Usage: "Print all packages (that contain Go code) from the current directory. Skip vendor/",
			Action: func(c *cli.Context) {
				wd := argOrWdir(c)
				sp := putil.Subpaths(wd, true)
				for _, p := range sp {
					fmt.Println(putil.Name(p))
				}
			},
		},
		{
			Name:  "paths",
			Usage: "Print all subpaths from the current directory.",
			Action: func(c *cli.Context) {
				wd := argOrWdir(c)
				sp := putil.Subpaths(wd, false)
				r := c.Bool("relative")
				for _, p := range sp {
					if r {
						rp, err := filepath.Rel(wd, p)
						if err == nil {
							p = rp
						}
					}
					fmt.Println(p)
				}
			},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "relative,r",
					Usage: "Print paths relative to the directory.",
				},
			},
		},
	}
}

func argOrWdir(c *cli.Context) string {
	a := c.Args()
	if len(a) < 1 {
		return wdir()
	}
	p, err := filepath.Abs(a[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get absolute path of %s: %s", a[0], err)
		os.Exit(1)
	}
	return p
}

func wdir() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get working directory: %s.\n", err)
		os.Exit(1)
	}
	return wd
}
