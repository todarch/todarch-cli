package util

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
	"github.com/todarch/todarch-cli/consts"
	"log"
)

func Log(msg interface{}) {
	if viper.GetBool(consts.VERBOSE) {
		log.Println(msg)
	}
}

func Debug(obj interface{}) {
	if viper.GetBool(consts.DEBUG) {
		log.Println(spew.Sdump(obj))
	}
}
