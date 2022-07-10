package main

import (
	"github.com/threadedstream/tinycgo/internal/gogen"
	parser "github.com/threadedstream/tinycgo/pkg/antlr"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
)

func main() {
	path := os.Args[1]
	input, _ := antlr.NewFileStream(path)
	lexer := parser.NewtinycLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewtinycParser(stream)
	tree := p.Program()
	visitor := gogen.NewGogenVisitor()
	visitor.Visit(tree)
}