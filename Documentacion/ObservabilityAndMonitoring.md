# Observability And Monitoring

## USAC SQUID GAMES | Cloud-Native Distributed System

Sistema distribuido usando Kubernetes, gRPC, PubSub, Kafka, RabbitMQ, Golang, NoSQL DB, Sockets io, etc.

# Demo

![Demo](https://i.ibb.co/7KDR3DN/sopesgif.gif)

---

# Pregunta 1

<div style="text-align: justify"><h2><b> ¿Cómo funcionan las métricas de oro, cómo se puede interpretar las 7 pruebas de faulty traffic, usando como base los gráficos y métricas que muestra el tablero de Linkerd Grafana? </b></h2> </div>
&nbsp;

## Prueba 1: Google PubSub Queue con 100% del tráfico

## Resultados

![Golden Metrics](https://i.imgur.com/sdcwD2M.jpg)

<div style="text-align: justify"><h3> En base a los resultados obtenidos, se puede observar en el caso de la métrica "Requests por segundo" que al enviar tráfico al endpoint del ingress, gracias a la concurrencia, se llegaron a ejecutar hasta 9.7 requests por segundo. Por otro lado, para el caso del Succes Rate, al no tener implementado en servicio de Faulty Traffic, se tuvo un 100% de efectividad en cada petición realizada. </h3></div>
&nbsp;

![Latency](https://i.imgur.com/vkVtEYB.jpg)

<div style="text-align: justify"><h3> Por último pero no menos importante, podemos observar, para el caso de la métrica Latency, durante el envío de tráfico la latencia de cada request tuvo picos de hasta 300ms y 80ms. </h3></div>

---

## Prueba 2: Kafka Queue con 100% del tráfico

## Resultados

![Golden Metrics](https://i.imgur.com/ZJicAcR.jpg)

<div style="text-align: justify"><h3> En base a los resultados obtenidos, se puede observar en el caso de la métrica "Requests por segundo" que al enviar tráfico al endpoint del ingress, gracias a la concurrencia, se llegaron a ejecutar hasta 4.70 requests por segundo. Por otro lado, para el caso del Succes Rate, al no tener implementado en servicio de Faulty Traffic, se tuvo un 100% de efectividad en cada petición realizada. </h3></div>
&nbsp;

![Latency](https://i.imgur.com/qJvDMqu.jpg)

<div style="text-align: justify"><h3> Por último, podemos observar, para el caso de la métrica Latency, durante el envío de tráfico la latencia de cada request tuvo picos de hasta 1.987s y 1.333s. </h3></div>

---

## Prueba 3: RabbitMQ Queue con 100% del tráfico

## Resultados

![Golden Metrics](https://i.imgur.com/viQulFC.jpg)

<div style="text-align: justify"><h3> En base a los resultados obtenidos, se puede observar en el caso de la métrica "Requests por segundo" que al enviar tráfico al endpoint del ingress, gracias a la concurrencia, se llegaron a ejecutar hasta 10.20 requests por segundo. Por otro lado, para el caso del Succes Rate, al no tener implementado en servicio de Faulty Traffic, se tuvo un 100% de efectividad en cada petición realizada. </h3></div>
&nbsp;

![Latency](https://i.imgur.com/R4G3JIm.jpg)

<div style="text-align: justify"><h3> Por último, podemos observar, para el caso de la métrica Latency, durante el envío de tráfico la latencia de cada request tuvo picos de 96ms y 22ms. </h3></div>

---

## Prueba 4: Google PubSub Queue con 50% del tráfico y 50% Faulty Traffic

## Resultados

![Succes Rate](https://i.imgur.com/bj1T9Qm.jpg)

<div style="text-align: justify"><h3> Para el caso del Succes Rate, al tener implementado en servicio de Faulty Traffic, la gráfica mostró una caida en la taza de éxito de las peticiones y un Succes Rate el 55.16%</h3></div>
&nbsp;


![Golden Metrics](https://i.imgur.com/Cgoclak.jpg)

<div style="text-align: justify"><h3> En base a los resultados obtenidos, se puede observar en el caso de la métrica "Requests por segundo" que al enviar tráfico al endpoint del ingress, gracias a la concurrencia, se llegaron a ejecutar hasta 5.0 requests por segundo. </h3></div>
&nbsp;

![Latency](https://i.imgur.com/WrQoHz0.jpg)

<div style="text-align: justify"><h3> Por último, podemos observar, para el caso de la métrica Latency, durante el envío de tráfico la latencia de cada request tuvo picos de 281ms y 75ms. </h3></div>

---

## Prueba 5: Kafka Queue con 50% del tráfico y 50% Faulty Traffic

## Resultados

![Succes Rate](https://i.imgur.com/Cn2CVm7.jpg)

<div style="text-align: justify"><h3> Para el caso del Succes Rate, al tener implementado en servicio de Faulty Traffic, la gráfica mostró una caida en la taza de éxito de las peticiones y un Succes Rate el 62.38%</h3></div>
&nbsp;


![Golden Metrics](https://i.imgur.com/N4udLm5.jpg)

<div style="text-align: justify"><h3> En base a los resultados obtenidos, se puede observar en el caso de la métrica "Requests por segundo" que al enviar tráfico al endpoint del ingress, gracias a la concurrencia, se llegaron a ejecutar hasta 5.50 requests por segundo. </h3></div>
&nbsp;

![Latency](https://i.imgur.com/0vJeAkd.jpg)

<div style="text-align: justify"><h3> Por último, podemos observar, para el caso de la métrica Latency, durante el envío de tráfico la latencia de cada request tuvo picos de 1.988s y 1.375s. </h3></div>

---

## Prueba 6: RabbitMQ Queue con 50% del tráfico y 50% Faulty Traffic

## Resultados

![Succes Rate](https://i.imgur.com/Afaa4OY.jpg)

<div style="text-align: justify"><h3> Para el caso del Succes Rate, al tener implementado en servicio de Faulty Traffic, la gráfica mostró una caida en la taza de éxito de las peticiones y un Succes Rate el 51.25%</h3></div>
&nbsp;


![Golden Metrics](https://i.imgur.com/dXbKKJf.jpg)

<div style="text-align: justify"><h3> En base a los resultados obtenidos, se puede observar en el caso de la métrica "Requests por segundo" que al enviar tráfico al endpoint del ingress, gracias a la concurrencia, se llegaron a ejecutar hasta 9.50 requests por segundo. </h3></div>
&nbsp;

![Latency](https://i.imgur.com/ElaFcze.jpg)

<div style="text-align: justify"><h3> Por último, podemos observar, para el caso de la métrica Latency, durante el envío de tráfico la latencia de cada request tuvo picos de 29ms y 12s. </h3></div>

---

## Prueba 7: PubSub 33.33%, Kafka 33.33% y RabbitMQ 33.33% 

## Resultados

![Golden Metrics](https://i.imgur.com/613fo9c.jpg)

<div style="text-align: justify"><h3> En base a los resultados obtenidos, se puede observar en el caso de la métrica "Requests por segundo" que al enviar tráfico al endpoint del ingress, gracias a la concurrencia, se llegaron a ejecutar hasta 4.7 requests por segundo por parte de RabbitMQ. Por otro lado, para el caso del Succes Rate, al no tener implementado en servicio de Faulty Traffic, se tuvo un 100% de efectividad en cada petición realizada. </h3></div>
&nbsp;

![Latency](https://i.ibb.co/rbBq6ry/latency6.jpg)

<div style="text-align: justify"><h3> Por último podemos observar, para el caso de la métrica Latency, durante el envío de tráfico la latencia de cada request tuvo picos de hasta 1.942s por parte de Kafka y 30ms por parte de RabbitMQ. </h3></div>

---

# Pregunta 2

<div style="text-align: justify"><h2><b> Menciona al menos 3 patrones de comportamiento que
hayas descubierto en las pruebas de faulty traffic.</b></h2> </div>
&nbsp;

## - Patrón de comportamiento 1

<div style="text-align: justify"><h3> En el caso de Traffic Splitting donde el 100% del tráfico se redireccionaba al path de RabbitMQ, se logró la mayor taza de Peticiones Por Segundo, logrando hasta 10.20RPS.</h3></div><br>


![Pattern](https://i.ibb.co/Df6PR3x/comportamiento-1.jpg)

## - Patrón de comportamiento 2

<div style="text-align: justify"><h3> A pesar de redirigir el 50% del tráfico como Faulty Traffic, la taza de peticiones por segundo de RabbitMQ se mantuvo alta logrando hasta 9.50RPS</h3></div><br>


![Pattern](https://i.ibb.co/g4VfJZB/comportamiento-2.jpg)

## - Patrón de comportamiento 3

<div style="text-align: justify"><h3> Por otro lado, al agregar los servicios de PubSub y Kafka al Traffic splitting, la taza de peticiones por segundo de RabbitMQ se redujo considerablemente, teniendo un RPS similar al de los otros 2 servicios.</h3></div><br>


![Pattern](https://i.ibb.co/NpWdTK6/comportamiento-3.jpg)

---

# Pregunta 3

<div style="text-align: justify"><h2><b> ¿Qué sistema de mensajería es más rápido? ¿Por qué? </b></h2> </div>
&nbsp;

<div style="text-align: justify"><h3> En una prueba realizada con 333 peticiones, en condiciones iguales para todos los servicios, RabbitMQ tuvo la menor latencia por una amplia diferencia logrando valores de 30ms. Por el contrario Kafka fue el servicio de mensajeria más lento con una latencia de 1.942 segundos.</h3></div><br>


![Pattern](https://i.ibb.co/rbBq6ry/latency6.jpg)

---

# Pregunta 4

<div style="text-align: justify"><h2><b> ¿Cuántos recursos utiliza cada sistema de mensajería? </b></h2> </div>
&nbsp;

# Google Pub/Sub

<div style="text-align: justify"><h3> Al utilizar el servicio de mensajeria de Google PubSub en el cluster, se llegó a usar un 19.70% del CPU y un 33.33% de la memoria. Como podemos observar en la gráfica, el servidor gRPC que actua como Publisher consumió más recursos que el Worker (Suscriber)</h3></div><br>

![Pattern](https://i.ibb.co/h7vLJM0/recursos-Pub-Sub.jpg)

---

# Kafka

<div style="text-align: justify"><h3> Al utilizar el servicio de mensajeria Kafka con Strimzi en el cluster, se llegó a usar un 16.00% del CPU y un 33.33% de la memoria. Como podemos observar en la gráfica, el cluster de Kafka instalado con Strimzi es el servicio que más CPU y RAM consume.</h3></div><br>

![Pattern](https://i.ibb.co/1dSkVBg/recursos-Kafka.jpg)
 
---

# RabbitMQ

<div style="text-align: justify"><h3> Al utilizar el servicio de mensajeria RabbitMQ en el cluster, se llegó a usar un 18.49% del CPU y un 33.33% de la memoria. Como podemos observar en la gráfica, a diferencia de PubSub esta vez el Worker (Consumidor de RabbitMQ) consumió más recursos que el Publisher (Servidor gRPC).</h3></div><br>

![Pattern](https://i.ibb.co/PY5gSz4/recursos-Rabbit.jpg)

---

# Pregunta 5

<div style="text-align: justify"><h2><b> ¿Cuáles son las ventajas y desventajas de cada servicio de
mensajería? </b></h2> </div>
&nbsp;

| Servicio | Ventajas | Desventajas
|:----:|:----:|:----:|
| Google PubSub| Está respaldado por Google, cuentas de servico y escalabidadlidad de recursos fáciles de gestionar. Se puede integrar con otros servicios de Google Cloud Platform | No es una solución tan veloz como otras alternativas. Se está sujeto a regiones disponibles de Google.
| Apache Kafka | Sistema escalable horizontalmente y tolerante a fallos. Referencia en la industria, lo usan las empresas más grandes del mundo. Garantías de entrega de mensajes exactly-once (exactamente una vez). | Kafka no es una tecnología que esté diseñada para manejar mensajes muy grandes (+ 1MB). De los 3 servicios implementados, Kafka fue el que tuvo mayor latencia (casi 2 segundos en algunos picos).
| RabbitMQ | En este proyecto fue el sistema de mensajeria implementado con menor latencia. Permite desacoplamiento, es asincrónico, y realiza recorte de picos. Mejor rendimiento y alta concurrencia. | Al estar desarrollado en Erlang es difícil de entender el código fuente. Su rendimiendo disminuye considerablemente en trabajos pesados.

---

# Pregunta 6

<div style="text-align: justify"><h2><b> ¿Cuál es el mejor sistema de mensajería? </b></h2> </div>

<div style="text-align: justify"><h3> Pese a que se redujo la taza de peticiones por segundo al momento de trabajar todos los servicios de mensajería al mismo tiempo; RabbitMQ fue la solucion con la menor latencia, además de un consumo de recursos e indice de concurrencia aceptables.</h3></div><br>

---

# Pregunta 7 

<div style="text-align: justify"><h2><b> ¿Cuál de las dos bases de datos se desempeña mejor y por
qué? (Redis vs MongoDB)</b></h2> </div>

<div style="text-align: justify"><h3> Si hablamos de desempeño, cada base de datos fue utilizada para propósitos distintos, por lo que cada una se desempeñó en base a la funcionalidad que se le dió. Si hablamos de la Base de Datos más rápida, sin duda fue Redis al ser una base de datos in-memory, MongoDB es una base de datos on-disk. Por otro lado Redis consume más recursos. MongoDB almacena la información en documentos y Redis como clave-valor. En este proyecto, Redis al actuar como memoria cache almacenando metadatos, tuvo un mejor rendimiento, sin embargo se pudo notar que a la hora de buscar escalabilidad y almacenamiento más general, MongoDB puede ser la mejor solución. </h3></div><br>

