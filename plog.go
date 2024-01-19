package plog

import (
	"log"
	"os"
)

var Error = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
