package main

import (
	"fmt"
	"os"
	"github.com/BurntSushi/xgb"
)

var Version = "dev"

func getVersion() string{
	return "LuckyWM version: "+Version
}

func getUsage() string{
	return "usage: luckywm [-v]"
}

func handleArgs(args []string) string{
	if len(args) == 2 && args[1] == "-v"{
		return getVersion()
	}
	return getUsage()
}


func main() {
	args := os.Args
	if len(args) > 1 {
		fmt.Println(handleArgs(args))
		os.Exit(0)
	}

}