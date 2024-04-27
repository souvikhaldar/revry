package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	path := flag.String("path", ".", "path to read the diary files")
	period := flag.String("period", "week", "time period of collecting journal entries")
	flag.Parse()
	currTime := time.Now()
	var limit time.Time
	if *period == "week" {
		limit = currTime.Add(-7 * 24 * time.Hour)
	} else {
		limit = currTime.Add(-31 * 24 * time.Hour)
	}

	layout := "02-01-2006"

	files, err := os.ReadDir(*path)
	if err != nil {
		log.Fatal(err)
	}

	var introspection string
	for _, file := range files {
		fileName := file.Name()
		fileName = strings.TrimSuffix(fileName, ".md")
		// Parse the string into a time.Time object
		parsedTime, err := time.Parse(layout, fileName)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return
		}
		if parsedTime.Before(limit) {
			continue
		}
		content, err := os.ReadFile(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		introspection += "\n\n"
		introspection += file.Name()
		introspection += "\n\n"
		introspection += string(content)

	}
	fmt.Println(introspection)

}
