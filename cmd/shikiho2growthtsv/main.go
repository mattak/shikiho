package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mattak/shikiho/pkg/shikiho"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
)

func runFromStdin(c *cli.Context) bool {
	stdin := bufio.NewScanner(os.Stdin)

	for stdin.Scan() {
		if err := stdin.Err(); err != nil {
			return true
		}

		filepath := stdin.Text()
		journal, err := shikiho.ParseJournal(filepath)
		if err != nil {
			panic(err)
		}

		sorter := shikiho.PerformanceSorter{Code: journal.Code, Performances: journal.Performances}
		sort.Sort(&sorter)
		fmt.Println(sorter.ToTsv())
	}

	return true
}

func main() {
	app := cli.NewApp()
	app.Name = "shikiho2growthtsv"
	app.Usage = "extract growth tsv from shikiho archive"
	app.Version = "0.0.1"
	app.ArgsUsage = "[html_file_path]"
	app.Action = func(c *cli.Context) error {
		if runFromStdin(c) {
			return nil
		}

		return errors.New("single mode: <html_path>, multi_mode: -d tmp")
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
