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
		"pushCommit":    pushCommit,
		"version":       versionFunc,
		"cleanCheckout": cleanCheckout,
		"saveCheckout": saveCheckout,
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
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func getVersion() string {
	return "0.0.1-bate"
}

func saveCheckout(branch string)  {
	runGitCommand("stash")
	runGitCommand("checkout", branch)
}

func cleanCheckout(branch string)  {
	runGitCommand("reset", "HEAD", "--hard")
	runGitCommand("checkout", branch)
}

func versionFunc()  {
	fmt.Printf("fgit version %s\n", getVersion())
	runGitCommand("--version")
}

func pushCommit(comment string, params ... string)  {
	//var comment = params[0]
	var pushCommand = []string{}
	pushCommand = append(pushCommand, "push")
	pushCommand = append(pushCommand, "origin")
	for i := range params {
		if strings.Contains(params[i], "--remote") {
			pars := strings.Split(params[i], "=")
			pushCommand[1] = pars[1]
		}

		if strings.Contains(params[i], "--branch") {
			pars := strings.Split(params[i], "=")
			pushCommand = append(pushCommand, pars[1])
		}
	}

	runGitCommand("pull")
	runGitCommand("add", ".")
	runGitCommand("commit", "-m", comment)

	runGitCommand(pushCommand...)
}

// 获取当前路径
func getCurrentPath() string {
	str, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return str
}


// 运行 git 命令
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
