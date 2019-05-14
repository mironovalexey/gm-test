package eof

import (
	"bytes"
	"github.com/yuin/goldmark"
	"io/ioutil"
)

func Run() {
	source, _ := ioutil.ReadFile("eof/test.md")
	var buf bytes.Buffer
	if err := goldmark.Convert(source, &buf); err != nil {
		panic(err)
	}
	println(buf.String())
}
