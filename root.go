package protocli

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/debug"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
	"github.com/yoozoo/protocli/generator"
	"github.com/yoozoo/protocli/util"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd *cobra.Command

// execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Run is the main entraince for protocli
func Run() {
	defer func() {
		util.CleanIncludePath()
		if r := recover(); r != nil {
			log.Printf("%s: %s", r, debug.Stack())
			os.Exit(1)
		}
	}()

	stat, err := os.Stdin.Stat()
	args := os.Args

	// when no any parameter and not reading from char device, treat it as being called by protoc
	if len(args) == 1 && err == nil && (stat.Mode()&os.ModeCharDevice) == 0 {
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			util.Die(fmt.Errorf("reading stdin error: %s", err.Error()))
		}

		response := generator.Generate(input)

		output, err := proto.Marshal(response)
		if err != nil {
			util.Die(fmt.Errorf("create response failed: %s", err.Error()))
		}
		_, err = os.Stdout.Write(output)
		if err != nil {
			util.Die(err)
		}
	} else {
		execute()
	}
}
