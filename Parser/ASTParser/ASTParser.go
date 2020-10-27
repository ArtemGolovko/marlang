package ASTParser

import (
	"../../Token"
	"../STParser"
)

const (
	NodeBody = iota
	FLOAT
	INT

	ADD
	SUB
	MUL
	DIV
	MOD
	POW
)

type (
	IToken = Token.Token
	STNode = STParser.Node
)

var __top_level_node *Node

// Node - struct containing all da shit
type Node struct {
	Name     int
	_type    bool
	Body     interface{}
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

func Parse(STree *STNode) *Node {
	// stream = CreateStream(tokens)
	__top_level_node = &Node{Name: NodeBody, _type: true, Parent: nil}

	// for {
	// 	currentToken := stream.Move()
	// 	if currentToken == nil {
	// 		break
	// 	}

	// 	if IsNumder(currentToken) {
	// 		if IsArithmeticOperator(stream.next) {
	// 			node := &Node{_type: true}
	// 			switch stream.next.Type {
	// 			case Token.ADD:
	// 				node.name = ADD
	// 			case Token.SUB:
	// 				node.name = SUB
	// 			case Token.MUL:
	// 				node.name = MUL
	// 			case Token.DIV:
	// 				node.name = DIV
	// 			case Token.MOD:
	// 				node.name = MOD
	// 			case TokenASTParser.POW:
	// 				node.name = POW
	// 			}
	// 			operand := stream.tokens[stream.index+1]
	// 			if operand.Type != currentToken.Type {
	// 				panic("Invalid type")
	// 			}
	// 			firstOperandNode := &Node{_type: false}
	// 			secondOperandNode := &Node{_type: false}
	// 			if currentToken.Type == Token.FLOAT {
	// 				firstOperandNode.name = FLOAT
	// 				secondOperandNode.name = FLOAT
	// 				firstOperandNode.body, _ = strconv.ParseFloat(currentToken.Value, 64)
	// 				secondOperandNode.body, _ = strconv.ParseFloat(operand.Value, 64)
	// 			} else {
	// 				firstOperandNode.name = INT
	// 				secondOperandNode.name = INT
	// 				firstOperandNode.body, _ = strconv.ParseInt(currentToken.Value, 10, 64)
	// 				secondOperandNode.body, _ = strconv.ParseInt(operand.Value, 10, 64)
	// 			}
	// 			node.AddChild(firstOperandNode)
	// 			node.AddChild(secondOperandNode)

	// 			// will change in future
	// 			__top_level_node.AddChild(node)
	// 			stream.index += 2
	// 		}
	// 	}
	// }

	return __top_level_node
}
