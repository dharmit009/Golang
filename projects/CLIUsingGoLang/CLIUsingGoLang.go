package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	fmt.Println("Infosys Cloud CLI")

	myscanner := bufio.NewScanner(os.Stdin)
	// using forever loop.
	for {
		fmt.Println("------------------")
		fmt.Print("=>")
		myscanner.Scan()
		userInput := myscanner.Text()
		fmt.Println("\n" + userInput)
		var commandString []string = strings.Split(userInput, " ")
		if strings.Compare("exit", commandString[0]) == 0 {
			os.Exit(0)
		} else if strings.Compare("checkip", commandString[0]) == 0 {
			commandOutput := execPINGCommand(commandString[1])
			fmt.Println("Command Output: \n" + commandOutput)
		} else {
			fmt.Println("Command not found!")
		}
	}

}

func execPINGCommand(cmdParam string) string {
	var ip string = cmdParam
	fmt.Println("=================================")
	fmt.Println("Pinging " + ip + " ...")
	fmt.Println("=================================")

	if c, err := exec.Command("ping", "-c 5", ip).CombinedOutput(); err != nil {
		s := "Error! command execution failed"
		return (s)
	} else {
		return string(c)
	}

}
