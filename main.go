package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"io/ioutil"
	"log"
)

func ReadInput() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("You: ")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		if strings.EqualFold("quit", strings.TrimSpace(input)) ||
			strings.EqualFold("exit", strings.TrimSpace(input)) {
			os.Exit(3)
		} else {
			result := tlAI(input)
			fmt.Printf("Robot: %s\n", result)
		}
	}
}

// start robot
func Start() {
	for {
		ReadInput()
	}
}

func main() {
	// init config
	if err := InitConf(); err != nil {
		log.Fatal(err)
	} else {
		Start()
	}
}
