package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	for true {
		fmt.Fprint(os.Stdout, "$ ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cleanedCommand := command[:len(command)-1]
		first, rest, success := strings.Cut(cleanedCommand, " ")
		if !success {
			first = cleanedCommand
		}
		switch first {
		case "exit":
			if rest == "0" {
				os.Exit(0)
			}
		case "echo":
			fmt.Println(rest)
		default:
			fmt.Println(cleanedCommand + ": command not found")
		}
	}
}
