package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func contains(s string, list []string) bool {
	for _, i := range list {
		if strings.Contains(s, i) {
			return true
		}
	}
	return false
}

func main() {
	flag.Parse()
	words := flag.Args()
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if !contains(text, words) {
			fmt.Print(text)
		}
	}
}
