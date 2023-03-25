package functory_method

import "strings"

// Load 加载配置文件
func Load(configFilePath string) {

}

func NewParserFactory(fileExt, category string) ParserFactory {
	switch fileExt {
	case "json":
		switch category {
		case "car":
			return jsonCarParserFactory{}
		}
		/*
			case "yaml":
			case "toml":
			case "ini":
			case "conf":
		*/
	}
	return nil
}

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

// 这里就不一样，这里直接创建一个解析器
// 没有确定是什么解析器
type ParserFactory interface {
	CreateParser() Parser
}

// jsonCarParserFactory
type jsonCarParserFactory struct{}

func (j jsonCarParserFactory) CreateParser() Parser {
	// 汽车解析器的一些配置
	return jsonCarParser{}
}

// jsonComputerParserFactory
// jsonLampParserFactory
// yamlCarParserFactory
// yamlComputerParserFactory
// yamlLampParserFactory
// tomlCarParserFactory
// tomlComputerParserFactory
// tomlLampParserFactory
// iniCarParserFactory
// iniComputerParserFactory
// iniLampParserFactory
// confCarParserFactory
// confComputerParserFactory
// confLampParserFactory

// -----------------------------------------------------

type Parser interface {
	Parse(data []byte)
}

// jsonCarParser
type jsonCarParser struct{}

func (j jsonCarParser) Parse(data []byte) {
	// 解析过程
}

// jsonComputerParser
// jsonLampParser
// yamlCarParser
// yamlComputerParser
// yamlLampParser
// tomlCarParser
// tomlComputerParser
// tomlLampParser
// iniCarParser
// iniComputerParser
// iniLampParser
// confCarParser
// confComputerParser
// confLampParser
