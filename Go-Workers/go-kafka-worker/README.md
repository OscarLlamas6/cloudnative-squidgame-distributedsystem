# Go Kafka Worker: USAC Squid Games - Distributed Cloud Native System

- Suscriptor de Kafka implementado en Golang


## Instalaciones necesarias:

```bash
#Golang dependencies
> go get github.com/segmentio/kafka-go
> go get github.com/joho/godotenv
> go get github.com/go-redis/redis/v8
> go get github.com/google/uuid
> go get go.mongodb.org/mongo-driver/mongo
```

## Variables de entorno necesarias

- Agregar un archivo llamado `.env` en la raiz del proyecto con las siguientes variables de entorno:

```bash


# Setear variables
KAFKA_HOST=<kafka-broker-service-ip>
KAFKA_PORT=<kafka-broker-service-port>
KAFKA_TOPIC=<kafka-topic>
REDIS_HOST=<redis-service-ip>
REDIS_PORT=<redis-service-port>
REDUS_PASS=<redis-pass>
MONGO_HOST=<mongo-service-ip>
MONGO_PORT=<mongo-service-port>
MONGO_DB=<mongo-db-name>
MONGO_COL=<mongo-collection-name>
MONGO_USER=<mongo-user>
MONGO_PASS=<mongo-pass>

```
### Crear Docker images

```bash

# Imagen Docker
> docker build -t oscarllamas6/go-kafka-worker:v1 .

```


### Ejecutar suscriptor

```bash

#Para iniciar suscriptor
> go run suscriber.go

```

