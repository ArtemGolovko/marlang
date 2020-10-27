package Repl

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"

	"../Lexer"
	"../Parser/ASTParser"
	"../Parser/STParser"
)

type ASTNode = ASTParser.Node

const PROMPT = "> "

func Start(stdin io.Reader, stdout io.Writer) {
	scanner := bufio.NewScanner(stdin)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "" {
			return
		}

		Parse(line)
	}
}

func ReadFile(filename string) {
	code, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	Parse(string(code))
	return
}

func Parse(code string) {
	lexer := Lexer.New(code)
	tokens := lexer.Lex()
	STree := STParser.Parse(tokens)
	ASTree := ASTParser.Parse(STree)

	fmt.Println("Tokens:")
	for _, token := range tokens {
		fmt.Println(token)
	}

	fmt.Println("Abstract Syntax Tree:")
	showASTNode(ASTree, "", 0)
}

func showASTNode(node *ASTNode, prefix string, index int) {
	if prefix == "" {
		fmt.Printf("%d: ", index)
		fmt.Println(node)
	} else {
		fmt.Printf("%v%d: ", prefix, index)
		fmt.Println(node)
	}
	for i, n := range node.GetChildren() {
		showASTNode(n, prefix+"   ", i)
	}
}
