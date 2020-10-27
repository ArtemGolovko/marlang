package STParser

import (
	"../../Token"
)

const (
	NodeMain = iota
	OperatorStatment
)

type (
	IToken = Token.Token
)

type TokenStream struct {
	index   int
	tokens  []*IToken
	next    *IToken
	current *IToken
}

func CreateStream(tokens []*IToken) *TokenStream {
	stream := &TokenStream{index: 0, tokens: tokens}
	return stream
}

func (this *TokenStream) Move() *IToken {
	if this.index < len(this.tokens) {
		this.current = this.tokens[this.index]
		this.index = this.index + 1
		if len(this.tokens)-1 > this.index {
			this.next = this.tokens[this.index]
		} else {
			this.next = nil
		}
		return this.current
	} else {
		return nil
	}
}

type Node struct {
	Name     int
	_type    bool
	Body     []*IToken
	Children []*Node
	Parent   *Node
}

func (this *Node) AddChild(node *Node) {
	node.Parent = this
	this.Children = append(this.Children, node)
}

func (this *Node) RemoveChild(index int) *Node {
	node := this.Children[index]
	for i := index; i < len(this.Children)-1; i++ {
		this.Children[i] = this.Children[i+1]
	}
	this.Children = this.Children[:len(this.Children)-1]
	return node
}

func (this *Node) GetChildren() []*Node {
	return this.Children
}

func (this *Node) GetChild(index int) *Node {
	if len(this.Children)-1 < index {
		return nil
	}
	return this.Children[index]
}

func (this *Node) GetParent() *Node {
	return this.Parent
}

func Parse(tokens []*IToken) *Node {
	__top_level_node := &Node{Name: NodeMain, _type: true}
	stream := CreateStream(tokens)

	for stream.Move() != nil {

	}

	return __top_level_node
}

func IsNumder(token *IToken) bool {
	return token.Type == Token.INT || token.Type == Token.FLOAT
}

func IsArithmeticOperator(token *IToken) bool {
	return token.Type == Token.ADD ||
		token.Type == Token.SUB ||
		token.Type == Token.MUL ||
		token.Type == Token.DIV ||
		token.Type == Token.MOD ||
		token.Type == Token.POW
}

func IsConpareOperator(token *IToken) bool {
	return token.Type == Token.EQU ||
		token.Type == Token.LAR ||
		token.Type == Token.LES ||
		token.Type == Token.LARe ||
		token.Type == Token.LESe ||
		token.Type == Token.NOTe
}

func IsLogicOperator(token *IToken) bool {
	return token.Type == Token.NOT ||
		token.Type == Token.AND ||
		token.Type == Token.OR
}

func IsAssignOperator(token *IToken) bool {
	return token.Type == Token.ASS ||
		token.Type == Token.ADDa ||
		token.Type == Token.SUBa ||
		token.Type == Token.MULa ||
		token.Type == Token.DIVa ||
		token.Type == Token.MODa ||
		token.Type == Token.POWa
}

func IsOperator(token *IToken) bool {
	return IsArithmeticOperator(token) ||
		IsConpareOperator(token) ||
		IsLogicOperator(token) ||
		IsAssignOperator(token)
}
