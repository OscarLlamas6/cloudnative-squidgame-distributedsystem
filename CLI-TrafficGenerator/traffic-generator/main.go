package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"traffic-generator/analizadores"
)

var (
	SquidGameSet *analizadores.SquidGameSet
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
	fmt.Println(string(getColor("purple")), "------------ SQUID GAME DISTRIBUTED SYSTEM ------------")
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

func main() {
	continuar := true
	LimpiarPantalla()
	for continuar {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(string(getColor("green")), "USAC ")
		fmt.Print(string(getColor("yellow")), ">> ")
		input, _ := reader.ReadString('\n')

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
			SquidGameSet = analizadores.Lexico(input)
		}

		/* if !analizadores.ErrorLex && !analizadores.SyntaxError {
			fmt.Println(string(getColor("green")), "EL ANALISIS FUE CORRECTO, LISTO PARA INICIAR SQUID GAME! :D")
		} */

	}

	fmt.Print(string(getColor("yellow")), "Gracias por jugar ")
	fmt.Print(string(getColor("red")), "USAC SQUID GAME")
	fmt.Print(string(getColor("yellow")), ", hasta la próxima! :D")
}
