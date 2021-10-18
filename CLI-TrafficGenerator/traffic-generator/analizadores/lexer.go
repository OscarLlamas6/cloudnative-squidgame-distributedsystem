package analizadores

import (
	"fmt"
	"strings"
	"unicode"
)

var (
	tokens                           = []*Token{}
	repetir, anular, ErrorLex bool   = false, false, false
	estado                    int    = 0
	lexemaact, lexemaerror    string = "", ""
)

// Lexico -> funcion para analizar cadena de entrada
// como parametro recibe una cadena ya sea ingresa en consola
// o el texto plano contenido en archivo leido
func Lexico(entrada string) *SquidGameSet {
	tokens = nil
	ErrorLex = false
	lexemaerror = ""
	entrada += " \n"
	split := strings.Split(entrada, "")
	for _, c := range split {
		anular = false
		for ok := true; ok; ok = repetir {
			repetir = false
			switch estado {
			case 0: //posibles transiciones para el Estado 0 (Aceptacion de TK_ESPACIO)

				if isWhiteSpace(c) {
					continue
				} else if c == "\"" { //inicio de cadenas
					lexemaact = ""
					estado = 1
				} else if isInt(c) { //numeros enteros
					estado = 2
					lexemaact = c
				} else if c == "-" { //inicio de bandera
					estado = 3
					lexemaact = c
				} else if isLetter(c) { //inicio de palabra reservada rungame
					estado = 4
					lexemaact = c
				} else {
					ErrorLex = true
					lexemaerror = lexemaact
					lexemaact = ""
				}

			case 1: //posibles transiciones para el estado 1 (Cadenas)

				if c != "\"" {
					estado = 1
					lexemaact += c
				} else if c == "\"" {
					estado = 0
					newToken := NewToken("TK_CADENA", lexemaact)
					tokens = append(tokens, newToken)
					lexemaact = ""
				}

			case 2: //posibles transiciones para el estado 2 (numeros enteros)
				if isInt(c) {
					estado = 2
					lexemaact += c
				} else {
					newToken := NewToken("TK_NUMERO", lexemaact)
					tokens = append(tokens, newToken)
					lexemaact = ""
					estado = 0
					repetir = true
				}

			case 3: //posibles transiciones para el estado 3 (banderas)

				if c == "-" {
					estado = 4
					lexemaact += c
				} else {
					ErrorLex = true
					lexemaerror = lexemaact
					lexemaact = ""
					estado = 0
					repetir = true
				}

			case 4: //posibles transiciones para el estado 5 (palabras reservadas)

				if isLetter(c) {
					estado = 4
					lexemaact += c
				} else {

					switch strings.ToLower(lexemaact) {

					case "--gamename":
						newToken := NewToken("TK_GAMENAME", lexemaact)
						tokens = append(tokens, newToken)
					case "--players":
						newToken := NewToken("TK_PLAYERS", lexemaact)
						tokens = append(tokens, newToken)
					case "--rungames":
						newToken := NewToken("TK_RUNGAMES", lexemaact)
						tokens = append(tokens, newToken)
					case "--concurrence":
						newToken := NewToken("TK_CONCURRENCE", lexemaact)
						tokens = append(tokens, newToken)
					case "--timeout":
						newToken := NewToken("TK_TIMEOUT", lexemaact)
						tokens = append(tokens, newToken)
					case "s", "m", "h", "seg", "segs", "min", "mins", "hr", "hrs", "segundos", "minutos", "horas", "seconds", "minutes", "hours":
						newToken := NewToken("TK_TIEMPO", lexemaact)
						tokens = append(tokens, newToken)
					case "rungame":
						newToken := NewToken("TK_RUNGAME", lexemaact)
						tokens = append(tokens, newToken)
					default:
						ErrorLex = true
						lexemaerror = lexemaact
						lexemaact = ""
						estado = 0
						repetir = true

					}
					estado = 0
					if !anular {
						repetir = true
					} else if anular {
						repetir = false
						anular = false
					}
					lexemaact = ""
				}

			default:
				fmt.Println("")
			}
		}
	}
	if !ErrorLex {
		//fmt.Print(string(getColor("white")), "Cantidad de tokens: ")
		//fmt.Println(string(getColor("yellow")), len(tokens))
		return Sintactico()
		//fmt.Println(string(getColor("blue")), "Analisis exitoso! :D")
	} else {
		fmt.Print(string(getColor("red")), "Error léxico encontrado. ")
		fmt.Println(string(getColor("yellow")), lexemaerror)
		return &SquidGameSet{
			Gamename:    "",
			Players:     0,
			Rungames:    0,
			Concurrence: 0,
			Timeout:     0,
		}
	}

}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isWhiteSpace(s string) bool {
	switch s {
	case " ", "\t", "\n", "\f", "\r":
		return true
	}
	return false
}

// Get Color Name
func getColor(colorName string) string {

	colors := map[string]string{
		"reset":  "\033[0m",
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"blue":   "\033[34m",
		"purple": "\033[35m",
		"cyan":   "\033[36m",
		"white":  "\033[37m",
	}
	return colors[colorName]
}
