package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	fYear        = flag.Uint("year", 0, "specifies year")
	fDay         = flag.Uint("day", 0, "specifies day")
	fBasedir     = flag.String("basedir", "../cpp", "base directory of C++ solutions")
	fTemplatedir = flag.String("templatedir", "templates/cpp", "directory of template files for C++ solutions")
)

type templateData struct {
	Year      uint
	FullYear  string
	Day       uint
	PaddedDay string
	FullDay   string
}

func imain() int {
	flag.Parse()

	year := *fYear
	if year == 0 {
		fmt.Println("a year must be specified with the -year flag")
		return 1
	}
	if year < 2015 || year > 2019 {
		fmt.Printf("invalid year %d: the year must be a year on which an AoC event was held\n", year)
		return 1
	}

	day := *fDay
	if day == 0 {
		fmt.Println("a day must be specified with the -day flag")
		return 1
	}
	if day > 25 {
		fmt.Printf("invalid day %d: the day must be in the range 1-25 (both inclusive)\n", day)
		return 1
	}

	basedir := *fBasedir
	if basedir == "" {
		fmt.Printf("a base directory for C++ solutions must be specified with the -basedir flag\n")
		return 1
	}

	templatedir := *fTemplatedir
	if templatedir == "" {
		fmt.Printf("a directory with template files for C++ solutions must be specified with the -templatedir flag\n")
		return 1
	}

	data := templateData{
		Year:      year,
		FullYear:  fmt.Sprintf("year%d", year),
		Day:       day,
		PaddedDay: fmt.Sprintf("%02d", day),
		FullDay:   fmt.Sprintf("day%02d", day),
	}
	fmt.Printf("template data: %+v\n", data)
	fmt.Printf("basedir: %s\n", basedir)
	fmt.Printf("templatedir: %s\n", templatedir)

	return 0
}

func main() {
	os.Exit(imain())
}
