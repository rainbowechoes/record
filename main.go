package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	userOp string
)

// Local connect mode
const localMode = "local"

// Remote connect mode
const remoteMode = "remote"

func main() {
	// choose the connect computer, localhost/remote
	fmt.Println("Welcome to Record, which can be used to generate op scripts")
	fmt.Printf("Choose the connect mode: %s/%s(local mode is default, Enter 'exit' will exit)\n", localMode, remoteMode)
	fmt.Scanln(&userOp)
	var connectType int
	if userOp == localMode || userOp == "" {
		connectType = 1
	} else if userOp == remoteMode {
		connectType = 2
	} else if userOp == "exit" {
		os.Exit(0)
	} else {
		fmt.Println("Choose error mode")
		os.Exit(1)
	}

	if connectType == 1 {
		fmt.Println("You choose local mode, Enter 'exit' will exit.")
		// call local shell
		for {
			output, err := readShellExecLocal()
			if err == nil {
				if output != "" {
					if output == "exit" {
						os.Exit(0)
					} else {
						fmt.Println(output)
					}
				}
			} else {
				log.Printf("exec shell occur error: %v, output: %s\n", err, output)
			}
		}
	}

	if connectType == 2 {
		fmt.Println("You choose remote mode")
		// ssh connect remote
	}

}

// read user shell operation and exec in local
func readShellExecLocal() (string, error) {
	localHostInfo()
	fmt.Scanln(&userOp)
	if userOp == "exit" {
		return "exit", nil
	}
	if userOp != "" {
		output, err := exec.Command(userOp).Output()
		return string(output), err
	}
	return "", nil
}

// read local host info
func localHostInfo() {
	pwd, err := exec.Command("pwd").Output()
	whoami, err := exec.Command("whoami").Output()
	if err != nil {
		log.Fatalln("occur error in get host info.")
		os.Exit(1)
	}
	fmt.Printf("%s@local %s>", strings.Replace(string(whoami), "\n", "", -1), strings.Replace(string(pwd), "\n", "", -1))

}
