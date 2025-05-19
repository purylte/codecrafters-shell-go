package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	for true {
		fmt.Fprint(os.Stdout, "$ ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cleanedCommand := command[:len(command)-1]
		if cleanedCommand == "exit 0" {
			os.Exit(0)
		}
		fmt.Println(cleanedCommand + ": command not found")
	}
}
