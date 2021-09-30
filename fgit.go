package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
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
		fmt.Println(funcArgs)
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

	in := make([]reflect.Value, len(params))
	fmt.Println(in)
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
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
	//var comment = params[0]
	var remote = ""
	var branch = ""
	for i := range params {
		if strings.Contains(params[i], "--remote") {
			pars := strings.Split(params[i], "=")
			remote = pars[1]
		}

		if strings.Contains(params[i], "--branch") {
			pars := strings.Split(params[i], "=")
			branch = pars[1]
		}
	}

	fmt.Println(remote)
	fmt.Println(branch)

	fmt.Printf("remote:%s\n", remote)
	fmt.Printf("branch:%s\n", branch)
	runGitCommand("pull")
	runGitCommand("add", ".")
	runGitCommand("commit", "-m", comment)
	var pushCommand = []string{}
	pushCommand = append(pushCommand, "push")
	fmt.Println(pushCommand)
	//runGitCommand(pushCommand...)
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
