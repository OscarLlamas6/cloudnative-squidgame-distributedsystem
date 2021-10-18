package helpers

// Token structure
type Token struct {
	Tipo   string
	Lexema string
}

//NewToken returned fuction
func NewToken(t string, l string) *Token {
	return &Token{
		Tipo:   t,
		Lexema: l,
	}
}

//GetTipo fuction
func (t *Token) GetTipo() string {
	return t.Tipo
}

//GetLexema fuction
func (t *Token) GetLexema() string {
	return t.Lexema
}
