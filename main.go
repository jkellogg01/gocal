package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	data, err := os.ReadFile("test/data/josh-calendar.ics")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(Unfold(string(data)), "\r\n")
	for _, line := range lines {
		cl, err := Scan(line)
		if err != nil {
			log.Fatal(err)
		}
		// if len(cl.Params) > 0 {
		fmt.Println(cl.ToString())
		// }
	}
}

func Unfold(data string) string {
	return regexp.MustCompile(`(?:\r\n )|(?:\r\n\t)`).ReplaceAllString(data, "")
}
