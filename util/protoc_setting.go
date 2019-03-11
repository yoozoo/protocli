package util

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	defaultProtocCmd = "protoc"
	//ProtocBin path for protoc binary under protoapi home
	ProtocBin = "bin/protoc"
	//ProtoapiCommonInclude protoapi common proto file directory
	ProtoapiCommonInclude = "include/protoapi/"

	protocInclude = "include"
)

var (
	pluginName     = ""
	protocliDirEnv = "_PATH"
)

// Init protoc cli name
func Init(name string) {
	pluginName = name
	protocliDirEnv = strings.ToUpper(name) + "_PATH"
}

//ClearDir remove all the files/dirs under a directory
func ClearDir(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

//GetProtocliHome return protoconf home dir
func GetProtocliHome() string {
	homedir := os.Getenv(protocliDirEnv)
	if len(homedir) == 0 {
		if usr, err := user.Current(); err == nil {
			homedir = usr.HomeDir + "/." + pluginName
		}
	}
	if len(homedir) == 0 {
		Die(fmt.Errorf("Failed to find " + pluginName + " home dir. Please make sure current user has home directory"))
	}
	return homedir + "/"
}

// GetDefaultProtoc retrieve protoc executable path and protoc Include path
func GetDefaultProtoc(incPath string) (protoc string, newProtocIncPath string) {
	homedir := GetProtocliHome()

	protoc = homedir + ProtocBin
	if runtime.GOOS == "windows" {
		protoc = protoc + ".exe"
	}
	// check existen
	if _, err := os.Stat(protoc); err != nil {
		Die(fmt.Errorf("Failed to find protoc. Please run `protoapi init` command to initialize. \n\nDetail: %s", err.Error()))
	} else {
		newProtocIncPath = homedir + protocInclude
		if _, err := os.Stat(newProtocIncPath); err != nil {
			Die(fmt.Errorf("Failed to find protoc include folder. Please run `protoapi init` command to initialize.\n\nDetail: %s", err.Error()))
		}
		if len(incPath) > 0 {
			newProtocIncPath += string(os.PathListSeparator) + incPath
		}
		return protoc, newProtocIncPath
	}

	return "", ""
}
