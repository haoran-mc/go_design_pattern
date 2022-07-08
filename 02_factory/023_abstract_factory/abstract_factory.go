package factory

import (
	"fmt"
	"strings"
)

// RuleConfigParser 规则解析器接口
type RuleConfigParser interface {
	ParseRule(data []byte)
}

// jsonRuleConfigParser
type jsonRuleConfigParser struct{}

func (j jsonRuleConfigParser) ParseRule(data []byte) {
	fmt.Println("json rule parsing...", string(data))
}

// --------------------------

// SystemConfigParser 系统解析器接口
type SystemConfigParser interface {
	ParseSystem(data []byte)
}

// jsonSystemConfigParser
type jsonSystemConfigParser struct{}

func (j jsonSystemConfigParser) ParseSystem(data []byte) {
	fmt.Println("json system parsing...", string(data))
}

// --------------------------

// ConfigParserFactory 工厂方法接口
type ConfigParserFactory interface {
	CreateRuleParser() RuleConfigParser
	CreateSystemParser() SystemConfigParser
}

// --------------------------

type jsonConfigParserFactory struct{}

func (j jsonConfigParserFactory) CreateRuleParser() RuleConfigParser {
	return jsonRuleConfigParser{}
}

func (j jsonConfigParserFactory) CreateSystemParser() SystemConfigParser {
	return jsonSystemConfigParser{}
}

// -----------------------------------------------------

func NewConfigParserFactory(fileExt string) ConfigParserFactory {
	switch fileExt {
	case "json":
		return jsonConfigParserFactory{}
	}
	return nil
}

// Load 加载配置文件
func Load(configFilePath string) {
	ext := GetFileExtension(configFilePath)
	parserFactory := NewConfigParserFactory(ext)
	parser := parserFactory.CreateRuleParser()
	configText := []byte("")
	parser.ParseRule(configText)
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
