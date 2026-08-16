package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	for {
		select {
		case <- ctx.Done():
			return	

		default:
			fmt.Print("$ ")
			command, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				fmt.Printf("Failed to read file: %s", err.Error())
				os.Exit(1)	
			}

			fmt.Println(command[:len(command)-1] + ": command not found") // remove the trailing newline
		}
	}
}
