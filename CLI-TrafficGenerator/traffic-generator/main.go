package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"
	"traffic-generator/helpers"
)

var (
	SquidGameSet          *helpers.SquidGameSet
	Host                  string
	TimeOutStatus         bool = false
	TimeOutValue          int64
	ConcurrenceValue      int64
	RungamesValue         int64
	PlayersValue          int64
	Success, Failed, Send int64
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

func GametimeOut(segundos int64) {
	time.Sleep(time.Duration(segundos) * time.Second)
	TimeOutStatus = true
}

func GetGamesConfig(squidgameset *helpers.SquidGameSet) []*helpers.SingleGame {

	TimeOutValue = squidgameset.GetTimeout()
	RungamesValue = squidgameset.GetRungames()
	ConcurrenceValue = squidgameset.GetConcurrence()
	PlayersValue = squidgameset.GetPlayers()

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

func sendRequest(numero int, game *helpers.SingleGame) {

	// Convert Json Body
	postBody := new(bytes.Buffer)
	json.NewEncoder(postBody).Encode(game)

	// Create Cliente
	client := &http.Client{}

	// Made Request
	req, _ := http.NewRequest("POST", Host+"/", postBody)

	// Add Headers
	req.Header.Add("Content-Type", "application/json")

	// Make Request
	resp, _ := client.Do(req)

	var res map[string]interface{}

	// Decoder Json
	json.NewDecoder(resp.Body).Decode(&res)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {

		fmt.Print(string(getColor("yellow")), "Rungame # ", numero+1)
		fmt.Print(string(getColor("purple")), " Status:")
		fmt.Print(string(getColor("blue")), resp.StatusCode)
		fmt.Print(string(getColor("purple")), " Response:")
		fmt.Println(string(getColor("yellow")), res["Mensaje"])
		Success++
	} else {
		fmt.Print(string(getColor("yellow")), "Rungame #", numero+1)
		fmt.Print(string(getColor("purple")), "Status:")
		fmt.Print(string(getColor("red")), resp.StatusCode)
		fmt.Print(string(getColor("purple")), "Response:")
		fmt.Println(string(getColor("red")), res["Mensaje"])
		Failed++
	}
}

func RunGame(games []*helpers.SingleGame) {

	TimeOutStatus = false
	Success = 0
	Failed = 0
	Send = 0

	//Inicio de Go Routine para controlar el timeout
	go GametimeOut(TimeOutValue)

	for _, game := range games {

		//Canal para ir pasando las peticiones y que se ejecuten concurrentemente
		var ch = make(chan int, ConcurrenceValue+1)
		//Creamos el grupo de concurrencia
		var wg sync.WaitGroup
		//Seteamos la cantidad de peticiones concurrence, dadas por el parametro --concurrence
		wg.Add(int(ConcurrenceValue))
		for i := 0; i < int(ConcurrenceValue); i++ {
			go func() {
				for {
					a, ok := <-ch
					//si el canal esta vacio, terminamos la goroutine
					if !ok || TimeOutStatus {
						wg.Done()
						return
					}
					//Aqui mandamos la peticion
					sendRequest(a, game)
					Send++
				}
			}()
			//Si se cumple el timeout interrumpimos el loop
			if TimeOutStatus {
				break
			}
		}
		//Iteramos sobre el valor del parametro --rungames
		for i := 0; i < int(RungamesValue); i++ {
			//Enviamos el contador al channel
			ch <- i
			if TimeOutStatus {
				break
			}
		}

		close(ch)
		wg.Wait()

		if TimeOutStatus {
			break
		}
	}

	if TimeOutStatus && (Send != RungamesValue) {
		fmt.Println("")
		fmt.Println(string(getColor("red")), "El juego no ha podido finalizar correctamente porque el timeout expiró :(")
	} else {
		fmt.Println("")
		fmt.Println(string(getColor("purple")), "Juego terminado exitosamente! :D")
	}

	fmt.Print(string(getColor("green")), "Peticiones exitosas:")
	fmt.Println(string(getColor("yellow")), Success)
	fmt.Print(string(getColor("green")), "Peticiones fallidas:")
	fmt.Println(string(getColor("yellow")), Failed)
	fmt.Print(string(getColor("green")), "Juegos realizados:")
	fmt.Println(string(getColor("cyan")), Send)
	fmt.Print(string(getColor("green")), "Juegos pendientes (Timeout):")
	fmt.Println(string(getColor("cyan")), RungamesValue-Send)
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
		} else if strings.ToLower(input) == "clear" || strings.ToLower(input) == "cls" {
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
			Host = strings.TrimSuffix(Host, "/")
			Host = strings.TrimSpace(Host)
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
		}

	}

	fmt.Print(string(getColor("yellow")), "Gracias por jugar ")
	fmt.Print(string(getColor("red")), "USAC SQUID GAME")
	fmt.Print(string(getColor("yellow")), ", hasta la próxima! :D")
}
