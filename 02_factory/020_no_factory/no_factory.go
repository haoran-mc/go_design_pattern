package factory

import (
	"errors"
	"fmt"
	"strings"
)

// ConfigParser 配置文件解析器接口
type ConfigParser interface {
	Parse(data []byte)
}

// jsonConfigParser json 格式的配置文件解析器
type jsonConfigParser struct{}

func (J jsonConfigParser) Parse(data []byte) {
	fmt.Println("json parsing...", string(data))
}

// yamlConfigParser yaml 格式的配置文件解析器
type yamlConfigParser struct{}

func (Y yamlConfigParser) Parse(data []byte) {
	fmt.Println("yaml parsing...", string(data))
}

// Load 加载配置文件
func Load(configFilePath string) {
	ext, err := GetFileExtension(configFilePath) // 获取配置文件扩展名
	if err != nil {
		panic("cannot get file extension")
	}

	var parser ConfigParser
	if ext == "json" { // 根据不同格式，获取不同解析器
		parser = jsonConfigParser{}
	} else if ext == "yaml" {
		parser = yamlConfigParser{}
	}

	configText := "name: mc" // 配置文件内容
	parser.Parse([]byte(configText))
}

// GetFileExtension 通过文件名获取扩展名
func GetFileExtension(filePath string) (string, error) {
	dotPos := strings.LastIndexByte(filePath, '.')
	fileSuffix := filePath[dotPos+1:]
	if fileSuffix == "json" {
		return "json", nil
	} else if fileSuffix == "yaml" {
		return "yaml", nil
	}
	return "", errors.New("cannot get file extension")
}
