package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// 字符串转数字
func Atoi(i string) int {
	v, err := strconv.Atoi(strings.TrimSpace(i))
	if err != nil {
		panic(err)
	}
	return v
}

// 数字转字符串
func Itoa(i int64) string {
	return strconv.Itoa(int(i))
}

// 两数最大值
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 两数最小值
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 序列化json
func JsonMarshalOb(ob interface{}) string {
	bt, err := json.Marshal(ob)
	if err != nil {
		panic(err)
		return ""
	}
	return string(bt)
}

// 反序列化
func JsonUnMarshalOb(obstr string, ob interface{}) error {
	err := json.Unmarshal([]byte(obstr), ob)
	return err
}

// 取md5 大写
func Md5(s string) string {
	return fmt.Sprintf("%X", md5.Sum([]byte(s)))
}

// 取md5 小写
func Md5Lower(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

