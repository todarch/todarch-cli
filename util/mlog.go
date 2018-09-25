package util

import (
	"log"
	"os"
)

func Log(msg interface{}) {
	if isDebugEnabled() {
		log.Println(msg)
	}
}

func isDebugEnabled() bool {
	debug := os.Getenv("TODARCH_DEBUG")
	return debug != ""
}
