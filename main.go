package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"flag"
	"strings"

	"github.com/MarcusNoble/gorgeous/filters"
)

var prefix string
var isJSON bool

// testOutput represents the JSON output of a test.
type testOutput struct {
	Time    string `json:"Time"`
	Action  string `json:"Action"`
	Package string `json:"Package"`
	Output  string `json:"Output"`
}

func main() {
	flag.StringVar(&prefix, "prefix", "", "give a prefix you want to automatically strip away (useful with docker-compose logs)")
	flag.BoolVar(&isJSON, "json", false, "pass this flag if you are dealing with JSON test output (defaults to false)")
	flag.Parse()

	hasFailed := false
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch isJSON {
		case true:
			var out testOutput
			if err := json.Unmarshal(scanner.Bytes(), &out); err != nil {
				log.Fatalln(err)
			}
			txt := strings.TrimSuffix(out.Output, "\n")
			if t := filters.Fail(txt); t != "" {
				hasFailed = true
			}
			for _, f := range filters.All {
				if t := f(txt); t != "" {
					fmt.Println(t)
				}
			}
		default:
			txt := scanner.Text()
			if prefix != "" {
				txt = strings.TrimPrefix(txt, prefix)
			}
			if t := filters.Fail(txt); t != "" {
				hasFailed = true
			}
			for _, f := range filters.All {
				if t := f(txt); t != "" {
					fmt.Println(t)
				}
			}
		}
	}

	if hasFailed {
		os.Exit(1)
	}
}
