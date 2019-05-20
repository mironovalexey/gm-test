package line

import (
	"bytes"
	"github.com/yuin/goldmark"
)

func Run() {
	source := []byte("Single `line`")
	var buf bytes.Buffer
	if err := goldmark.Convert(source, &buf); err != nil {
		panic(err)
	}
	println(buf.String())
}
