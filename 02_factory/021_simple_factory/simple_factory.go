package factory

import (
	"fmt"
	"strings"
)

// ConfigParser 解析器接口
// 它是一个解析器：可能是 json 解析器；也可能是 yaml 解析器...
type ConfigParser interface {
	Parse(data []byte)
}

// jsonConfigParser json 配置文件解析类
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

// Load 加载配置文件
// 这个就是通过工厂模式优化的函数
// 放入一个后缀，就可以得到由工厂生成的解析器
func Load(configFilePath string) {
	ext := GetFileExtension(configFilePath)
	parser := NewConfigParser(ext) // 工厂生成解析器
	configText := []byte("")
	parser.Parse(configText)
}

// NewConfigParser 是生产 ConfigParser 的工厂
// 参数是配置文件，如果文件后缀是 json，那么就会生产出一个解析 json
// 的解释器，如果文件后缀是 yaml，那么就会生产一个解析 yaml 的解析器
func NewConfigParser(fileExt string) ConfigParser {
	switch fileExt {
	case "json":
		return jsonConfigParser{}
	case "yaml":
		return yamlConfigParser{}
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
