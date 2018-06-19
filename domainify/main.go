package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"
)

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func main() {
	execpath, _ := os.Executable()
	file, err := os.OpenFile(filepath.Join(filepath.Dir(execpath), "./tlds.txt"), os.O_APPEND|os.O_RDWR, os.ModeAppend)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	br := bufio.NewReader(file)
	var tlds []string

	for {
		line, _, err := br.ReadLine()
		if err != nil {
			if err != io.EOF {
				fmt.Println("Reading file error: " + err.Error())
			}
			break
		}
		tlds = append(tlds, string(line))
	}

	if len(tlds) < 1 {
		fmt.Println("No tlds!")
		os.Exit(1)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune
		for _, r := range text {
			if unicode.IsSpace(r) {
				r = '-'
			}
			if !strings.ContainsRune(allowedChars, r) {
				continue
			}
			newText = append(newText, r)
		}
		fmt.Println(string(newText) + tlds[rand.Intn(len(tlds))])
	}
}
