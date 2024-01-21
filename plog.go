package plog

import (
	"fmt"
	"log"
	"os"
)

var Error = log.New(os.Stdout, "\u001B[38;5;162m  Error in: \u001B[0m", log.Lshortfile)
var Warn = log.New(os.Stdout, "\u001B[38;5;214m  Warning in: \u001B[0m", log.Lshortfile)
var Info = log.New(os.Stdout, "\u001B[38;5;51m  Info: \u001B[0m", log.Lshortfile)
var Success = log.New(os.Stdout, "\u001B[38;5;77m  Success: \u001B[0m", log.Lshortfile)

func CustomLog(color int, icon string, prefix string) *log.Logger {
	return log.New(os.Stdout, fmt.Sprintf("\u001B[38;5;%dm %s %s: \u001B[0m", color, icon, prefix), log.Lshortfile)
}
