package navigation

import (
	"testing"
)

const PWD = "/Users/aaron/Documents/codecrafters/codecrafters-shell-go/app/navigation"

func TestCorrectParsing(t *testing.T) {
	const test = "./"
	const ans = PWD
	result := traverse(PWD, test)

	if ans != result {
t.Errorf("Expect: %v, Actual %v", ans, result)
	}
}

func TestCurrDirAndTraverse(t *testing.T) {
	const test = "./test"
	const ans = PWD + "/test"

	result := traverse(PWD, test)
	if ans != result {
		t.Errorf("Expect: %v, Actual %v", ans, result)
	}
}

func TestParentDir(t *testing.T) {
	const test = "../"
	const ans = "/Users/aaron/Documents/codecrafters/codecrafters-shell-go/app"

	result := traverse(PWD, test)
	if ans != result {
		t.Errorf("Expect: %v, Actual %v", ans, result)
	}
}

func TestInitDir(t *testing.T) {
	const path1 = "./hello"
	const init1 = PWD
	const path2 = "/hello/bye"
	const init2 = "/hello"
}
