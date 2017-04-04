package main

import (
	"bufio"
	"os"
	"strings"
	"log"
	"github.com/biezhi/witty/utils"
	"fmt"
)

func ReadInput() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print(utils.Colorize("You: ", "note"))
	input, err := inputReader.ReadString('\n')
	if err == nil {
		if strings.EqualFold("quit", strings.TrimSpace(input)) ||
			strings.EqualFold("exit", strings.TrimSpace(input)) {
			os.Exit(3)
		} else {
			result := tlAI(input)
			fmt.Print(utils.Colorize("Robot: ", "succ"))
			fmt.Print(utils.Colorize(result, "fail"), "\n")
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
