# Apache Kafka

- Guia para instalar Kafka.

# Contenido
- [Windows](#windows) 
- [Linux](#linux)    
- [Strimzi](#strimzi)

# Windows

  - Asegurarnos que Java 8 SDK está instalado

```bash
# la salida debe ser similar a esta
> java -version
java version "1.8.0_311"
Java(TM) SE Runtime Environment (build 1.8.0_311-b11)
Java HotSpot(TM) 64-Bit Server VM (build 25.311-b11, mixed mode)
```

  - Podemos instalarlo desde el instalador de la pagina oficial de Java o bien usar el manejador de packetes https://chocolatey.org/ y ejecutar el siguiente comando

```bash
> choco install jdk8
``` 

- Descargar los binarios de Kafka desde la pagina oficial de Kafka https://kafka.apache.org/downloads  
Link de descarga directa: https://archive.apache.org/dist/kafka/2.8.0/kafka_2.13-2.8.0.tgz

&nbsp;
<p align="center" >
  <img src="https://i.ibb.co/3BKJmsy/image.png" width="316" height="230" />
</p>
&nbsp;

- Situar carpeta en la raiz del disco local, renombrar la carpeta como "Kafka" y crear subcarpetas "kafka-logs" y "zookeeper-data" 

&nbsp;
<p align="center" >
  <img src="https://i.ibb.co/HdLr9kB/image.png" width="312" height="290" />
</p>
&nbsp;

- Cambiar configuraciones default

    -  Actualizar la ruta del directorio de datos de zookeeper en "config/zookeeper.properties"

    ```bash
    # the directory where the snapshot is stored.
    dataDir=c:/kafka/zookeeper-data
    # the port at which the clients will connect
    ```

    -  Actualizar la ruta del directorio de Logs de Karfa en "config/server.properties"

    ```bash
    # A comma separated list of directories under which to store log files
    log.dirs=c:/kafka/kafka-logs
    ```

    - 3 propiedades que deben ser unicas en cada broker

    ```bash

    #Archivo /config/server.properties

    # Id del broker, debe ser un numero entero para cada uno.
    broker.id=0

    # Direccion del socket del broker
    listeners=PLAINTEXT://:9092

    # Directorio para los logs del broker
    log.dirs=c:/kafka/kafka-logs
    ```
## Iniciar servicios

```bash
# Iniciar servidor zookeeper
> .\bin\windows\zookeeper-server-start.bat ..\..\config\zookeeper.properties

# Iniciar servidor broker
> .\bin\windows\kafka-server-start.bat ..\..\config\server.properties
```

## Crear topics

```bash
> bin/windows/kafka-topics.bat --create --topic squidgames --zookeeper localhost:2181 --partitions 2 --replication-factor 2

# --partitions = cantidad de brokers
# --replicationfactor = copias de los datos en caso de que un broker falle
# --zookeeper direccion del servicio Zookeeper 

#ejemplo
> bin/windows/kafka-topics.bat --create --topic squidgames --zookeeper localhost:2181 --partitions 1 --replication-factor 1
Created topic squidgames. #<-- Mensaje de confirmacion

```

## Listar Topics

```bash
# Consultando lista de Topics con el zookeeper en la direccion "localhost:2181"
> bin/windows/kafka-topics.bat --list --zookeeper localhost:2181
squidgames
```

## Ver detalles de un topic en específico

```bash
# Consultando detalles de Topic "squidgames" con el zookeeper en la direccion "localhost:2181"
> bin/windows/kafka-topics.bat --describe --topic squidgames --zookeeper localhost:2181
```

## Enviar mensaje a un Topic

```bash
# Enviando mensaje a Topic "squidgames" del broker "localhost:9092"
> bin/windows/kafka-console-producer.bat --broker-list localhost:9092 --topic squidgames

# --broker-list = lista de direcciones de los brokers (ej: localhost:9093,localhost:9094,localhost:9095)

# --topic = nombre del Topic

```
## Consumir mensajes de un Topic

```bash
# Consumiendo mensajes del Topic "squidgames" del broker "localhost:9092"
> bin/windows/kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic squidgames --from-beginning

# --bootstrap-server direccion del broker del consumiremos los mensajes

# --topic = nombre del Topic

# --from-beginning = Leer todos los mensajes existentes. incluso aquellos que ya consumimos antes.

```

# Linux

- Asegurarnos que Java 8 SDK está instalado

```bash
# la salida debe ser similar a esta
> java -version
java version "1.8.0_311"
Java(TM) SE Runtime Environment (build 1.8.0_311-b11)
Java HotSpot(TM) 64-Bit Server VM (build 25.311-b11, mixed mode)
```

  - En caso de no tenerlo instalador, ejectuar los siguientes comandos:

```bash
> sudo apt-get update
> sudo apt-get install openjdk-8-jdk
``` 

- Descargar los binarios de Kafka desde la pagina oficial de Kafka https://kafka.apache.org/downloads  
Link de descarga directa: https://archive.apache.org/dist/kafka/2.8.0/kafka_2.13-2.8.0.tgz

- Descomprimir en una ubicacion de nuestra eleccion con el siguiente comando:

```bash
> tar -xvzf kafka_2.13-2.8.0.tgz
```

## Iniciar servicios

```bash
# Iniciar servidor zookeeper
> .\bin\zookeeper-server-start.sh ..\..\config\zookeeper.properties

# Iniciar servidor broker
> .\bin\kafka-server-start.sh ..\..\config\server.properties
```

## Crear topics

```bash
> .\bin\kafka-topics.sh --create --topic squidgames --zookeeper localhost:2181 --partitions 2 --replication-factor 2

# --partitions = cantidad de brokers
# --replicationfactor = copias de los datos en caso de que un broker falle
# --zookeeper direccion del servicio Zookeeper 

#ejemplo
> .\bin\kafka-topics.sh --create --topic squidgames --zookeeper localhost:2181 --partitions 1 --replication-factor 1
Created topic squidgames. #<-- Mensaje de confirmacion

```

## Listar Topics

```bash
# Consultando lista de Topics con el zookeeper en la direccion "localhost:2181"
> .\bin\kafka-topics.sh --list --zookeeper localhost:2181
squidgames
```

## Ver detalles de un topic en específico

```bash
# Consultando detalles de Topic "squidgames" con el zookeeper en la direccion "localhost:2181"
> .\bin\windows\kafka-topics.sh --describe --topic squidgames --zookeeper localhost:2181
```

## Enviar mensaje a un Topic

```bash
# Enviando mensaje a Topic "squidgames" del broker "localhost:9092"
> .\bin\windows\kafka-console-producer.sh --broker-list localhost:9092 --topic squidgames

# --broker-list = lista de direcciones de los brokers (ej: localhost:9093,localhost:9094,localhost:9095)

# --topic = nombre del Topic

```
## Consumir mensajes de un Topic

```bash
# Consumiendo mensajes del Topic "squidgames" del broker "localhost:9092"
> .\bin\kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic squidgames --from-beginning

# --bootstrap-server direccion del broker del consumiremos los mensajes

# --topic = nombre del Topic

# --from-beginning = Leer todos los mensajes existentes. incluso aquellos que ya consumimos antes.
```

# Apache Kafka con Strimzi

- Instalar y setear Kafa con Strimzi
```bash
# Instalando Strimzi. Cambiar <namespace> por el nombre correcto.
> kubectl apply -f 'https://strimzi.io/install/latest?namespace=<namespace>'

# Chequear que el pod de Strimzi este corriendo. Cambiar <namespace> por el nombre correcto.
> kubectl get pods -n <namespace>
```

- Definiendo YAMLs para crear cluster, topics o cualquier otro recurso de Kafka deseado.

```yaml
# Ejemplo de un cluster simple
apiVersion: kafka.strimzi.io/v1beta2
kind: Kafka
metadata:
  name: squidgames-cluster
  namespace: squidgames
spec:
  kafka:
    version: 3.0.0
    replicas: 2
    listeners:
      - name: plain
        port: 9092
        type: internal
        tls: false
      - name: tls
        port: 9093
        type: internal
        tls: true
    config:
      offsets.topic.replication.factor: 2
      transaction.state.log.replication.factor: 2
      transaction.state.log.min.isr: 2
      log.message.format.version: '3.0'
      inter.broker.protocol.version: '3.0'
    storage:
      type: ephemeral
  zookeeper:
    replicas: 2
    storage:
      type: ephemeral
  entityOperator:
    topicOperator: {}
    userOperator: {}

```

```yaml
# Ejemplo de un Topic en Kafka
apiVersion: kafka.strimzi.io/v1beta2
kind: KafkaTopic
metadata:
  name: squidgames
  namespace: squidgames
  labels:
    strimzi.io/cluster: squidgames-cluster
spec:
  partitions: 10
  replicas: 2
  config:
    retention.ms: 604800000
    segment.bytes: 1073741824
```

Para más ejemplos e información de recurso de Kafka visitar [https://operatorhub.io/operator/strimzi-kafka-operator](https://operatorhub.io/operator/strimzi-kafka-operator=)


- Aplicar manifiestos YAMLs para crear objetos.

```bash
# Creando cluster definido en archivo strimzi.yaml
> kubectl apply -f strimzi.yaml

# Listar pods con watcher
> kubectl get pods -n squidgames -w

# Revisar servicios de Kafka en K8s
> kubectl get kafka -n squidgames

# Obtener info detallada de servicos de Kafka
> kubectl get kafka -n squidgames -o yaml

# Listar topics de Kafka 
> kubectl get kafkatopic -n squidgames

```