package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	bg := &BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := bg.Synonyms(word)
		if err != nil {
			log.Printf("Failed when looking for synonyms for %s: %s\n", word, err.Error())
			continue
		}
		syns = append([]string{word}, syns...)
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
