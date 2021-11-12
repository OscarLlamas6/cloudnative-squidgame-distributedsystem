# USAC SQUID GAMES | Cloud-Native Distributed System

Sistema distribuido usando Kubernetes, gRPC, PubSub, Kafka, RabbitMQ, Golang, NoSQL DB, Sockets io, etc.

# Demo

![Demo](https://i.ibb.co/7KDR3DN/sopesgif.gif)

---

# Diagrama general del sistema

![Diagrama](https://i.ibb.co/g63JWMj/SOPES1-PROYECTO2.png)

---

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


## Descripción

Se solicita construir un sistema genérico de arquitectura distribuida que muestre estadísticas en tiempo real utilizando Kubernetes y service mesh como Linkerd y otras tecnologías Cloud Native. En la última parte se utilizará una service mesh para dividir el tráfico. Adicionalmente, se añadirá Chaos Mesh para implementar Chaos Engineering. Este proyecto se aplicará a la visualización de los resultados de juegos implementados por los estudiantes.


## Explicación

Ingress recibe el tráfico y se redirige a una API que escribe en una cola. Antes de eso, recibe el número de jugadores y elige aleatoriamente un juego de
algoritmo, para elegir el ganador del juego actual, luego escribe los datos necesarios en la cola. El Worker de Go Queue leerá estos datos primero, luego
se escribirán en las bases de datos, Redis para los datos en tiempo real en los paneles y MongoDB para los registros de transacciones.


