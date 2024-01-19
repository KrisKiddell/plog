package plog

import (
	"log"
	"os"
)

var Error = log.New(os.Stdout, "\uea87 does this show", log.LstdFlags|log.Lshortfile)
