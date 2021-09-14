package fileconfig

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var bindingRootDirectory = GetBindingRootDirectory()

const Provider = "provider"
const BindingType = "type"

func ReadServiceBindingConfig() ([]ServiceBinding, error) {
	fs, err := os.Stat(bindingRootDirectory)
	if err != nil {
		return nil, err
	}
	if !fs.IsDir() {
		return nil, errors.New(fmt.Sprintf("Service Binding root %s is not a directory.", bindingRootDirectory))
	}

	lstFile, err := ioutil.ReadDir(bindingRootDirectory)
	if err != nil {
		return nil, err
	}

	// each directory is a set of binding config
	bindings := []ServiceBinding{}
	for _, dir := range lstFile {
		if dir.IsDir() {
			provider, bindingType, content, err := readBindingContent(dir)
			if err != nil {
				return nil, err
			}
			bindings = append(bindings, ServiceBinding{
				Name:        dir.Name(),
				Provider:    provider,
				Properties:  content,
				BindingType: bindingType,
			})
		}
	}
	return bindings, err
}

func readBindingContent(bindigDir os.FileInfo) (string, string, map[string]string, error) {
	subDir := filepath.Join(bindingRootDirectory, bindigDir.Name())
	subDirFiles, err := ioutil.ReadDir(subDir)
	if err != nil {
		return "", "", nil, err
	}
	var provider, bindingType string
	content := make(map[string]string)
	for _, prop := range subDirFiles {
		//skip sub-dir and hidden files
		if prop.IsDir() || prop.Name()[0:1] == "." {
			continue
		}
		buf, err := ioutil.ReadFile(filepath.Join(subDir, prop.Name()))
		if err != nil {
			return "", "", nil, err
		}
		if prop.Name() == BindingType {
			bindingType = string(buf)
		} else if prop.Name() == Provider {
			provider = string(buf)
		} else {
			content[prop.Name()] = string(buf)
		}
	}
	return provider, bindingType, content, nil
}
