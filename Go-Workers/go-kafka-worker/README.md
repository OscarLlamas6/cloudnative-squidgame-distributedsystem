# Go Kafka Worker: USAC Squid Games - Distributed Cloud Native System

- Suscriptor de Kafka implementado en Golang


## Instalaciones necesarias:

```bash
#Golang dependencies
> go get github.com/segmentio/kafka-go
> go get github.com/joho/godotenv
> go get github.com/go-redis/redis/v8
> go get github.com/google/uuid
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
MONGO_HOST=<mongo-service-ip>
MONGO_PORT=<mongo-service-port>
MONGO_DB=<mongo-db-name>
MONGO_COL=<mongo-collection-name>

```

### Ejecutar suscriptor

```bash

#Para iniciar suscriptor
> go run suscriber.go

```

