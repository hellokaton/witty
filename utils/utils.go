package utils

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

// color output
func Colorize(text string, status string) string {
	out := ""
	switch status {
	case "succ":
		out = "\033[32;1m"    // Blue
	case "fail":
		out = "\033[31;1m"    // Red
	case "warn":
		out = "\033[33;1m"    // Yellow
	case "note":
		out = "\033[34;1m"    // Green
	default:
		out = "\033[0m"    // Default
	}
	return out + text + "\033[0m"
}

// load config
func LoadConfig(filename string) (map[string]string, error) {
	var cmap map[string]string
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return nil, err
	}

	if err := json.Unmarshal(bytes, &cmap); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return nil, err
	}
	return cmap, nil
}
