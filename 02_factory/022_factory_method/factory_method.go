package factory

import (
	"fmt"
	"strings"
)

// Load 加载配置文件
func Load(configFilePath string) {
	ext := GetFileExtension(configFilePath)
	parserFactory := NewConfigParserFactory(ext) // 通过大工厂获得一个小工厂
	parser := parserFactory.CreateParser()       // 通过小工厂生成一个解析器
	configText := []byte("")
	parser.Parse(configText)
}

// NewConfigParserFactory 生产解析器工厂的工厂，参数是配置文件
// 如果文件后缀是 json，那么就会生产一个 json 解析器工厂
// 如果文件后缀是 yaml，那么就会生产一个 yaml 解析器工厂
// ...
func NewConfigParserFactory(fileExt string) ConfigParserFactory {
	switch fileExt {
	case "json":
		return jsonConfigParserFactory{}
	case "yaml":
		return yamlConfigParserFactory{}
	}
	return nil
}

// GetFileExtension 通过文件名获取扩展名
func GetFileExtension(filePath string) string {
	dotPos := strings.LastIndexByte(filePath, '.')
	fileSuffix := filePath[dotPos+1:]
	if fileSuffix == "json" {
		return "json"
	} else if fileSuffix == "yml" {
		return "yml"
	}
	return ""
}

// -----------------------------------------------------

// ConfigParserFactory 工厂方法接口
type ConfigParserFactory interface {
	CreateParser() ConfigParser
}

// jsonConfigParserFactory jsonConfigParser 的工厂类
type jsonConfigParserFactory struct{}

func (J jsonConfigParserFactory) CreateParser() ConfigParser {
	return jsonConfigParser{}
}

// yamlConfigParserFactory yamlConfigParser 的工厂类
type yamlConfigParserFactory struct{}

func (Y yamlConfigParserFactory) CreateParser() ConfigParser {
	return yamlConfigParser{}
}

// -----------------------------------------------------

// ConfigParser 解析器接口
// 它是一个解析器：可能是 json 解析器；也可能是 yaml 解析器...
type ConfigParser interface {
	Parse(data []byte)
}

// json 配置文件解析类
type jsonConfigParser struct{}

// Parse 解析 json 格式的配置文件
func (J jsonConfigParser) Parse(data []byte) {
	fmt.Println("json parsing...", string(data))
}

// yamlConfigParser yaml 配置文件解析类
type yamlConfigParser struct{}

// Parse 解析 yaml 格式的配置文件
func (Y yamlConfigParser) Parse(data []byte) {
	fmt.Println("yaml parsing...", string(data))
}
