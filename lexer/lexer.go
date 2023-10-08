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
Num:
Str:
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
