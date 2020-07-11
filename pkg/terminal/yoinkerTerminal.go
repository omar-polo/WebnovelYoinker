package terminal

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

// StartTerminal start a terminal applcation to use Webnovel Yoinker
func StartTerminal() {
	app := cli.NewApp()
	app.Name = "WebnovelYoinker terminal"
	app.Usage = "Lets you download webnovels and exports them as epub or pdf"

	scrapeFlags := []cli.Flag{
		&cli.PathFlag{
			Name:     "in",
			Value:    "~/",
			Required: true,
		},
		&cli.PathFlag{
			Name:     "out",
			Value:    "~/Downloads",
			Required: true,
		},

		&cli.IntFlag{
			Name:        "r",
			Usage:       "Sets the number of go routines used to scrape, aka how many should be downloaded simultaenously",
			DefaultText: "3",
			Value:       3,
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:   "scrape",
			Usage:  "Downloads specified books and exports them",
			Flags:  scrapeFlags,
			Action: scrapeCommand,
		},
	}

	err := app.Run(os.Args)
	logErr(err)
}

func scrapeCommand(c *cli.Context) error {
	jobs := getBookConfigs(c.String("in"))
	yoinker.BeginMultiConvert(jobs, c.Int("r"), c.String("out"))
	return nil
}

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getBookConfigs(path string) []yoinker.BookMetadata {
	books := []yoinker.BookMetadata{}
	rawBooks, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("%v: %v", err.Error(), err)
	}

	err = yaml.Unmarshal(rawBooks, &books)
	if err != nil {
		log.Fatalf("Unknown format: %v", err)
	}

	return books
}
