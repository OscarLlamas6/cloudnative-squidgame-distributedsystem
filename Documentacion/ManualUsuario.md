# Manual de usuario

## USAC SQUID GAMES | Cloud-Native Distributed System

Sistema distribuido usando Kubernetes, gRPC, PubSub, Kafka, RabbitMQ, Golang, NoSQL DB, Sockets io, etc.

## Información General
- Curso: Sistemas Operativos 1
- Segundo Semestre 2021
- Lenguajes: Golang, Javascript.
- Grupo 18

&nbsp;
---
## Integrantes

|Carné | Nombre |
|:----:|:----:|
|201602625| Oscar Alfredo Llamas Lemus|
|201709309| Jose Alejandro Santizo Cotto|
|201801628| Sergio Alexander Echigoyen Gómez|

# Contenido 

- [Descripción](#descripción)
- [Funcionalidad](#funcionalidad)
    - [Traffic Generator](#traffic-generator)
    - [Dashboard](#dashboard)

# Demo

![Demo](https://i.ibb.co/7KDR3DN/sopesgif.gif)

---


## Descripción

Se solicita construir un sistema genérico de arquitectura distribuida que muestre estadísticas en tiempo real utilizando Kubernetes y service mesh como Linkerd y otras tecnologías Cloud Native. En la última parte se utilizará una service mesh para dividir el tráfico. Adicionalmente, se añadirá Chaos Mesh para implementar Chaos Engineering. Este proyecto se aplicará a la visualización de los resultados de juegos implementados por los estudiantes.

---

## Diagrama general del sistema

![Diagrama](https://i.ibb.co/g63JWMj/SOPES1-PROYECTO2.png)


# Funcionalidad

## Traffic Generator

### Pantalla principal del generador de tráfico
<p align="center" >
  <img src="https://i.ibb.co/NLrSG7D/image.png" width="854" height="386" />
</p>
<div style="text-align: center">  </div>
&nbsp;


### Comandos para generador de tráfico

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
### Ingresar el comando y especificar el host al cual enviar el tráfico
<p align="center" >
  <img src="https://i.ibb.co/GHFPrkk/image.png" width="839" height="190" />
</p>
<div style="text-align: center">  </div>
&nbsp;

### Presionar ENTER para iniciar generación de tráfico
<p align="center" >
  <img src="https://i.ibb.co/swC8GrR/image.png" width="852" height="261" />
</p>
<div style="text-align: center">  </div>
&nbsp;

### Una vez finalizadas las peticiones, se mostrará un resumen de las peticiones ejecutadas.
<p align="center" >
  <img src="https://i.ibb.co/j6K0zbV/image.png" width="882" height="556" />
</p>
<div style="text-align: center">  </div>
&nbsp;

## Dashboard

- Para ingresar al Dashbord y visualizar los distintos reportes, entrar a squidgames.IP-INGRESS-CONTROLLER.nip.io/
<p align="center" >
  <img src="https://i.ibb.co/CbnNSZf/image.png" width="839" height="1017" />
</p>
<div style="text-align: center">  </div>
&nbsp;