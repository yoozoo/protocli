package protocli

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/yoozoo/protocli/util"
)

var protocliIncPath string
var includeFiles = make(map[string]string)

const embeddedDir = "/proto/"

// RegisterIncludeFile add file into includeFiles map
func RegisterIncludeFile(name string, content string) {
	includeFiles[name] = content
}

// GetIncludePath returns the actual include path
func GetIncludePath(userPath string, filePath string) string {
	// extract common includes if needed
	protocliIncPath := util.GetProtocliHome() + util.ProtocliCommonInclude
	if !CheckIncludesExist() {
		path, err := ioutil.TempDir("", "protoapi_inc_")
		if err != nil {
			panic(err)
		}

		err = ExtractIncludes(path)
		if err != nil {
			panic(err)
		}
		protocliIncPath = path
	}

	// return path concatenated with user include path
	var result = protocliIncPath + string(os.PathListSeparator) + filePath
	if len(userPath) > 0 {
		result += string(os.PathListSeparator) + userPath
	}
	return result
}

// CheckIncludesExist check if include files all exists
func CheckIncludesExist() bool {
	for n := range includeFiles {
		if _, err := os.Stat(protocliIncPath + n); err != nil {
			return false
		}
	}
	return true
}

// CleanIncludePath cleanup include path if needed
func CleanIncludePath() {
	// make sure we remove the dir under os.TempDir only
	if len(protocliIncPath) > 0 && strings.HasPrefix(protocliIncPath, os.TempDir()) {
		os.RemoveAll(protocliIncPath)
	}
}

// ExtractIncludes extract common protoconf includes
func ExtractIncludes(path string) error {

	needClean := true
	defer func() {
		if needClean {
			os.RemoveAll(path)
		}
	}()

	for n, c := range includeFiles {
		err := ioutil.WriteFile(path+string(os.PathSeparator)+n, []byte(c), 0666)
		if err != nil {
			return err
		}
	}
	needClean = false
	return nil
}
