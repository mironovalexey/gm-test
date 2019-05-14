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
			//Note,
		),
	)

	// md.Convert(source, &buf)
	// ->
	//
	reader := text.NewReader(source)
	doc := md.Parser().Parse(reader)

	//doc.Dump(source, 10)
	for c := doc.FirstChild(); c != nil; c = c.NextSibling() {
		name := c.Kind().String()
		println(name)
	}

	md.Renderer().Render(&buf, source, doc)

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

var key = parser.NewContextKey()
var keyStack []ast.Node

func (b *noteParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	line, _ := reader.PeekLine()
	pos := pc.BlockOffset()
	if string(line[pos:pos+6]) == "%NOTE%" {
		reader.Advance(6)
		node := ast.NewBlockquote()
		//println("OPEN", node)
		keyStack = append(keyStack, node)
		pc.Set(key, keyStack)
		return node, parser.HasChildren
	}
	return nil, parser.NoChildren
}

func (b *noteParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	pos := pc.BlockOffset()
	line, _ := reader.PeekLine()
	if len(line) < 8 {
		return parser.Continue | parser.HasChildren
	}
	println("'"+string(line[pos:])+"'", pos)
	lastNode := keyStack[len(keyStack)-1]
	if string(line[pos:pos+9]) == "%ENDNOTE%" && node == lastNode {
		reader.Advance(9)
		keyStack = keyStack[:len(keyStack)-1]
		println("CLOSE")
		return parser.Close
	}
	println("---\n")
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
