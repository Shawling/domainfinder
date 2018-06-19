package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	execPath, _ := os.Executable()
	var cmdChain = []*exec.Cmd{
		exec.Command(filepath.Join(filepath.Dir(execPath), "./lib/synonyms")),
		exec.Command(filepath.Join(filepath.Dir(execPath), "lib/sprinkle")),
		exec.Command(filepath.Join(filepath.Dir(execPath), "lib/coolify")),
		exec.Command(filepath.Join(filepath.Dir(execPath), "lib/domainify")),
		exec.Command(filepath.Join(filepath.Dir(execPath), "lib/available")),
	}

	cmdChain[0].Stdin = os.Stdin
	cmdChain[len(cmdChain)-1].Stdout = os.Stdout
	for i := 0; i < len(cmdChain)-1; i++ {
		thisCmd := cmdChain[i]
		nextCmd := cmdChain[i+1]
		stdout, err := thisCmd.StdoutPipe()
		if err != nil {
			log.Fatalln(err)
		}
		nextCmd.Stdin = stdout
	}
	for _, cmd := range cmdChain {
		if err := cmd.Start(); err != nil {
			log.Fatalln(err)
		} else {
			// 关闭时关闭所有子程序
			defer cmd.Process.Kill()
		}
	}
	for _, cmd := range cmdChain {
		// 等待子程序结束，否则主程序会提前结束
		if err := cmd.Wait(); err != nil {
			log.Fatalln(err)
		}
	}
}
