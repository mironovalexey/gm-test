package hattrs

import (
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"io/ioutil"
)

func Run() {
	source, _ := ioutil.ReadFile("hattrs/test.md")
	var buf bytes.Buffer

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(parser.WithAttribute()),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)

	if err := md.Convert(source, &buf); err != nil {
		panic(err)
	}

	println(buf.String())
}
