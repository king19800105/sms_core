package core

import (
	"os"
	"path/filepath"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"strings"
)

type FileContent map[interface{}]interface{}

type YMLParser struct {
	filePath    string
	fileContent FileContent
	PointPath   []string
}



// 构造
func NewYmlParser() Parser {
	return &YMLParser{}
}

func (yml *YMLParser) ReadFileByPath(path string) FileContent {
	m := make(FileContent)
	d, err := os.Stat(path)

	if nil != err {
		panic(err)
	}

	if !d.IsDir() {
		m = parseYAMLFile(path, m)
	} else {
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			m = parseYAMLFile(path, m)
			return nil
		})
	}

	yml.fileContent = m

	return m
}

func (yml YMLParser) GetContentByCutPoint(point string) (interface{}, bool) {
	yml.PointPath = strings.Split(point, ".")
	return yml.parser()
}

func (yml *YMLParser) SetFileContent(content FileContent) Parser {
	yml.fileContent = content

	return yml
}

func parseYAMLFile(ymlFile string, m FileContent) FileContent {
	_, err := os.Stat(ymlFile)
	ext := filepath.Ext(ymlFile)

	if nil == err && ".yml" == ext {
		fileContent, err := ioutil.ReadFile(ymlFile)
		err = yaml.Unmarshal([]byte(fileContent), m)

		if nil != err {
			panic(err)
		}
	}

	return m
}

func (yml YMLParser) parser() (interface{}, bool) {
	var result interface{}
	var cnt = 0

	for _, item := range yml.PointPath {
		value, ok := yml.fileContent[item]
		cnt++

		if !ok {
			return nil, false
		}

		result = value
		yml.fileContent, ok = value.(FileContent)

		if !ok && cnt == len(yml.PointPath) {
			return result, true
		}
	}

	return result, true
}
