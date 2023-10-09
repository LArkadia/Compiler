package lexer

import (
	"compiler/lexer/trees"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type tokensGrouped struct {
	TokenType string   `json:"type"`
	Spaced    bool     `json:"spaced"`
	Operators []string `json:"operators"`
	Operation []string `json:"operation"`
}
type rawTokens struct {
	Tokens []tokensGrouped `json:"tokens"`
}

func loadRawTokens(jsonSource string) rawTokens {
	rawTokensFile, err := os.ReadFile(jsonSource)
	var jsonTokens rawTokens
	if err != nil {
		fmt.Println("Error while reading tokens file!!")
	}
	if err := json.Unmarshal(rawTokensFile, &jsonTokens); err != nil {
		fmt.Println("Error while unmarshalling the tokens!!")
	}
	return jsonTokens
}
func unwrapTokens(jsonTokens rawTokens) (*trees.Node, map[rune]bool, map[string]string) {
	tokensSequences, nonSpacedTokens, tokens := new(trees.Node), make(map[rune]bool), make(map[string]string)
	for _, tokenGroup := range jsonTokens.Tokens {
		for i, tokenOperator := range tokenGroup.Operators {
			tokens[tokenOperator] = tokenGroup.TokenType + tokenGroup.Operation[i]
			if !tokenGroup.Spaced {
				nonSpacedTokens[[]rune(tokenOperator)[0]] = true
			}
			tokensSequences.Add([]rune(tokenOperator))
		}
	}
	return tokensSequences, nonSpacedTokens, tokens
}
func LoadLexerDefinition(jsonSourcePath string) (*trees.Node, map[rune]bool, map[string]string) {
	return unwrapTokens(loadRawTokens(jsonSourcePath))
}
func Lexate(code []rune, tokens_sequences *trees.Node, non_spaced_tokens map[rune]bool, tokens_tag map[string]string, numbers *trees.Node, spaces *trees.Node) string {
	var lexed_code strings.Builder
	code = append(code, ' ')
	buffer := make([]rune, 0)
	pointer := 0
	code_len := len(code)
start:
	for pointer < code_len && spaces.IsComplete(code[pointer:pointer+1]) {
		pointer++
	}
	if pointer >= code_len {
		goto end
	}
	switch {
	case tokens_sequences.Contains(code[pointer : pointer+1]):
		buffer = append(buffer, code[pointer])
		pointer++
		goto Op
	case numbers.IsComplete(code[pointer : pointer+1]):
		buffer = append(buffer, code[pointer])
		pointer++
		goto Num
	case code[pointer] == '"':
		pointer++
		goto Str
	default:
		buffer = append(buffer, code[pointer])
		pointer++
		goto Id
	}
Op:
	//Comments handler
	if len(buffer) == 2 {
		switch {
		case buffer[0] == '/' && buffer[1] == '/':
			for pointer < code_len && code[pointer] != '\n' {
				pointer++
			}
			buffer = []rune{}
			goto start
		case buffer[0] == '/' && buffer[1] == '*':
			for pointer < code_len-1 && (code[pointer] != '*' || code[pointer+1] != '/') {
				pointer++
			}
			pointer += 2
			buffer = []rune{}
			goto start
		}
	}
	//Tokens Handler
	switch {
	case tokens_sequences.Contains(append(buffer, code[pointer])):
		buffer = append(buffer, code[pointer])
		pointer++
		goto Op
	case (tokens_sequences.IsComplete(buffer) && (spaces.Contains(code[pointer:pointer+1]) || (non_spaced_tokens[code[pointer]])) || tokens_sequences.IsComplete(buffer) && non_spaced_tokens[buffer[0]]):
		goto add_tkn_OP

	default:
		goto Id
	}
Num:
	switch {
	case numbers.IsComplete(append(buffer, code[pointer])):
		buffer = append(buffer, code[pointer])
		pointer++
		goto Num
	default:
		goto add_tkn_Num
	}
Str:
	switch {
	case pointer < code_len && code[pointer] != '"':
		buffer = append(buffer, code[pointer])
		pointer++
		goto Str
	default:
		pointer++
		goto add_tkn_Str
	}
Id:
	switch {
	case !spaces.Contains(code[pointer:pointer+1]) && !non_spaced_tokens[code[pointer]]:
		buffer = append(buffer, code[pointer])
		pointer++
		goto Id
	default:
		goto add_tkn_Id

	}
	//Add_tokens
add_tkn_OP:
	lexed_code.WriteString(tokens_tag[string(buffer)] + "\n")
	buffer = []rune{}
	goto start
add_tkn_Num:
	lexed_code.WriteString("Number_" + string(buffer) + "\n")
	buffer = []rune{}
	goto start
add_tkn_Str:
	lexed_code.WriteString("String_" + string(buffer) + "\n")
	buffer = []rune{}
	goto start
add_tkn_Id:
	lexed_code.WriteString("ID_" + string(buffer) + "\n")
	buffer = []rune{}
	goto start
end:
	return lexed_code.String()
}
func Lexate_file(file string) string {
	var lexated_code strings.Builder
	source_code_string, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error while reading the source code file: ", err)
	}
	NTree := trees.GetNumbersTree()
	WS := trees.GetWhiteSpaces()
	tokens_sequences, non_spaced_tokens, tokens := LoadLexerDefinition("tokens.json")
	lexated_code.WriteString(Lexate([]rune(string(source_code_string)), tokens_sequences, non_spaced_tokens, tokens, NTree, WS))
	return lexated_code.String()
}
