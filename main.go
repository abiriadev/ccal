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

func printMonth(today time.Time) {
	year, month, _ := today.Date()
	fmt.Println(center(fmt.Sprintf("%s %d", month, year), WIDTH))
	fmt.Println("Su Mo Tu We Th Fr Sa")

	daysin, weekoffset := daysin(year, month), weekoffset(year, month)

	bgred := color.New(color.BgRed).SprintFunc()

	for ofs, d := 0, 1; ; ofs++ {
		if ofs < int(weekoffset) {
			fmt.Printf("  ")
		} else {
			if d == today.Day() {
				fmt.Printf("%s", bgred(fmt.Sprintf("%2d", d)))
			} else {
				fmt.Printf("%2d", d)
			}
			d++
			if d > daysin {
				fmt.Printf(strings.Repeat(" ", 3*((WEEK-1)-(ofs%WEEK))))
				break
			}
		}

		if ofs%WEEK == WEEK-1 {
			fmt.Println()
		} else {
			fmt.Printf(" ")
		}
	}

	fmt.Printf("\n%s\n", strings.Repeat(" ", WIDTH))
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

	printMonth(t)
}
