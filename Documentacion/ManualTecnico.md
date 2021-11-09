# Manual técnico

## USAC SQUID GAMES | Cloud-Native Distributed System

Sistema distribuido usando Kubernetes, gRPC, PubSub, Kafka, RabbitMQ, Golang, NoSQL DB, Sockets io, etc.

## Información General
- Curso: Sistemas Operativos 1
- Segundo Semestre 2021
- Lenguajes: Golang & JavaScript
- Grupo 18

---
## Integrantes

|Carné | Nombre |
|:----:|:----:|
|201602625| Oscar Alfredo Llamas Lemus|
|201709309| Jose Alejandro Santizo Cotto|
|201801628| Sergio Alexander Echigoyen Gómez|

# Contenido 

- [Demo](#demo) 
- [Descripción general](#descripción-general)
- [Diagrama del sistema](#diagrama-general-del-sistema) 
- [Herramientas utilizadas](#herramientas-utilizadas)    
    - [Bases de datos](#bases-de-datos)
        - [Redis](#redis)
        - [MongoDB](#mongodb)
    - [Lenguajes utilizados](#lenguajes-utilizados)
        - [Golang](#go) 
        - [Javascript](#javascript)      
    - [Google Cloud Platform](#google-cloud-platform)
        - [Google Kubernetes Engine](#google-kubernetes-engine)
        - [Virtual Machines](#virtual-machines)
        - [Pub-Sub](#pub-sub)
    - [Apache Kafka](#apache-kafka)
    - [RabbitMQ](#rabbitmq)
    - [Linkerd](#linkerd)
    - [Chaos Mesh](#chaos-mesh)
    - [Referencias](#referencias)
---

# Demo

![Demo](https://i.ibb.co/7KDR3DN/sopesgif.gif)

---


## Descripción

Se solicita construir un sistema genérico de arquitectura distribuida que muestre estadísticas en tiempo real utilizando Kubernetes y service mesh como Linkerd y otras tecnologías Cloud Native. En la última parte se utilizará una service mesh para dividir el tráfico. Adicionalmente, se añadirá Chaos Mesh para implementar Chaos Engineering. Este proyecto se aplicará a la visualización de los resultados de juegos implementados por los estudiantes.

---

## Diagrama general del sistema

![Diagrama](https://i.ibb.co/g63JWMj/SOPES1-PROYECTO2.png)

---

## Bases de datos

- Para el almacenamiento de los datos se implementaron dos bases de datos en maquinas virtuales de Google cloud Platform.

### Redis

<div style="text-align: justify"> Redis es un almacén de estructura de datos de valores de clave en memoria rápido y de código abierto. Redis incorpora un conjunto de estructuras de datos en memoria versátiles que le permiten crear con facilidad diversas aplicaciones personalizadas. Entre los casos de uso principales de Redis se encuentran el almacenamiento en caché, la administración de sesiones, pub/sub y las clasificaciones. </div><br>

<p align="center" >
  <img src="https://i.ibb.co/cQksGzj/redis.png" width="490" height="175" />
</p><br>

### MongoDB

<div style="text-align: justify">. MongoDB es una base de datos orientada a documentos. Esto quiere decir que en lugar de guardar los datos en registros, guarda los datos en documentos. Estos documentos son almacenados en BSON, que es una representación binaria de JSON.</div><br>

<p align="center" >
  <img src="https://i.ibb.co/mFBQwkk/mongo.png" width="490" height="160" />
</p><br>

---

## Lenguajes utilizados
<div style="text-align: justify"> Para la implementación de las distintas tecnologias utilizadas en el proyecto se hizo uso de varios lenguajes de programación los cuales tuvieron uno o varios propósitos específicos. </div><br>

### Go

<div style="text-align: justify"> Golang "Go" es un lenguaje de programación creado por Rober Griesemer, Rob Pike y Ken Thompson, desarrolladores de Google. Desarrollado con el propósito de adaptarse a las nuevas circunstancias tecnológicas, para sustituir aquellos viejos lenguajes de programación como C, que a pesar de ser de los lenguajes más rápidos, también es uno de los más antiguos y por lo que muchos de los paradigmas para los que fue diseñado, ya están obsoletos. En este proyecto fue utilizado en este proyecto para distintos propósitos los cuales fueron: implementación de un Load Tester, una REST API con conexión a CosmosDB & CloudSQL,una Cloud Function y un publisher para Google Pub/Sub.</div><br>

<p align="center" >
  <img src="https://i.ibb.co/khDvtK7/golang-Programing.jpg" width="325" height="170" />
</p><br>

### Javascript
<div style="text-align: justify"> Tradicionalmente JavaScript ha sido empleado para crear efectos y animaciones en la web. Sin embargo, con el paso del tiempo, ha vivido un gran desarrollo y sus funcionalidades se han ampliado enormemente Actualmente el uso más común de JavaScript es la programación de respuestas a eventos en un sitio web. En este proyecto fue utilizado para desarrollar una aplicación web con React.</div><br>

<p align="center" >
  <img src="https://i.ibb.co/02GS2zX/js.jpg" width="325" height="170" />
</p><br>

---

## Google Cloud Platform
<div style="text-align: justify"> Uno de los propósitos principales de este proyecto fue el de explorar e implementar diferentes servicios ofrecidos por Google, los cuales se detallan a continuación.</div>

---

### Google Kubernetes Engine

<div style="text-align: justify"> Google Kubernetes Engine (GKE) proporciona un entorno administrado para implementar, administrar y escalar sus aplicaciones alojadas en contenedores con la infraestructura de Google. El entorno de Kubernetes Engine consta de varias máquinas (específicamente, instancias de Compute Engine) que se agrupan para formar un clúster de contenedores. </div><br>
<p align="center" >
  <img src="https://i.ibb.co/B2kQvrt/0-f-EAM7mj-Adz-IC4-YTB.png" width="400" height="170" />
</p><br>

### Virtual Machines

<div style="text-align: justify"> Google Compute engine es un servicio de procesamiento seguro y personalizable que te permite crear y ejecutar máquinas virtuales en la infraestructura de Google. </div><br>
<p align="center" >
  <img src="https://i.ibb.co/LPTBmT6/googlevm.jpg" width="325" height="170" />
</p><br>


### Pub-Sub

<div style="text-align: justify"> Google Pub/Sub Pub/Sub permite crear sistemas de consumidores y productores de eventos, llamados publicadores y suscriptores </div><br>
<p align="center" >
  <img src="https://i.ibb.co/0yj1bTw/googlepubsub.jpg" width="325" height="170" />
</p><br>

---

## Apache Kafka

<div style="text-align: justify"> Apache Kafka es una plataforma distribuida de transmisión de datos que permite publicar, almacenar y procesar flujos de registros, así como suscribirse a ellos, de forma inmediata. Está diseñada para administrar los flujos de datos de varias fuentes y distribuirlos a diversos usuarios. </div><br>

![apache kafka](https://i.ibb.co/ZL1sHWs/apachekafka.png)

---

## RabbitMQ

<div style="text-align: justify"> RabbitMQ es un software de encolado de mensajes llamado bróker de mensajería o gestor de colas. Dicho de forma simple, es un software donde se pueden definir colas, las aplicaciones se pueden conectar a dichas colas y transferir/leer mensajes en ellas. </div><br>

![rabbitmq](https://i.ibb.co/wScwgmW/rabbitmq.png)

---

## Linkerd

<div style="text-align: justify"> Linkerd es un proxy de red open source diseñado para ser desplegado como Service Mesh y que está basado en finagle y netty. Su principal cometido es hacer de link, como su nombre indica, entre las diferentes piezas de sistemas distribuidos y es un buen compañero para nuestras arquitecturas de microservicios </div><br>

![rabbitmq](https://i.ibb.co/J7n22h3/linkerd.png)

---

## Chaos Mesh

<div style="text-align: justify"> Chaos Mesh es una plataforma de ingeniería del caos nativa de la nube de código abierto. Ofrece varios tipos de simulación de fallas y tiene una enorme capacidad para organizar escenarios de fallas. Con Chaos Mesh, puede simular convenientemente varias anomalías que podrían ocurrir en la realidad durante los entornos de desarrollo, prueba y producción y encontrar problemas potenciales en el sistema. </div><br>

![rabbitmq](https://i.ibb.co/RvKtMKY/chaos-mesh.jpg)

---

## Referencias

- [Redis](https://aws.amazon.com/es/redis/)
- [MongoDB](https://www.genbeta.com/desarrollo/mongodb-que-es-como-funciona-y-cuando-podemos-usarlo-o-no)
- [Google Kubernetes Engine](https://www.qwiklabs.com/focuses/878?locale=es&parent=catalog#:~:text=Google%20Kubernetes%20Engine%20(GKE)%20proporciona,con%20la%20infraestructura%20de%20Google.&text=En%20este%20lab%2C%20adquirir%C3%A1%20experiencia,implementaci%C3%B3n%20de%20aplicaciones%20con%20GKE.)
- [Apache Kafka](https://www.redhat.com/es/topics/integration/what-is-apache-kafka)
- [RabbitMQ](https://blog.bi-geek.com/rabbitmq-para-principiantes/)
- [Linkerd](https://www.paradigmadigital.com/dev/probando-linkerd-el-pionero-de-los-services-mesh/#:~:text=Linkerd%20es%20un%20proxy%20de,basado%20en%20finagle%20y%20netty.&text=Lo%20que%20s%C3%AD%20podemos%20asegurar,arquitecturas%20distribuidas%20de%20grandes%20compa%C3%B1%C3%ADas.)
- [Chaos Mesh](https://chaos-mesh.org/docs/)