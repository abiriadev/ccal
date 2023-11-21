package main

import (
	"fmt"
	"strings"
	"time"
)

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
	fmt.Println(center(fmt.Sprintf("%s %d", month, year), 7*2+6))
	fmt.Println("Su Mo Tu We Th Fr Sa")

	daysin, weekoffset := daysin(year, month), weekoffset(year, month)

	for ofs, d := 0, 1; ; ofs++ {
		if ofs < int(weekoffset) {
			fmt.Printf("  ")
		} else {
			fmt.Printf("%2d", d)
			d++
			if d > daysin {
				break
			}
		}

		if ofs%7 == 6 {
			fmt.Println()
		} else {
			fmt.Printf(" ")
		}
	}

	fmt.Printf("\n")
}

func main() {
	t := time.Now()

	printMonth(t)
}
