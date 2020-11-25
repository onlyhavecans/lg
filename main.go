package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

/*
logtodayone
Based on the ruby version by Brett Terpstra (http://brettterpstra.com)
Use and modify freely, attribution appreciated

This script works with the Day One command line utility
It parses an input string for a [date string] at the
beginning to parse natural language dates

Example usage:
logtodayone "This is a entry."
logtodayone "[yesterday 3pm] Something I did yesterday at 3:00PM"
*/

const (
	dayOneBin  = "/usr/local/bin/dayone2"
	dayOneTime = "2006/01/02 15:04:05"
	exitFail   = 1
)

type DayOne struct {
	DayOneCmd *exec.Cmd
	stdout    io.Writer
}

func (d *DayOne) Log(date time.Time, message string) error {
	dateString := date.Format(dayOneTime)
	dateArg := fmt.Sprintf("--date=\"%s\"", dateString)

	var cmd *exec.Cmd
	if d.DayOneCmd == nil {
		cmd = exec.Command(dayOneBin)
	} else {
		cmd = d.DayOneCmd
	}

	if d.stdout == nil {
		d.stdout = os.Stdout
	}

	cmd.Args = []string{cmd.Path, dateArg, "new", message}
	out, err := cmd.Output()
	if err != nil {
		return err
	}

	outString := strings.TrimSpace(string(out[:]))

	fmt.Fprintf(d.stdout, "%s\n", outString)

	return nil
}

func main() {
	dayOne := DayOne{}

	if err := run(os.Args, dayOne); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(exitFail)
	}

}

func run(args []string, dayOne DayOne) error {
	if len(args) < 2 {
		return errors.New("you need to add the entry on the command line")
	}

	var input string
	input = strings.Join(args[1:], " ")
	input = strings.TrimSpace(input)

	re := regexp.MustCompile(`^\[.*?\]\s*`)
	matched := re.MatchString(input)

	date := time.Now()
	if matched {
		w := when.New(nil)
		w.Add(en.All...)
		w.Add(common.All...)

		ds := re.FindString(input)
		d, err := w.Parse(ds, time.Now())
		if err != nil {
			return err
		}
		if d != nil {
			date = d.Time
			input = re.ReplaceAllString(input, "")
		}
	}

	if err := dayOne.Log(date, input); err != nil {
		return err
	}

	return nil
}
