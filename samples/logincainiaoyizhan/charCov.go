package main

import (
	"github.com/axgle/mahonia"
)

// ansiToUtf8
func ansiToUtf8(bs []byte) string {
	decoder := mahonia.NewDecoder("gbk")
	return decoder.ConvertString(string(bs))
}

// utf8ToAnsi
func utf8ToAnsi(bs []byte) string {
	decoder := mahonia.NewEncoder("gbk")
	return decoder.ConvertString(string(bs))
}
