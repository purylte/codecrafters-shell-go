package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Shell struct {
	commands map[string]func(string)
}

func NewShell() *Shell {
	s := &Shell{commands: make(map[string]func(string))}
	s.commands["exit"] = s.exit
	s.commands["echo"] = s.echo
	s.commands["type"] = s.typeCmd
	return s
}

// Run starts the shell loop.
func (s *Shell) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$ ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		parts := strings.SplitN(input, " ", 2)
		cmd := parts[0]
		var args string
		if len(parts) > 1 {
			args = parts[1]
		}
		if cmdFunc, exists := s.commands[cmd]; exists {
			cmdFunc(args)
		} else {
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}
}

// exit terminates the shell.
func (s *Shell) exit(arg string) {
	code := 0
	if arg != "" {
		if n, err := strconv.Atoi(arg); err == nil {
			code = n
		} else {
			fmt.Fprintln(os.Stderr, "invalid exit code, defaulting to 0")
		}
	}
	os.Exit(code)
}

// echo prints the arguments.
func (s *Shell) echo(arg string) {
	fmt.Println(arg)
}

// typeCmd checks if a given command is a builtin.
func (s *Shell) typeCmd(arg string) {
	if _, exists := s.commands[arg]; exists {
		fmt.Printf("%s is a shell builtin\n", arg)
	} else {
		fmt.Printf("%s: not found\n", arg)
	}
}

func main() {
	shell := NewShell()
	shell.Run()
}
