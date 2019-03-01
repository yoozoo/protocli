package main

import (
	"github.com/yoozoo/protocli"
)

func main() {
	protocli.PluginName = "foobar"
	protocli.Execute()
}
