package navigation

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

func Exec(args string) {
	path := traversePath(args)	
	changeDir(path)
}

func traversePath(path string) string {
	var res []string
	pathArr := strings.Split(path, "/")
	
	for i, arg := range(pathArr) {
		if i == 0 {
			// handle abs path	
			if len(arg) == 0 {
				continue
			}

			wd := getWd()
			wdArr := strings.Split(wd[1:], "/")
			res = append(res, wdArr...)	
			
			// if ../, then we need to continue
			if arg == "." {
				continue
			}
		}
	
		if arg == ".." {
			// probably need error handle
			res = res[:len(res)-1]
			continue
		}

		res = append(res, arg)
	}

	return "/" + strings.Join(res, "/")
}

func getWd() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return dir
}

func changeDir(dir string) {
	err := os.Chdir(dir)
	if errors.Is(err, fs.ErrNotExist) {
		fmt.Println("cd: " + dir + ": No such file or directory")	
	} else if err != nil {
		log.Fatal("Failed to change dir")
	}
}
