package read

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
)

var bindingRootDirectory = GetBindingRootDirectory()

func ReadBindingContent() (map[string]string, error) {
	content := make(map[string]string)
	return content, nil
}

func ReadFile(file string) (string, error) {
	absFilePath := BuildFullFilePath(file)

	// If no file is provided then we don't try to read it
	if absFilePath == "" {
		return "", nil
	}

	// Read the file
	buf, err := ioutil.ReadFile(absFilePath)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func BuildFullFilePath(filename string) string {
	// If the value is in quotes, unquote it
	unquotedFile, err := strconv.Unquote(filename)
	if err != nil {
		// values without quotes will raise an error, ignore it.
		unquotedFile = filename
	}

	// If no file is provided, leave val unchanged.
	if unquotedFile == "" {
		return ""
	}

	// Ensure the absolute file path is used
	absFilePath := unquotedFile
	if !filepath.IsAbs(unquotedFile) {
		absFilePath = filepath.Join(bindingRootDirectory, unquotedFile)
	}
	return absFilePath
}
