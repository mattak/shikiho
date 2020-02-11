package main

import (
	"errors"
	"fmt"
	"github.com/mattak/shikiho/pkg/shikiho"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "shikiho"
	app.Usage = "extract shikiho data json from html archive"
	app.Version = "0.0.1"
	app.ArgsUsage = "[html_file_path]"
	app.Action = func(c *cli.Context) error {
		if len(c.Args()) < 1 {
			return errors.New("ERROR: argument requires 1 html file path")
		}
		filepath := c.Args().Get(0)

		journal, err := shikiho.ParseJournal(filepath)
		if err != nil {
			panic(err)
		}

		fmt.Println(journal.ToJson())
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
