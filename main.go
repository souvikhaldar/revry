package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	path := flag.String("path", ".", "path to read the diary files")
	flag.Parse()

	files, err := os.ReadDir(*path)
	if err != nil {
		log.Fatal(err)
	}
	var introspection string
	for _, file := range files {
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
