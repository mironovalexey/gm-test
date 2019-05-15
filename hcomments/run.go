package hcomments

import (
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"io/ioutil"
)

func Run() {
	source, _ := ioutil.ReadFile("hcomments/test.md")
	var buf bytes.Buffer

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)

	if err := md.Convert(source, &buf); err != nil {
		panic(err)
	}

	println(buf.String())
}
