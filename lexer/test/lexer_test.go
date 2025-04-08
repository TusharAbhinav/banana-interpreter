package test

import (
	"testing"

	lexer "github.com/TusharAbhinav/monkey/lexer"
	token "github.com/TusharAbhinav/monkey/token"
)

// TestExtendedOperators tests the lexer with more complex expressions and operators
func TestExtendedOperators(t *testing.T) {
	input := `let max = 100;
let min = 1;
let x = 5;
let y = 10;

let compare = x < y;
let not_compare = !(x < y);
let complex = (5 + 10 * 2) / 3;
let mixed = min + (max * 2);`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "max"},
		{token.ASSIGN, "="},
		{token.INT, "100"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "min"},
		{token.ASSIGN, "="},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "y"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "compare"},
		{token.ASSIGN, "="},
		{token.IDENT, "x"},
		{token.LT, "<"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "not_compare"},
		{token.ASSIGN, "="},
		{token.BANG, "!"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.LT, "<"},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "complex"},
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.PLUS, "+"},
		{token.INT, "10"},
		{token.ASTERISK, "*"},
		{token.INT, "2"},
		{token.RPAREN, ")"},
		{token.SLASH, "/"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "mixed"},
		{token.ASSIGN, "="},
		{token.IDENT, "min"},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.IDENT, "max"},
		{token.ASTERISK, "*"},
		{token.INT, "2"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runLexerTest(t, input, tests)
}

// TestLargeNumbers tests handling of different sizes of integer literals
func TestLargeNumbers(t *testing.T) {
	input := `5;
10;
5555555;
999999;
0;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "5555555"},
		{token.SEMICOLON, ";"},
		{token.INT, "999999"},
		{token.SEMICOLON, ";"},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runLexerTest(t, input, tests)
}

// TestIdentifiers tests different identifier names
func TestIdentifiers(t *testing.T) {
	input := `let x = 5;
let _y = 10;
let longVariableName = 15;
let camelCaseVar = 20;
let snake_case_var = 25;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "_y"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "longVariableName"},
		{token.ASSIGN, "="},
		{token.INT, "15"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "camelCaseVar"},
		{token.ASSIGN, "="},
		{token.INT, "20"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "snake_case_var"},
		{token.ASSIGN, "="},
		{token.INT, "25"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runLexerTest(t, input, tests)
}

// TestNestedExpressions tests nested expressions with multiple levels of parentheses
func TestNestedExpressions(t *testing.T) {
	input := `let result = (5 + (10 * (2 + 3)));
let complex = ((((1 + 2) + 3) + 4) + 5);`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.INT, "10"},
		{token.ASTERISK, "*"},
		{token.LPAREN, "("},
		{token.INT, "2"},
		{token.PLUS, "+"},
		{token.INT, "3"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "complex"},
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.LPAREN, "("},
		{token.LPAREN, "("},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.PLUS, "+"},
		{token.INT, "2"},
		{token.RPAREN, ")"},
		{token.PLUS, "+"},
		{token.INT, "3"},
		{token.RPAREN, ")"},
		{token.PLUS, "+"},
		{token.INT, "4"},
		{token.RPAREN, ")"},
		{token.PLUS, "+"},
		{token.INT, "5"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runLexerTest(t, input, tests)
}

// TestConditionalStatements tests more complex conditional statements
func TestConditionalStatements(t *testing.T) {
	input := `if (x < 10) {
  if (y > 5) {
    return true;
  } else {
    return x + y;
  }
} else {
  return false;
}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "y"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	runLexerTest(t, input, tests)
}

// TestFunctionDefinitions tests more complex function definitions
func TestFunctionDefinitions(t *testing.T) {
	input := `let multiply = fn(x, y) { return x * y; };
let factorial = fn(n) {
  if (n == 0) {
    return 1;
  } else {
    return n * factorial(n - 1);
  }
};
let apply = fn(x, f) { return f(x); };`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "multiply"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.ASTERISK, "*"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "factorial"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "n"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "n"},
		{token.EQ, "=="},
		{token.INT, "0"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "n"},
		{token.ASTERISK, "*"},
		{token.IDENT, "factorial"},
		{token.LPAREN, "("},
		{token.IDENT, "n"},
		{token.MINUS, "-"},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "apply"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "f"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "f"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runLexerTest(t, input, tests)
}

// TestIllegalCharacters tests handling of illegal characters
func TestIllegalCharacters(t *testing.T) {
	input := `let x = 5;
let @ = 10;
let # = 15;
let y = $;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.ILLEGAL, "@"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.ILLEGAL, "#"},
		{token.ASSIGN, "="},
		{token.INT, "15"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "y"},
		{token.ASSIGN, "="},
		{token.ILLEGAL, "$"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runLexerTest(t, input, tests)
}

// TestKeywordsAsIdentifiers tests how the lexer handles keywords when used in identifier-like contexts
func TestKeywordsAsIdentifiers(t *testing.T) {
	input := `let if_statement = 5;
let function_name = 10;
let return_value = 15;
let true_statement = 20;
let false_flag = 25;
let else_block = 30;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "if_statement"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "function_name"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "return_value"},
		{token.ASSIGN, "="},
		{token.INT, "15"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "true_statement"},
		{token.ASSIGN, "="},
		{token.INT, "20"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "false_flag"},
		{token.ASSIGN, "="},
		{token.INT, "25"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "else_block"},
		{token.ASSIGN, "="},
		{token.INT, "30"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runLexerTest(t, input, tests)
}

// Helper function to run lexer tests
func runLexerTest(t *testing.T, input string, tests []struct {
	expectedType    token.TokenType
	expectedLiteral string
}) {
	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
