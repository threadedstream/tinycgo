package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/threadedstream/tinycgo/internal/gogen"
	parser "github.com/threadedstream/tinycgo/pkg/antlr"
	"os"
	"strings"
)

const (
	prologue = `
package main

func main() {
		
`

	epilogue = `}`
)

func main() {
	path := os.Args[1]
	input, _ := antlr.NewFileStream(path)
	lexer := parser.NewtinycLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewtinycParser(stream)
	p.BuildParseTrees = true
	tree := p.Program()
	fmt.Println(tree.ToStringTree(nil, p))
	visitor := gogen.NewGogenVisitor()
	progText := visitor.Visit(tree).(string)

	// write to a file
	goSrc := strings.Builder{}
	goSrc.WriteString(prologue)
	goSrc.WriteString(progText)
	goSrc.WriteString(epilogue)

	var err error
	var file *os.File
	if file, err = os.OpenFile(os.Args[2], os.O_RDWR, os.ModePerm); err == nil {
		if _, err = file.Write([]byte(goSrc.String())); err != nil {
			panic(err)
		}
		goto good
	}
	panic(err)
good:
	return
}
