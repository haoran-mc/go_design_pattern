package prototype

import (
	"encoding/json"
)

// 搜索框搜索的关键字
type Value struct {
	Word      string // 关键字
	Visit     int    // 关键字被搜索次数
	UpdatedAt *int64 // 更新时间的时间戳
}

// 存储关键字的散列表
type KeywordsMap map[string]*Value

// 原型模式
func (old_words KeywordsMap) Clone(updated_words []*Value) (new_words KeywordsMap) {
	new_words = make(KeywordsMap)

	for k, v := range old_words {
		// 这里是浅拷贝，直接拷贝了地址
		new_words[k] = v
	}

	// 替换掉需要更新的字段，这里用的是深拷贝
	for _, word := range updated_words {
		new_words[word.Word] = word.Clone()
	}
	return new_words
}

// 深拷贝：使用序列化与反序列化的方式
func (v *Value) Clone() *Value {
	var newValue Value
	b, _ := json.Marshal(v)
	json.Unmarshal(b, &newValue)
	return &newValue
}
