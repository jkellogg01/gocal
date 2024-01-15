package main

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
)

func main() {
	data, err := os.ReadFile("test/data/abe-lincoln.ics")
	if err != nil {
		log.Fatal(err)
	}

	iCalComponent, err := Compile(Unfold(string(data)))
	if err != nil {
		log.Fatal(err)
	}
	output, err := json.MarshalIndent(iCalComponent, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	outfile, err := os.Create("./output/test.json")
	if err != nil {
		log.Fatal(err)
	}

	outfile.Write(output)
}

func Unfold(data string) string {
	return regexp.MustCompile(`(?:\r\n )|(?:\r\n\t)`).ReplaceAllString(data, "")
}
