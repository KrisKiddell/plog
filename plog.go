package plog

import (
	"log"
	"os"
)

var Error = log.New(os.Stdout, "\uea87", log.LstdFlags|log.Lshortfile)
