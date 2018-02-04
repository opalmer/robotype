package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

var (
	delay = flag.Duration(
		"delay", time.Second*5,
		"The delay between reading the file and typing.")
	notrim = flag.Bool(
		"no-trim", false,
		"If provided, don't trim whitespace. Note, this only applies to file based input.")
	clipboard = flag.Bool(
		"clipboard", false,
		"If provided, read input from the clipboard instead of a file.")
	keydelay = flag.Duration("key-delay", time.Millisecond * 90, "The time between each key press")
)

func main() {
	flag.Parse()
	var text string

	if *clipboard {
		txt, err := robotgo.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		text = txt

	} else {
		if len(flag.Args()) != 1 {
			fmt.Printf("usage: %s <filepath>\n", os.Args[0])
			os.Exit(1)
		}

		file, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}

		contents, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		text = string(contents)
		if !*notrim {
			text = strings.TrimSpace(text)
		}
	}

	fmt.Printf("Sleeping %v seconds...\n", delay.Seconds())
	time.Sleep(*delay)

	for _, character := range text {
		robotgo.TypeStr(string(character))
		time.Sleep(*keydelay)
	}
}
