package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	app := &cli.App{
		Name: "randomWordchoice",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "separator",
				Aliases:  []string{"sep", "s"},
				Usage:    "string to separate words",
				Value:    " ",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "titlecase",
				Aliases:  []string{"t"},
				Usage:    "capitalise first letter of each word",
				Value:    false,
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				log.Fatalf("Must pass list of files containing words")
			}
			sources := c.Args().Slice()
			str, err := generateRandomWords(sources, c.String("sep"), c.Bool("titlecase"))
			fmt.Println(str)
			return err
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func generateRandomWords(sources []string, sep string, titlecase bool) (string, error) {
	words := make([]string, len(sources))
	for i, path := range sources {
		lines := readLines(path)
		word := chooseRandom(lines)
		if titlecase {
			words[i] = strings.Title(word)

		} else {
			words[i] = word
		}
	}
	return strings.Join(words, sep), nil
}

func readLines(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")
}

func chooseRandom(list []string) string {
	return list[rand.Intn(int(len(list)))]
}
