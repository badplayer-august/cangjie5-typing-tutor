package cangjie

import (
	"bytes"
	_ "embed"
	"encoding/gob"
)

type CharInfo struct {
	Char string   `json:"char"`
	Key  []string `json:"key"`
}

type Dataset struct {
	A []CharInfo `json:"A"`
	B []CharInfo `json:"B"`
	C []CharInfo `json:"C"`
	N []CharInfo `json:"N"`
}

var (
	//go:embed data.gob
	f []byte
	KeyToName map[rune]string = map[rune]string{
		'a': "日", 'b': "月",
		'c': "金", 'd': "木", 'e': "水", 'f': "火", 'g': "土",
		'h': "竹", 'i': "戈", 'j': "十", 'k': "大", 'l': "中",
		'm': "一", 'n': "弓", 'o': "人", 'p': "心", 'q': "手",
		'r': "口", 's': "尸", 't': "廿", 'u': "山", 'v': "女",
		'w': "田", 'x': "難", 'y': "卜", 'z': "重",
	}
)

func LoadCangjie() Dataset {
	var (
		data    = Dataset{}
		buf     = bytes.NewBuffer(f)
		decoder = gob.NewDecoder(buf)
	)
	if err := decoder.Decode(&data); err != nil {
		panic(err)
	}
	return data
}
