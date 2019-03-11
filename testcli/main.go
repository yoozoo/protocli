package main

import (
	"github.com/yoozoo/protocli"
)

func main() {
	protocli.Init("foobar", "0.0.2")
	protocli.Run()
}
