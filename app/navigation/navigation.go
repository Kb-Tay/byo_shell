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
	path := traverse("", args)
	changeDir(path)
}


// Checks if path starts from abs or relative path 
func getInitPath(path string) (string, string) {
	if path[0] == '.' {
		return path, getWd()	
	}
	
	return "", path
}

func traverse(currPath string, targetPath string) string {
	if len(targetPath) == 0 {
		return currPath
	}

	// only applies if the initial path is / (absolute path)
	if targetPath[0] == '/' {
		if len(targetPath) == 1 { // the path is just '/'
			return "/" 
		}

		if targetPath[1] == '/' {
			return traverse(targetPath[1:], "/")
		}

		s := strings.SplitN(targetPath[1:], "/", 2)
	
		if len(s) == 1 {
			return "/" + s[0] 
		}

		return traverse("/" + s[0], s[1])
	}

	// handling relative directories
	if targetPath[0] == '.' {
		res := strings.SplitN(targetPath, "/", 2)
	
		if res[0] == "." {
			dir := getWd()

			
			return traverse(dir, res[1])
		}

		if res[0] == ".." {
			dir := getWd()
			parentDir := moveToParentDir(dir)
			return traverse(parentDir, res[1])
		}
	}
	
	// '/blah/../../' 
	res := strings.SplitN(targetPath, "/", 2);

	if len(res) > 1 {
		return traverse(res[1], currPath + "/" + res[0])
	}
	
	return currPath + "/" + res[0]
}


func getWd() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return dir
}

// moves from curr dir to parent dir
func moveToParentDir(currDir string) string {
	sArr := strings.Split(currDir, "/")
	return strings.Join(sArr[:len(sArr) - 1], "/")
}


func changeDir(dir string) {
	err := os.Chdir(dir)
	if errors.Is(err, fs.ErrNotExist) {
		fmt.Println("cd: " + dir + ": No such file or directory")	
	}

	if err != nil {
		log.Fatal("Failed to change dir")
	}
}
