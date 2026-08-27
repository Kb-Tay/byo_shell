package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/navigation"
	"github.com/codecrafters-io/shell-starter-go/app/redirect"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

var commandMap = map[string]string{
	"echo":  "echo is a shell builtin",
	"exit":  "exit is a shell builtin",
	"type":  "type is a shell builtin",
	"pwd": "pwd is a shell builtin",
}

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	for {
		select {
		case <-ctx.Done():
			return

		default:
			fmt.Print("$ ")
			input, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				fmt.Printf("Failed to read file: %s", err.Error())
				os.Exit(1)
			}

			args := strings.Split(input[:len(input)-1], " ")
			cmd := args[0]

			switch cmd {
			case "type":
				if len(args) < 1 {
					continue
				}

				info, ok := commandMap[args[1]]
				if ok {
					fmt.Println(info)
					continue
				}

				isInPath, path := locateExecInPath(args[1])

				if isInPath {
					fmt.Println(args[1] + " is " + path)
					continue
				}

				fmt.Println(strings.Join(args[1:], "") + ": not found")
			case "exit":
				return
		
			case "cd":
				if len(args) < 1 {
					continue
				}
				navigation.Exec(args[1])

			case "pwd":
				pwd, err := os.Getwd()
				if err != nil {
					log.Fatal("Failed to get working dir")
				}

				fmt.Println(pwd)
			case "echo":
				if len(args) == 1 {
					fmt.Println("")
					continue
				}
				
				if redirect, out, _ := redirect.IsRedirect(args[1:]); redirect {
					os.WriteFile(out, []byte(args[1]), 0644)
					continue	
				}

				fmt.Println(strings.Join(args[1:], " "))
			default:
				isInPath, _ := locateExecInPath(cmd)

				if isInPath {
					execCommand(cmd, args[1:])
					continue
				}

				fmt.Println(cmd + ": command not found") // remove the trailing newline
			}

		}
	}
}

func execCommand(cmd string, args []string) {
	var command *exec.Cmd 	

	if len(args) == 0 {
		command = exec.Command(cmd)
	}

	isRedirect, pathToWrite, i := redirect.IsRedirect(args)

	// for eg: ls > hello.txt
	if isRedirect && i == 0 {
		command = exec.Command(cmd)
	}

	if isRedirect && i > 0 {
		command = exec.Command(cmd, args[:i]...)
	}

	if !isRedirect {
		command = exec.Command(cmd, args...)	
	}

	var out strings.Builder
	command.Stdout = &out
	err := command.Run()

	if err != nil {
		fmt.Println(err.Error())
	}

	if isRedirect {
		os.WriteFile(pathToWrite, []byte(out.String()), 0644)
	} else {
		fmt.Println(out.String())
	}
}



// resolve the PATH
func locateExecInPath(targetFile string) (bool, string) {
	path, _ := os.LookupEnv("PATH")
	dirs := strings.SplitSeq(path, ":")

	for dir := range dirs {
		if dir == "" {
			dir = "."
		}

		filePath := filepath.Join(dir, targetFile)
		info, err := os.Stat(filePath)

		if err != nil {
			continue
		}

		if info.Mode().Perm()&0111 != 0 {
			return true, filePath
		}
	}

	return false, ""
}
