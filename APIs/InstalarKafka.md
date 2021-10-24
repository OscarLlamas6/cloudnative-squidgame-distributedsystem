# Local Kafka instance

- Guia para instalar Kafka localmente.

# Contenido
- [Windows](#windows) 
- [Linux](#linux)    

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
