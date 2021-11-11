# Chaos Mesh

## USAC SQUID GAMES | Cloud-Native Distributed System

Sistema distribuido usando Kubernetes, gRPC, PubSub, Kafka, RabbitMQ, Golang, NoSQL DB, Sockets io, etc.

## Usando Siege para generar tráfico

```bash

# Instalar Siege
> sudo apt update -y
> sudo apt install siege -y
> siege --version

# Ejemplo de uso de siege
> siege -b <IP> -t=1M
```

---

# Pregunta 1

<div style="text-align: justify"><h2><b>¿Cómo se reflejan en los dashboards de Linkerd los experimentos de Chaos Mesh?</b></h2></div><br>

## Experimento 1: Pod Kill
<br>
- Definición del experimento:

```yaml
# Pod Kill
kind: Schedule
apiVersion: chaos-mesh.org/v1alpha1
metadata:
  name: experiment1
  namespace: chaos-testing
spec:
  schedule: '@every 10s'
  concurrencyPolicy: Forbid
  historyLimit: 1
  type: PodChaos
  podChaos:
    selector:
      namespaces:
        - squidgames
      labelSelectors:
        app: grpc-pubsub-client
    mode: one
    action: pod-kill
    duration: 1m
    gracePeriod: 0
```
<br>
- Dashboard Linkerd  

<br>

![linkerd](https://i.ibb.co/whFF0YM/linkerd1.jpg)

<div style="text-align: justify"><h3> En base a la gráfica mostrada en el dashboard de Linkerd, los resultados del experimento PodKill refleja unos picos y cortes en la gráfica que representan el momento en que los Pods se "mueren". </h3></div>

---

## Experimento 2: Pod Failure
<br>
- Definición del experimento:

```yaml
# Pod Failure
kind: Schedule
apiVersion: chaos-mesh.org/v1alpha1
metadata:
  name: experiment1
  namespace: chaos-testing
spec:
  schedule: '@every 10s'
  concurrencyPolicy: Forbid
  historyLimit: 1
  type: PodChaos
  podChaos:
    selector:
      namespaces:
        - squidgames
      labelSelectors:
        app: grpc-pubsub-client
    mode: one
    action: pod-failure
    duration: 1m
    gracePeriod: 0
```
<br>
- Dashboard Linkerd  

<br>

![linkerd](https://i.ibb.co/hBfnSSk/linkerd2.jpg)

<div style="text-align: justify"><h3> En base a la gráfica mostrada en el dashboard de Linkerd, los resultados del experimento PodFailure refleja unos picos y cortes en la gráfica que representan el momento en que los Pods fallan. </h3></div>

---

## Experimento 3: Container Kill
<br>
- Definición del experimento:

```yaml
# Container Kill
kind: Schedule
apiVersion: chaos-mesh.org/v1alpha1
metadata:
  name: experiment1
  namespace: chaos-testing
spec:
  schedule: '@every 10s'
  concurrencyPolicy: Forbid
  historyLimit: 1
  type: PodChaos
  podChaos:
    selector:
      namespaces:
        - squidgames
      labelSelectors:
        app: grpc-pubsub-client
    mode: one
    action: container-kill
    containerNames: ['grpc-pubsub-client']
    duration: 1m
    gracePeriod: 0
```
<br>
- Dashboard Linkerd  

<br>

![linkerd](https://i.ibb.co/1KWWVt7/linkerd3.jpg)

<div style="text-align: justify"><h3> En base a la gráfica mostrada en el dashboard de Linkerd, los resultados del experimento ContainerKill refleja unos picos y cortes en la gráfica que representan el momento en que los contenedores "mueren". </h3></div>

---

## Experimento 4: Network Emulation (Netem) Chaos
<br>
- Definición del experimento:

```yaml
# Slow Network Emulation (Netem)
kind: Schedule
apiVersion: chaos-mesh.org/v1alpha1
metadata:
  name: squidgames-networks-delay
  namespace: chaos-testing
spec:
  schedule: '@every 10s'
  concurrencyPolicy: Forbid
  historyLimit: 1
  type: NetworkChaos
  networkChaos:
    selector:
      namespaces:
        - squidgames
      labelSelectors:
        app: grpc-pubsub-client
    mode: one
    action: delay
    delay:
      latency: '10ms'
    duration: "30s"
```
<br>
- Dashboard Linkerd  

<br>

![linkerd](https://i.ibb.co/xJhdmVy/linkerd4.jpg)

<div style="text-align: justify"><h3> En base a la gráfica mostrada en el dashboard de Linkerd, los resultados del experimento SlowNetwork refleja unos picos y cortes en la gráfica que representan las subidas en la latencia. </h3></div>

---

## Experimento 5: DNS Chaos
<br>
- Definición del experimento:

```yaml
# DNS Chaos
apiVersion: chaos-mesh.org/v1alpha1
kind: DNSChaos
metadata:
  name: dns-chaos-example
  namespace: chaos-testing
spec:
  action: error
  mode: all
  patterns:
    - grpc-rabbit-server-service
    - grpc-pubsub-server-service
    - grpc-kafka-server-service
  selector:
    namespaces:
      - squidgames"
```
<br>
- Dashboard Linkerd  

<br>

![linkerd](https://i.ibb.co/jGYcC3K/linkerd5.jpg)

<div style="text-align: justify"><h3> En base a la gráfica mostrada en el dashboard de Linkerd, los resultados del experimento DNSChaos refleja unos picos y cortes en la gráfica que representan el momento en que los DNS de los servicios de los clientes gRPC devuelven un error. </h3></div>

---
<br>

# Pregunta 2

<div style="text-align: justify"><h2><b>¿En qué se diferencia cada uno de los experimentos realizados? </b></h2></div><br>

| Experimento | Descripción
|:----:|:----:|
| Pod Kill | El "mata" el Pod seleccionado, sin embargo gracias al ReplicaSet del deployment este se vuelve a reiniciar para asegurarel estado deseado.
| Pod Failure | El Pod seleccionado no estará disponible en un período de tiempo específico.
| Container Kill | El contenedor seleccionado se "mata" en el Pod seleccionado.
| Slow Network Emulation | Emula las malas condiciones de la red, como grandes retrasos, alta tasa de pérdida de paquetes, reordenación de paquetes, etc.
| DNS Chaos | El pod seleccionado se inyectará con errores de dns, como error, random.

---
<br>

# Pregunta 3

<div style="text-align: justify"><h2><b>¿Cuál de todos los experimentos es el más dañino?</b></h2></div><br>

<div style="text-align: justify"><h3> De los 5 experimentos realizados, PodFailure y DNSChaos parecieron ser los más dañinos, a diferencia de Pod Kill el cual reiniciaba el Pod instantáneamente, Pod Failure e DNS Chaos inhabilitan los servicios en su totalidad, uno arrojando error por falla de Pod y el otro por erro de DNS. Por otra parte Slow Network a pesar de afectar la latencia, permite que el sistema siga funcionando correctamente. </h3></div>