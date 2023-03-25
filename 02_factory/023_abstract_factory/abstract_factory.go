package factory

import (
	"fmt"
	"strings"
)

// Load 加载配置文件
func Load(configFilePath string) {
	ext := GetFileExtension(configFilePath)
	parserFactory := NewConfigParserFactory(ext) // 获取解析器工厂
	parser := parserFactory.CreateCarParser()    // 这里我们指定要解析汽车
	configText := []byte("")
	parser.ParseCar(configText)
}

func NewConfigParserFactory(fileExt string) ConfigParserFactory {
	switch fileExt {
	case "json":
		return jsonParserFactory{}
		/*
			case "yaml":
				return yamlParserFactory{}
			case "toml":
				return tomlParserFactory{}
			case "ini":
				return iniParserFactory{}
			case "conf":
				return confParserFactory{}
		*/
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

// --------------------------

// ConfigParserFactory 工厂方法接口
type ConfigParserFactory interface {
	CreateCarParser() CarParser
	CreateComputerParser() ComputerParser
	CreateLampParser() LampParser
}

type jsonParserFactory struct{}

func (j jsonParserFactory) CreateCarParser() CarParser {
	return jsonCarParser{}
}

func (j jsonParserFactory) CreateComputerParser() ComputerParser {
	return jsonComputerParser{}
}

func (j jsonParserFactory) CreateLampParser() LampParser {
	return jsonLampParser{}
}

// type yamlParserFactory struct{}
// func (y yamlParserFactory) CreateCarParser()
// func (y yamlParserFactory) CreateComputerParser()
// func (y yamlParserFactory) CreateLampParser()

type tomlParserFactory struct{}
type iniParserFactory struct{}
type confParserFactory struct{}

// -----------------------------------------------------

// 汽车解析器
type CarParser interface {
	ParseCar(data []byte)
}

type jsonCarParser struct{}

func (j jsonCarParser) ParseCar(data []byte) {
	fmt.Println("json car parsing...", string(data))
}

// yamlCarParser
// tomlCarParser
// iniCarParser

// 电脑解析器
type ComputerParser interface {
	ParseComputer(data []byte)
}

type jsonComputerParser struct{}

func (j jsonComputerParser) ParseComputer(data []byte) {
	fmt.Println("json computer parsing...", string(data))
}

// yamlComputerParser
// tomlComputerParser
// iniComputerParser

// 台灯解析器
type LampParser interface {
	ParseLamp(data []byte)
}

type jsonLampParser struct{}

func (j jsonLampParser) ParseLamp(data []byte) {
	fmt.Println("json lamp parsing...", string(data))
}

// yamlLampParser
// tomlLampParser
// iniLampParser
