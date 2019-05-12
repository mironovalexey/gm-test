package note

import (
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	"io/ioutil"
	"os"
)

func Run() {

	source, _ := ioutil.ReadFile("note/test.md")
	var buf bytes.Buffer

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			Note,
		),
	)

	md.Convert(source, &buf)

	ioutil.WriteFile("note/test.html", buf.Bytes(), os.ModePerm)

}

// ---

type note struct {
}

var Note = &note{}

func (e *note) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(NewNoteParser(), 999),
		),
	)
}

// ---

type noteParser struct {
}

var defaultParser = &noteParser{}

func NewNoteParser() parser.BlockParser {
	return defaultParser
}

// ---

var oKey = parser.NewContextKey()

func (b *noteParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	line, _ := reader.PeekLine()
	pos := pc.BlockOffset()
	if string(line[pos:pos+6]) == "%NOTE%" {
		reader.Advance(6)
		return ast.NewBlockquote(), parser.HasChildren
	}
	return nil, parser.NoChildren
}

func (b *noteParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	line, _ := reader.PeekLine()
	pos := pc.BlockOffset()
	if string(line[pos:pos+9]) == "%ENDNOTE%" {
		reader.Advance(9)
		return parser.Close
	}
	return parser.Continue | parser.HasChildren
}

func (b *noteParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
	// nothing to do
}

func (b *noteParser) CanInterruptParagraph() bool {
	return true
}

func (b *noteParser) CanAcceptIndentedLine() bool {
	return false
}
