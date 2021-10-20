package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	SyntaxError                             bool = false
	token                                   int  = -1
	tokenAux                                *Token
	gamename                                string = "1 | Red Light, Green Light"
	players, rungames, concurrence, timeout int64  = 20, 50, 5, 180
)

func resetearValores() {
	gamename = "1 | Red Light, Green Light"
	players = 20
	rungames = 50
	concurrence = 5
	timeout = 180
}

func NoBlankStrings(list []string) bool {
	for _, s := range list {
		blank := strings.TrimSpace(s) == ""
		if blank {
			return true
		}
	}
	return false
}

//Sintactico fuction
func Sintactico() *SquidGameSet {
	SyntaxError = false
	tokenAux = nextToken()
	token = -1

	if token < (len(tokens) - 1) {
		tokenAux = nextToken()
		resetearValores()
		inicio()
	}

	if !SyntaxError && token >= (len(tokens)-1) {
		split := strings.Split(gamename, "|")
		if (len(split) >= 2) && (len(split)%2 == 0) && (!NoBlankStrings(split)) {
			Ejecutado = true
		} else {
			fmt.Println(string(getColor("red")), "El parámetro --gamename debe cumplir con el formato GAME_NUMBER|GAME_NAME")
			fmt.Println(string(getColor("yellow")), "Ej: --gamename \"1 | Game1 | 2 | Game2\"")
		}
	} else {
		fmt.Println(string(getColor("cyan")), "Error sintáctico encontrado")
	}

	return &SquidGameSet{
		Gamename:    gamename,
		Players:     players,
		Rungames:    rungames,
		Concurrence: concurrence,
		Timeout:     timeout,
	}
}

func inicio() {
	if tokenAux.GetTipo() == "TK_RUNGAME" {
		tokenAux = nextToken()
		instrucciones()
	} else {
		fmt.Print(string(getColor("red")), "Una instrucción debe iniciar con la palabra reservada ")
		fmt.Print(string(getColor("green")), "rungame")
		fmt.Println(string(getColor("red")), ".")
		fmt.Println(string(getColor("yellow")), "Ej: rungame --gamename \"1 | Game1\" [OPTIONS]")
		SyntaxError = true
	}
}

func instrucciones() {

	if tokenAux.GetTipo() == "TK_GAMENAME" {
		tokenAux = nextToken()
		if tokenCorrecto(tokenAux, "TK_CADENA") {
			//Seteando nombre del juego
			gamename = tokenAux.GetLexema()
			_ = gamename
			otraInstruccion()
		} else {
			fmt.Println(string(getColor("red")), "El parámetro --gamename debe ser una cadena. ")
			fmt.Println(string(getColor("yellow")), "Ej: --gamename \"1 | Game1\"")
			SyntaxError = true
		}
	} else if tokenAux.GetTipo() == "TK_PLAYERS" {
		tokenAux = nextToken()
		if tokenCorrecto(tokenAux, "TK_NUMERO") {
			//Seteando número de jugadores
			players, _ = strconv.ParseInt(tokenAux.GetLexema(), 10, 64)
			_ = players
			otraInstruccion()
		} else {
			fmt.Println(string(getColor("red")), "El parámetro --players debe ser un número entero positivo. ")
			fmt.Println(string(getColor("yellow")), "Ej: --players 60")
			SyntaxError = true
		}
	} else if tokenAux.GetTipo() == "TK_RUNGAMES" {
		tokenAux = nextToken()
		if tokenCorrecto(tokenAux, "TK_NUMERO") {
			//Seteando número de veces a ejecutar cada juego
			rungames, _ = strconv.ParseInt(tokenAux.GetLexema(), 10, 64)
			_ = rungames
			otraInstruccion()
		} else {
			fmt.Println(string(getColor("red")), "El parámetro --rungames debe ser un número entero positivo. ")
			fmt.Println(string(getColor("yellow")), "Ej: --rungames 50")
			SyntaxError = true
		}
	} else if tokenAux.GetTipo() == "TK_CONCURRENCE" {
		tokenAux = nextToken()
		if tokenCorrecto(tokenAux, "TK_NUMERO") {
			//Seteando número de peticiones simultaneas a la API para ejecutar los juegos.
			concurrence, _ = strconv.ParseInt(tokenAux.GetLexema(), 10, 64)
			_ = concurrence
			otraInstruccion()
		} else {
			fmt.Println(string(getColor("red")), "El parámetro --concurrence debe ser un número entero positivo. ")
			fmt.Println(string(getColor("yellow")), "Ej: --concurrence 10")
			SyntaxError = true
		}
	} else if tokenAux.GetTipo() == "TK_TIMEOUT" {
		tokenAux = nextToken()
		if tokenCorrecto(tokenAux, "TK_NUMERO") {
			tokenAux = nextToken()
			numeroTimeout, _ := strconv.ParseInt(tokenAux.GetLexema(), 10, 64)
			if tokenCorrecto(tokenAux, "TK_TIEMPO") {

				//Seteando el timeout para el proceso de enviado de peticiones
				unidadTiempo := tokenAux.GetLexema()
				switch strings.ToLower(unidadTiempo) {
				case "s", "seg", "segs", "segundos", "seconds":
					timeout = numeroTimeout
				case "m", "min", "mins", "minutos", "minutes":
					timeout = numeroTimeout * 60
				case "h", "hr", "hrs", "horas", "hours":
					timeout = numeroTimeout * 3600
				default:
					timeout = numeroTimeout
				}
				_ = timeout
				otraInstruccion()
			} else {
				fmt.Println(string(getColor("red")), "El parámetro --timeout debe llevar una dimensional de tiempo. ")
				fmt.Print(string(getColor("green")), "Usar s, segs, segundos ó seconds para ")
				fmt.Println(string(getColor("blue")), "segundos")
				fmt.Print(string(getColor("green")), "Usar m, min, mins, minutos ó minutes para")
				fmt.Println(string(getColor("blue")), "minutos")
				fmt.Print(string(getColor("green")), "Usar h, hr, hrs, horas ó hours para")
				fmt.Println(string(getColor("blue")), "horas")
				fmt.Println(string(getColor("yellow")), "Ej: --timeout 3 m")
				SyntaxError = true
			}
		} else {
			fmt.Println(string(getColor("red")), "El parámetro --timeout debe ser un número entero positivo. ")
			fmt.Println(string(getColor("yellow")), "Ej: --timeout 3 m")
			SyntaxError = true
		}
	}
}

func tokenCorrecto(taux *Token, tipo string) bool {
	if taux != nil {
		return taux.GetTipo() == tipo
	}
	return false
}

func otraInstruccion() {
	if token < (len(tokens) - 1) {
		tokenAux = nextToken()
		instrucciones()
	}
}

func nextToken() *Token {
	if token < (len(tokens) - 1) {
		token++
		return tokens[token]
	}
	return &Token{
		Tipo:   "TK_NULL",
		Lexema: "",
	}
}
