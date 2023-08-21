package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var verbose = flag.Bool("v", false, "verbose mode")

func parseTime(t string) (time.Time, error) {
	var layout string

	if strings.ContainsAny(t, "amAMpmPM") {
		if strings.Count(t, ":") == 2 {
			layout = "3:04:05pm"
		} else {
			layout = "3:04pm"
		}
	} else {
		if strings.Count(t, ":") == 2 {
			layout = "15:04:05"
		} else {
			layout = "15:04"
		}
	}

	return time.Parse(layout, t)
}

func getNextTargetTime(parsedTime time.Time) time.Time {
	now := time.Now()
	target := time.Date(now.Year(), now.Month(), now.Day(), parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), 0, time.Local)

	if target.Before(now) {
		target = target.Add(24 * time.Hour)
	}

	return target
}

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatalf("Usage: %s [-v] <time in format H:MM[:SS][am/pm] or HH:MM[:SS] (24h)>", os.Args[0])
	}

	parsedTime, err := parseTime(flag.Arg(0))
	if err != nil {
		log.Fatalf("Failed to parse time: %v", err)
	}

	targetTime := getNextTargetTime(parsedTime)

	if *verbose {
		duration := time.Until(targetTime)
		hours := duration / time.Hour
		minutes := (duration % time.Hour) / time.Minute
		seconds := (duration % time.Minute) / time.Second

		fmt.Printf("Will wait for %d hours %d minutes %d seconds until %s\n", hours, minutes, seconds, targetTime.Format("3:04:05pm"))
	}

	for {
		current := time.Now()
		if current.After(targetTime) || current.Equal(targetTime) {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
