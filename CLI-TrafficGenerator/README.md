# Traffic Generator: USAC Squid Games - Distributed Cloud Native System

- Generador de peticiones HTTP implementado en Golang


<p align="center" >
  <img src="https://i.ibb.co/wp3b8rq/logo.png" width="750" height="400" />
</p>
&nbsp;

## Ejecutar Traffic Generator

```bash

#Para iniciar cliente
> cd traffic-generator
> go run ./main.go

```

## Comandos

|Comando | Parámetro | Descripción | Ejemplo 
|:----:|:----:|:----:|:----:|
|`rungame`| - | Palabra reservada inicial para definir una instrucción. | `rungame --gamename "1 \| Game1 \| 2 \| Game2" --players 30 `
|`--gamename`| string | Nombre del o los juegos usando el formato GAME_NUMBER\|GAME_NAME. | `--gamename "1 \| Game1 \| 2 \| Game2"` 
|`--rungames`| int | Número de veces que se ejecutará cada juego. | `--rungames 100`
|`--players`| int | Número de jugadores por juego. | `--players 60`
|`--concurrence`| int | Número de peticiones concurrentes generadas por el Traffic Generator | `--concurrence 25`
|`--timeout`| int + time unit | Tiempo en que el generador de tráfico debe parar. | `--timeout 3m o --timeout 180s`

### Ejemplos de comandos válidos

```bash

> rungame --gamename "1 | Game1 | 2 | Game2" --players 66 --rungames 30000 --concurrence 10 --timeout 3m

> rungame --gamename "1| Red Light, Green Light" --players 18 --rungames 100 --concurrence 5 --timeout 180 s

# Si alguno de los parámetros no se especifíca en la instrución, el generador de tráfico usará los siguientes valores:
#players = 20
#rungames = 50
#concurrence = 5
#timeout = 180 (segundos)
#gamename = "1 | Red Light, Green Light"
```