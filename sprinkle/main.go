package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	execpath, _ := os.Executable() // 获得程序路径
	file, err := os.Open(filepath.Join(filepath.Dir(execpath), "./transforms.txt"))
	defer file.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	br := bufio.NewReader(file)
	var transforms []string

	for {
		line, _, err := br.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Println("Reading txt error:" + err.Error())
			}
			break
		}
		transforms = append(transforms, string(line))
	}

	if len(transforms) < 1 {
		log.Fatal("Couldn't read any words from txt!")
	}

	// transforms.txt 第一行作标记字符
	otherWord := transforms[0]

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
