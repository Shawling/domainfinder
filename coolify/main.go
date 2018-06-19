package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	duplicate bool = true
	remove         = false
)

func randBool() bool {
	return rand.Intn(2) == 0
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := []byte(s.Text())
		if randBool() {
			var cIndex = -1
			for i, c := range text {
				switch c {
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
					if randBool() {
						cIndex = i
					}
				}
			}
			if cIndex >= 0 {
				switch randBool() {
				case duplicate:
					//...表示将slice 的每个元素作为多个参数传入
					text = append(text[:cIndex+1], text[cIndex:]...)
				case remove:
					text = append(text[:cIndex], text[cIndex+1:]...)
				}
			}
		}
		fmt.Println(string(text))
	}
}
