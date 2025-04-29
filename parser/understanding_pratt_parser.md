# Understanding Pratt Parsing

## What is Pratt Parsing?

Pratt parsing, also known as Top-Down Operator Precedence parsing, is an elegant technique for parsing expressions developed by Vaughan Pratt in 1973. It's particularly well-suited for handling expressions with operators of varying precedence and associativity while maintaining the simplicity of recursive descent parsing.

## Why Use Pratt Parsing?

- **Intuitive** - The approach closely mirrors how humans mentally parse expressions
- **Flexible** - Easily handles operators with different precedence levels
- **Extensible** - New operators can be added without major restructuring
- **Efficient** - Provides linear time parsing for most expression grammars

## Core Concepts

### 1. Binding Power (Precedence)

At the heart of Pratt parsing is the concept of binding power (BP), which represents operator precedence. Higher numbers indicate stronger binding.

Example precedence levels:
- `*`, `/`: 50 (high precedence)
- `+`, `-`: 40 (lower precedence)
- `==`, `!=`: 30 (even lower precedence)

### 2. Prefix and Infix Parsing Functions

Pratt parsing distinguishes between two types of expressions:

- **Prefix expressions**: Start with an operator followed by operands (e.g., `-5`, `!true`)
- **Infix expressions**: Have an operator between operands (e.g., `a + b`, `x * y`)

Each token type is associated with parsing functions:
- **Prefix parsing functions**: Handle expressions that start with a particular token
- **Infix parsing functions**: Handle expressions where the token appears between expressions

### Basic Implementation in Go

Here's a simplified structure of a Pratt parser in Go:

```go
type Parser struct {
    tokens   []Token
    position int
    prefixParseFns map[TokenType]prefixParseFn
    infixParseFns  map[TokenType]infixParseFn
}

type prefixParseFn func() Expression
type infixParseFn func(Expression) Expression

func (p *Parser) parseExpression(precedence int) Expression {
    // Get the prefix function for the current token
    prefix := p.prefixParseFns[p.currentToken.Type]
    if prefix == nil {
        // Error: no prefix parse function for this token
        return nil
    }
    
    // Parse the left-hand side of the expression
    leftExp := prefix()
    
    // Continue parsing infix expressions while the next token has higher precedence
    for p.peekToken.Type != EOF && precedence < p.peekPrecedence() {
        infix := p.infixParseFns[p.peekToken.Type]
        if infix == nil {
            return leftExp
        }
        
        p.nextToken()
        leftExp = infix(leftExp)
    }
    
    return leftExp
}
```

## Registering Parser Functions

For each token type, you need to register appropriate parsing functions:

```go
func (p *Parser) registerPrefix(tokenType TokenType, fn prefixParseFn) {
    p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType TokenType, fn infixParseFn) {
    p.infixParseFns[tokenType] = fn
}

func NewParser(tokens []Token) *Parser {
    p := &Parser{
        tokens: tokens,
        prefixParseFns: make(map[TokenType]prefixParseFn),
        infixParseFns: make(map[TokenType]infixParseFn),
    }
    
    // Register prefix parse functions
    p.registerPrefix(INT, p.parseIntegerLiteral)
    p.registerPrefix(IDENT, p.parseIdentifier)
    p.registerPrefix(MINUS, p.parsePrefixExpression)
    p.registerPrefix(BANG, p.parsePrefixExpression)
    p.registerPrefix(LPAREN, p.parseGroupedExpression)
    
    // Register infix parse functions
    p.registerInfix(PLUS, p.parseInfixExpression)
    p.registerInfix(MINUS, p.parseInfixExpression)
    p.registerInfix(SLASH, p.parseInfixExpression)
    p.registerInfix(ASTERISK, p.parseInfixExpression)
    p.registerInfix(EQ, p.parseInfixExpression)
    p.registerInfix(NOT_EQ, p.parseInfixExpression)
    
    return p
}
```

## Example Parse Functions

Here are examples of parsing functions for different expression types:

### Integer Literal

```go
func (p *Parser) parseIntegerLiteral() Expression {
    lit := &IntegerLiteral{Token: p.currentToken}
    
    value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
    if err != nil {
        // Error handling
        return nil
    }
    
    lit.Value = value
    return lit
}
```

### Infix Expression

```go
func (p *Parser) parseInfixExpression(left Expression) Expression {
    expression := &InfixExpression{
        Token:    p.currentToken,
        Operator: p.currentToken.Literal,
        Left:     left,
    }
    
    precedence := p.currentPrecedence()
    p.nextToken()
    expression.Right = p.parseExpression(precedence)
    
    return expression
}
```

### Prefix Expression

```go
func (p *Parser) parsePrefixExpression() Expression {
    expression := &PrefixExpression{
        Token:    p.currentToken,
        Operator: p.currentToken.Literal,
    }
    
    p.nextToken()
    expression.Right = p.parseExpression(PREFIX)
    
    return expression
}
```

## Implementing Precedence

Define precedence levels for your operators:

```go
const (
    _ int = iota
    LOWEST
    EQUALS      // ==
    LESSGREATER // > or <
    SUM         // +
    PRODUCT     // *
    PREFIX      // -X or !X
    CALL        // myFunction(X)
)

var precedences = map[TokenType]int{
    EQ:       EQUALS,
    NOT_EQ:   EQUALS,
    LT:       LESSGREATER,
    GT:       LESSGREATER,
    PLUS:     SUM,
    MINUS:    SUM,
    SLASH:    PRODUCT,
    ASTERISK: PRODUCT,
    LPAREN:   CALL,
}

func (p *Parser) peekPrecedence() int {
    if p, ok := precedences[p.peekToken.Type]; ok {
        return p
    }
    return LOWEST
}

func (p *Parser) currentPrecedence() int {
    if p, ok := precedences[p.currentToken.Type]; ok {
        return p
    }
    return LOWEST
}
```

## Parsing Process Example

To understand how Pratt parsing works in practice, let's trace through parsing an expression like `5 + 3 * 2`:

1. Start with `parseExpression(LOWEST)`
2. Parse `5` as an integer literal (prefix parsing function)
3. Check if the next token (`+`) has higher precedence than `LOWEST` - it does
4. Apply the infix parsing function for `+`, making `5` the left side
5. Parse the right side by calling `parseExpression(SUM)` 
6. For the right side, first parse `3` as an integer literal
7. Check if the next token (`*`) has higher precedence than `SUM` - it does
8. Apply the infix parsing function for `*`, making `3` the left side
9. Parse `2` as the right side of `*`
10. No more tokens with higher precedence, so return `3 * 2` as the right side of `+`
11. Final result is a tree representing `5 + (3 * 2)`

This naturally enforces operator precedence without needing separate grammar rules.

## Advantages of Pratt Parsing

1. **Simple Implementation**: Much simpler than traditional parsers for expressions
2. **Performance**: Linear time complexity for most expressions
3. **Readability**: The code structure closely follows the grammar structure
4. **Extensibility**: Adding new operators just requires registering new parsing functions

## Further Reading

- Vaughan Pratt's original paper: "Top Down Operator Precedence" (1973)
- "Writing An Interpreter In Go" by Thorsten Ball
- "Pratt Parsers: Expression Parsing Made Easy" by Bob Nystrom
- The article "Simple Top-Down Parsing in Python" by Fredrik Lundh

## Conclusion

Pratt parsing provides an elegant solution for parsing expressions with different operator precedences. Its balance of simplicity and power makes it an excellent choice for interpreters and compilers, especially for expression-heavy languages.