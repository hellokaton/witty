package main

import (
	"bufio"
	"os"
	"strings"
	"log"
	"github.com/biezhi/agon/color"
)

func ReadInput() {
	inputReader := bufio.NewReader(os.Stdin)
	color.Print(color.Yellow, "You: ")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		if strings.EqualFold("quit", strings.TrimSpace(input)) ||
			strings.EqualFold("exit", strings.TrimSpace(input)) {
			os.Exit(3)
		} else {
			result := tlAI(input)
			color.Print(color.Green, "Robot: ")
			color.Println(color.Purple, result)
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
