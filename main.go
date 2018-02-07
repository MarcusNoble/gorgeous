package main

import (
	"bufio"
	"fmt"
	"os"

	"flag"
	"strings"

	"github.com/ladydascalie/gorgeous/filters"
)

var prefix string

func main() {
	flag.StringVar(&prefix, "prefix", "", "give a prefix you want to automatically strip away (useful with docker-compose logs)")
	flag.Parse()

	//multi := io.MultiReader(os.Stdin, os.Stderr)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		if prefix != "" {
			txt = strings.TrimPrefix(txt, prefix)
		}
		for _, f := range filters.All {
			if t := f(txt); t != "" {
				fmt.Println(t)
			}
		}
	}
}
