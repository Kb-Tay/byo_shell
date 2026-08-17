package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

var commandMap = map[string]string{
	"echo": "echo is a shell builtin",
	"exit": "exit is a shell builtin",
	"type": "type is a shell builtin",
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
					if resolvePath(args[1], true) {
						break
					}
					fmt.Println(strings.Join(args[1:], "") + ": not found")	
					break	
				}

				fmt.Println(info)
			case "exit":
				return
			case "echo":
				fmt.Println(strings.Join(args[1:], " "))
			default:
				if !isCommandBuiltIn(cmd) && resolvePath(cmd, false) {
					command := exec.Command(cmd, args[1:]...)
					var out strings.Builder
					command.Stdout = &out
					command.Run()
//					stdout, err := command.StdoutPipe()
//					if err != nil {
//						return 
//					}
//					command.Start()
//
//					bytes, err := io.ReadAll(stdout)
//					if err != nil {
//						return	
//					}
//					command.Wait()
					fmt.Print(out.String())	
					continue
				}

				fmt.Println(cmd + ": command not found") // remove the trailing newline
			}

		}
	}
}

func isCommandBuiltIn(cmd string) bool {
	_, ok := commandMap[cmd]
	return ok
}

// resolve the PATH
func resolvePath(targetFile string, isPrint bool) bool {
	path, _ := os.LookupEnv("PATH")
	dirs := strings.SplitSeq(path, ":")

	for dir := range dirs {
		filePath, isFound := traverseDirs(dir, targetFile)
	
		if isFound {
			if isPrint {
				fmt.Println(targetFile + " is " + filePath)
			}
			return true
		}
	}
	// traversal of folders
	// check if folder or file, if file then compare the filename. If not recursely call read on the folder 
	return false
}

func traverseDirs(dir string, target string) (string, bool) {
	entries, err := os.ReadDir(dir)

	if err != nil {
		return "", false	
	}
	
	for _, entry := range entries {
		if entry.IsDir() {
			filePath, isFound := traverseDirs(dir + entry.Name(), target)

			if isFound {
				return filePath, true
			}
		}
	
		// check for file execution permission
		if entry.Name() == target && isFileExec(entry){
			return dir + "/" + entry.Name(), true 
		}
	}

	return "", false 
}

func isFileExec(dirEntry os.DirEntry) bool {
	fi, _ := dirEntry.Info()
	return fi.Mode().Perm()&0111 != 0
}


