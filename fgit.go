package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		runGitCommand()
		return
	}

	var innerComand = []string{"pushCommit"}
		pushCommit(args[1])
	if inArray(args[0], innerComand) {
		
	} else {
		_, _ = runGitCommand(args...)
	}
}

func pushCommit(comment string)  {
	runGitCommand("pull")
	runGitCommand("add", ".")
	runGitCommand("commit", "-m", comment)
	runGitCommand("push")
}

func inArray(need interface{}, needArr []string) bool {
	for _,v := range needArr{
		if need == v{
			return true
		}
	}
	return false
}

// 获取当前路径
func getCurrentPath() string {
	str, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return str
}


func runGitCommand(arg ...string) (string, error) {

	cmd := exec.Command("git", arg...)
	cmd.Dir = getCurrentPath()
	//cmd.Stderr = os.Stderr
	msg, err := cmd.CombinedOutput() // 混合输出stdout+stderr
	cmd.Run()

	fmt.Println(string(msg))
	// 报错时 exit status 1
	return string(msg), err
}