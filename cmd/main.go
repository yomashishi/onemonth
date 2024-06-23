package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/yomashishi/onemonth"
)

var (
	year         int
	month        int
	timeLayout   string
	outputLayout string = "%s\n"
)

func init() {
	var (
		usageYear   = "number of years"
		defaultYear = time.Now().Local().Year()
	)
	flag.IntVar(&year, "year", defaultYear, usageYear)
	flag.IntVar(&year, "y", defaultYear, usageYear+" (shorthand)")

	var (
		usageMonth   = "number of months"
		defaultMonth = int(time.Now().Local().Month())
	)
	flag.IntVar(&month, "month", defaultMonth, usageMonth)
	flag.IntVar(&month, "m", defaultMonth, usageMonth+" (shorthand)")

	var (
		usageTimeLayout   = "string of time format layout."
		defaultTimeLayout = time.DateOnly
	)
	flag.StringVar(&timeLayout, "layout", defaultTimeLayout, usageTimeLayout)
	flag.StringVar(&timeLayout, "l", defaultTimeLayout, usageTimeLayout+" (shorthand)")

	if layout := getConfig("layout.time"); len(layout) > 0 {
		timeLayout = layout
	}
	if layout := getConfig("layout.output"); len(layout) > 0 {
		outputLayout = layout
	}
}

func getConfig(configPath string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	cfgPath := filepath.Join(home, ".config", "onemonth", configPath)

	f, err := os.Open(cfgPath)
	if err != nil {
		return ""
	}
	defer f.Close()

	// read first line
	scanner := bufio.NewScanner(f)
	if scanner.Scan() {
		return strings.ReplaceAll(scanner.Text(), `\n`, "\n")
	} else {
		return ""
	}
}

func main() {
	flag.Parse()

	month, err := onemonth.NewOneMonth(year, month)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	month.Iterate(func(day time.Time) {
		fmt.Printf(outputLayout, day.Format(timeLayout))
	})
}
