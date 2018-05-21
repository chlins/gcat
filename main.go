package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func main() {
	// Target file provided by os.Args[1].
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Read file error, %s\n", err.Error())
		os.Exit(1)
	}

	// Analyze code language.
	lexer := lexers.Get(path.Ext(os.Args[1])[1:])
	if lexer == nil {
		if lexers.Analyse(string(file)) != nil {
			lexer = lexers.Analyse(string(file))
		} else {
			lexer = lexers.Fallback
		}
	}

	lexer = chroma.Coalesce(lexer)

	// Formatter
	f := formatters.Get("terminal16m")
	if f == nil {
		f = formatters.Fallback
	}

	// Determine style.
	s := styles.Get("monokai")
	if s == nil {
		s = styles.Fallback
	}

	builder := s.Builder().Add(chroma.Background, "#fff")
	style, err := builder.Build()
	if err != nil {
		fmt.Printf("Internel error, %s\n", err.Error())
	}

	iterator, err := lexer.Tokenise(nil, string(file))
	if err != nil {
		fmt.Printf("Cat file error, %s\n", err.Error())
		os.Exit(1)
	}

	f.Format(os.Stdout, style, iterator)
}
