package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
	"traffic-generator/helpers"
)

var (
	SquidGameSet *helpers.SquidGameSet
	Host         string
)

//LimpiarPantalla fuction
func LimpiarPantalla() {

	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	fmt.Println(string(getColor("purple")), "---------------------- SQUID GAME DISTRIBUTED SYSTEM ----------------------")
	fmt.Println()
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

func GetGamesConfig(squidgameset *helpers.SquidGameSet) []*helpers.SingleGame {

	Games := []*helpers.SingleGame{}
	split := strings.Split(squidgameset.GetGameName(), "|")
	count := len(split) / 2
	x := 0

	for i := 0; i < count; i++ {

		newGame := helpers.NewSingleGame(
			strings.TrimSpace(split[x]),
			strings.TrimSpace(split[x+1]),
			squidgameset.GetPlayers(),
			squidgameset.GetRungames(),
			squidgameset.GetConcurrence(),
			squidgameset.GetTimeout())

		Games = append(Games, newGame)
		x = x + 2
	}

	//imprimiendo para verificar estructura
	//jsonF, _ := json.Marshal(Games)
	//fmt.Println(string(jsonF))

	return Games
}

func RunGame(games []*helpers.SingleGame) {

}

func main() {
	continuar := true
	LimpiarPantalla()
	for continuar {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(string(getColor("green")), "USAC ")
		fmt.Print(string(getColor("yellow")), ">> ")
		input, _ := reader.ReadString('\n')
		helpers.Ejecutado = false

		if runtime.GOOS == "windows" {
			input = strings.TrimRight(input, "\r\n")
		} else {
			input = strings.TrimRight(input, "\n")
		}

		if strings.ToLower(input) == "exit" {
			continuar = false
		} else if strings.ToLower(input) == "clear" {
			LimpiarPantalla()
		} else {
			//ejecutando analizador
			SquidGameSet = helpers.Lexico(input)
		}

		if !helpers.ErrorLex && !helpers.SyntaxError && helpers.Ejecutado {
			fmt.Println("")
			fmt.Println(string(getColor("cyan")), "Ingrese el host al cual enviar el tráfico")
			fmt.Print(string(getColor("blue")), "HOST ")
			fmt.Print(string(getColor("yellow")), ">> ")
			fmt.Scanln(&Host)
			fmt.Print(string(getColor("cyan")), "Seteando host -> ")
			fmt.Println(string(getColor("red")), Host)
			fmt.Print(string(getColor("cyan")), "Cargando configuración:")

			for i := 0; i < 53; i++ {
				fmt.Print(string(getColor("yellow")), "#")
				time.Sleep(25 * time.Millisecond)
			}

			fmt.Println("")
			fmt.Print(string(getColor("cyan")), "Configuración cargada correctamente.")
			fmt.Println(string(getColor("yellow")), "Presione ENTER para iniciar el juego. :D")
			var wait string
			fmt.Scanln(&wait)
			fmt.Println(string(getColor("red")), "Ejecutando Squid Games... ")
			RunGame(GetGamesConfig(SquidGameSet))
			fmt.Println(string(getColor("green")), "Juego finalizado.")
		}

	}

	fmt.Print(string(getColor("yellow")), "Gracias por jugar ")
	fmt.Print(string(getColor("red")), "USAC SQUID GAME")
	fmt.Print(string(getColor("yellow")), ", hasta la próxima! :D")
}
