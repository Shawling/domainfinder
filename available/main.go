package main

import (
	"bufio"
	"fmt"
	"os"
)

var marks = map[bool]string{true: "✅", false: "❎"}

func main() {
	apiKey := os.Getenv("WHOIS_APIKEY")
	whois := &WhoisAPI{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		domain := s.Text()
		available, err := whois.Available(domain)
		if err != nil {
			fmt.Printf("Failed while Checking is domain %s available: %s\n", domain, err.Error())
			continue
		}
		fmt.Println(domain + " " + marks[available])
	}
}
