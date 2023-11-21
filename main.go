package main

import (
	"errors"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

const WEEK = 7
const MAX_WEEKS = 6
const WIDTH = WEEK*3 - 1

func daysin(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func weekoffset(year int, month time.Month) time.Weekday {
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Weekday()
}

func center(text string, pad int) string {
	if len(text) >= pad {
		return text
	}

	var buf strings.Builder

	rest := pad - len(text)
	l, r := rest-rest/2, rest/2

	for _, p := range []string{
		strings.Repeat(" ", l),
		text,
		strings.Repeat(" ", r),
	} {
		buf.WriteString(p)
	}

	return buf.String()
}

func printMonth(today time.Time, labels []Label) []string {
	bgred := color.New(color.BgRed).SprintFunc()

	year, month, _ := today.Date()

	daysin, weekoffset := daysin(year, month), int(weekoffset(year, month))

	buf := make([]string, 0, MAX_WEEKS+2)

	buf = append(buf, center(fmt.Sprintf("%s %d", month, year), WIDTH))
	buf = append(buf, "Su Mo Tu We Th Fr Sa")

	grid := make([]string, WEEK*MAX_WEEKS)

	for i := range grid {
		grid[i] = "  "
	}

	for d := 1; d <= daysin; d++ {
		var c string
		if d == today.Day() {
			c = bgred(fmt.Sprintf("%2d", d))
		} else {
			c = fmt.Sprintf("%2d", d)
		}
		grid[weekoffset+d-1] = c
	}

	for i := 0; i < 6; i++ {
		buf = append(buf, strings.Join(grid[i*WEEK:(i+1)*WEEK], " "))
	}

	return buf
}

// inclusive range with color
type Label struct {
	start int
	end   int
	color string
}

var labelRegex = regexp.MustCompile(`^(\d+)-(\d+)$`)

func parseLabel(arg string) (Label, error) {
	m := labelRegex.FindStringSubmatch(arg)

	if len(m) != 3 {
		return Label{}, errors.New("invalid label format")
	}

	start, err := strconv.Atoi(m[1])
	if err != nil {
		return Label{}, err
	}

	end, err := strconv.Atoi(m[2])
	if err != nil {
		return Label{}, err
	}

	color := "red"

	return Label{
		start,
		end,
		color,
	}, nil
}

func main() {
	flag.Parse()

	args := flag.Args()

	labels := make([]Label, 0, len(args))

	for _, arg := range args {
		l, err := parseLabel(arg)
		if err != nil {
			panic(err)
		}
		labels = append(labels, l)
	}

	t := time.Now()

	calbuf := printMonth(t, labels)

	fmt.Printf("%s\n", strings.Join(calbuf, "\n"))
}
