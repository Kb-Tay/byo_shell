package main

import (
	"fmt"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	fmt.Print("$ ")
	var arg string
	fmt.Scanf("%s", &arg)
	fmt.Printf("%s: command not found", arg)
}
