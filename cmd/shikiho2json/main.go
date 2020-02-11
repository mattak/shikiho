package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mattak/shikiho/pkg/shikiho"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func writeFile(readFile, outDirectory string) {
	journal, err := shikiho.ParseJournal(readFile)
	if err != nil {
		panic(err)
	}

	outfile := path.Join(outDirectory, journal.Code+".json")
	err = ioutil.WriteFile(outfile, []byte(journal.ToJson()), 0644)

	if err != nil {
		panic(err)
	}
}

func runSingle(c *cli.Context) bool {
	if c.Args().Len() < 1 {
		return false
	}

	filepath := c.Args().Get(0)

	journal, err := shikiho.ParseJournal(filepath)
	if err != nil {
		panic(err)
	}

	fmt.Println(journal.ToJson())
	return true
}

func runFromStdin(c *cli.Context) bool {
	directory := c.String("directory")
	if len(directory) < 1 {
		return false
	}

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		println("not exists directory")
		err := os.MkdirAll(directory, 0777)
		if err != nil {
			panic(err)
		}
	}

	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		if err := stdin.Err(); err != nil {
			return true
		}

		readFile := stdin.Text()
		writeFile(readFile, directory)
	}

	return true
}

func main() {
	app := cli.NewApp()
	app.Name = "shikiho"
	app.Usage = "extract shikiho data json from html archive"
	app.Version = "0.0.1"
	app.ArgsUsage = "[html_file_path]"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "directory",
			Aliases: []string{"d"},
			Usage:   "directory to output parse result json",
		},
	}
	app.Action = func(c *cli.Context) error {
		if runFromStdin(c) {
			return nil
		}

		if runSingle(c) {
			return nil
		}

		return errors.New("single mode: <html_path>, multi_mode: -d tmp")
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
