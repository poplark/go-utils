package common

import (
	json "encoding/json"
	"github.com/bitly/go-simplejson"
	io "io/ioutil"
)

type JsonConfigStruct struct {
}

func NewJsonConfig() *JsonConfigStruct {
	return &JsonConfigStruct{}
}

func (self *JsonConfigStruct) Load(filename string) (*simplejson.Json, error) {
	rawData, err := io.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data := []byte(rawData)
	return simplejson.NewJson(data)
}

func (self *JsonConfigStruct) Load2(filename string, v interface{}) {
	rawData, err := io.ReadFile(filename)
	if err != nil {
		return
	}

	data := []byte(rawData)

	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
