package prototype

import (
	"reflect"
	"testing"
	"time"
)

func TestKeywordsClone(t *testing.T) {
	// 旧版本数据
	old_time, _ := time.Parse("2006", "2020")
	old_version := old_time.Unix()
	old_words := KeywordsMap{
		"keyA": &Value{
			Word:      "keyA",
			Visit:     1,
			UpdatedAt: &old_version,
		},
		"keyB": &Value{
			Word:      "keyB",
			Visit:     2,
			UpdatedAt: &old_version,
		},
		"keyC": &Value{
			Word:      "keyC",
			Visit:     3,
			UpdatedAt: &old_version,
		},
	}

	// 新版本数据
	new_version := time.Now().Unix()
	updated_words := []*Value{
		{
			Word:      "keyB",
			Visit:     10,
			UpdatedAt: &new_version,
		},
	}

	// updatedWords 是从数据库中获取到的所有已更新数据
	new_words := old_words.Clone(updated_words)

	if !reflect.DeepEqual(old_words["keyA"], new_words["keyA"]) {
		t.Error("keyA not equal")
	}
	if reflect.DeepEqual(old_words["keyB"], new_words["keyB"]) {
		t.Error("keyB not equal")
	}
	if !reflect.DeepEqual(updated_words[0], new_words["keyB"]) {
		t.Error("updatedWords and keyB not equal")
	}
	if !reflect.DeepEqual(old_words["keyC"], new_words["keyC"]) {
		t.Error("keyC not equal")
	}
}
