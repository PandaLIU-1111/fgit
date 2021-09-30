package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"reflect"
)



func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		return
	}

	funcs := map[string]interface{} {
		"pushCommit": pushCommit,
		"version": versionFunc,
	}

	if inArray(args[0], funcs) {
		funcArgs := args[1:]
		_, err := call(funcs, args[0], funcArgs...)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		_, _ = runGitCommand(args...)
	}
	return
}

func call(m map[string]interface{}, name string, params ... string) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	fmt.Println(in)
	result = f.Call(in)
	return
}

func getVersion() string {
	return "0.0.1-bate"
}

func versionFunc()  {
	fmt.Printf("fgit version %s\n", getVersion())
	runGitCommand("--version")
}

func pushCommit(comment string, params ... string)  {
	fmt.Println(params)
	var remote = ""
	var branch = ""

	flag.StringVar(&remote, "origin", "", "--remote")
	flag.StringVar(&branch, "branch", "", "--branch")
	flag.Parse()

	fmt.Printf("remote:%s\n", remote)
	fmt.Printf("branch:%s\n", branch)
	runGitCommand("pull")
	runGitCommand("add", ".")
	runGitCommand("commit", "-m", comment)
	runGitCommand("push")
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


func inArray(need interface{}, needArr map[string]interface{}) bool {
	for k,_ := range needArr{
		if need == k{
			return true
		}
	}
	return false
}
