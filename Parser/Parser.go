package parser

import (
	"../Token"
)

const (
	NodeBody = iota
)

type IToken = Token.Token

var __top_level_node *Node
var stream *TokenStream

type TokenStream struct {
	index  int
	tokens []*IToken
	next   *IToken
}

func CreateStream(tokens []*IToken) *TokenStream {
	stream := &TokenStream{index: 0, tokens: tokens}
	return stream
}

func (this *TokenStream) Move() *IToken {
	if this.index < len(this.tokens) {
		next := this.tokens[this.index]
		this.index = this.index + 1
		this.next = this.tokens[this.index]
		return next
	} else {
		return nil
	}
}

// Node - struct containing all da shit
type Node struct {
	name int
	_type bool
	body []*Node
	symbolName string
	stringBody string
}



func Parse(tokens []*IToken) *Node {
	stream = CreateStream(tokens)
	__top_level_node = &Node{name: NodeBody, _type: true}

	return __top_level_node
}
