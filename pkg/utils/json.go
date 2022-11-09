package utils

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

func Struct2Json(obj interface{}) string {
	str, err := json.Marshal(obj)
	if err != nil {
		logx.Error("[struct2json]can not convert: %v", err)
	}
	return string(str)
}

func Json2Struct(str string, obj interface{}) {
	err := json.Unmarshal([]byte(str), obj)
	if err != nil {
		logx.Error("[json2struct]can not convert: %v", err)
	}
}

func Struct2StructByJson(s1, s2 interface{}) {
	jsonStr := Struct2Json(s1)
	Json2Struct(jsonStr, s2)
}
