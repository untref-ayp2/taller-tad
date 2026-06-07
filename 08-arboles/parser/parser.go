package parser

import "github.com/untref-ayp2/data-structures/tree"

// Parse analiza una expresión aritmética y devuelve su árbol binario.
// Soporta números enteros de varios dígitos y los operadores +, -, *, /.
// Respeta la prioridad de operadores y paréntesis.
func Parse(expr string) (*tree.TreeNode[string], error) {
	tokens, err := tokenize(expr)
	if err != nil {
		return nil, err
	}
	p := parser{tokens: tokens}
	result := p.parseExpr()
	if p.pos != len(p.tokens) {
		return nil, ErrExpresionInvalida
	}
	return result, nil
}

type tokenType int

const (
	tokNum tokenType = iota
	tokOp
	tokOpen
	tokClose
)

type token struct {
	typ tokenType
	val string
}

func tokenize(expr string) ([]token, error) {
	var tokens []token
	i := 0
	for i < len(expr) {
		c := expr[i]
		if c == ' ' {
			i++
			continue
		}
		if c >= '0' && c <= '9' {
			start := i
			for i < len(expr) && expr[i] >= '0' && expr[i] <= '9' {
				i++
			}
			tokens = append(tokens, token{tokNum, expr[start:i]})
			continue
		}
		if c == '+' || c == '-' || c == '*' || c == '/' {
			tokens = append(tokens, token{tokOp, string(c)})
			i++
			continue
		}
		if c == '(' {
			tokens = append(tokens, token{tokOpen, "("})
			i++
			continue
		}
		if c == ')' {
			tokens = append(tokens, token{tokClose, ")"})
			i++
			continue
		}
		return nil, ErrCaracterInesperado
	}
	return tokens, nil
}

type parser struct {
	tokens []token
	pos    int
}

func (p *parser) peek() token {
	if p.pos >= len(p.tokens) {
		return token{}
	}
	return p.tokens[p.pos]
}

func (p *parser) consume() token {
	tok := p.tokens[p.pos]
	p.pos++
	return tok
}

// expr → term (('+' | '-') term)*
func (p *parser) parseExpr() *tree.TreeNode[string] {
	left := p.parseTerm()
	for p.peek().typ == tokOp && (p.peek().val == "+" || p.peek().val == "-") {
		op := p.consume().val
		right := p.parseTerm()
		node := tree.NewTreeNode(op)
		node.SetLeft(left)
		node.SetRight(right)
		left = node
	}
	return left
}

// term → factor (('*' | '/') factor)*
func (p *parser) parseTerm() *tree.TreeNode[string] {
	left := p.parseFactor()
	for p.peek().typ == tokOp && (p.peek().val == "*" || p.peek().val == "/") {
		op := p.consume().val
		right := p.parseFactor()
		node := tree.NewTreeNode(op)
		node.SetLeft(left)
		node.SetRight(right)
		left = node
	}
	return left
}

// factor → NUMBER | '(' expr ')'
func (p *parser) parseFactor() *tree.TreeNode[string] {
	if p.peek().typ == tokNum {
		return tree.NewTreeNode(p.consume().val)
	}
	if p.peek().typ == tokOpen {
		p.consume() // '('
		node := p.parseExpr()
		p.consume() // ')'
		return node
	}
	return nil
}
