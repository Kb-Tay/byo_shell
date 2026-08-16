package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

var commandMap = map[string]string{
	"echo": "echo is a shell builtin",
	"exit": "exit is a shell builtin",
}

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	for {
		select {
		case <- ctx.Done():
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
				if !ok {
					fmt.Println(strings.Join(args[1:], "") + ": command not found")	
					break	
				}

				fmt.Println(info)
			case "exit":
				return
			case "echo":
				fmt.Println(strings.Join(args[1:], " "))
			default:
				fmt.Println(cmd + ": command not found") // remove the trailing newline
			}

		}
	}
}

