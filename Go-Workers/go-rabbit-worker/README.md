# Go RabbitMQ Worker: USAC Squid Games - Distributed Cloud Native System

- Suscriptor de RabbitMQ implementado en Golang


## Instalaciones necesarias:

```bash
#Golang dependencies
> go get github.com/streadway/amqp
> go get github.com/joho/godotenv
> go get github.com/go-redis/redis/v8
> go get github.com/google/uuid
> go get go.mongodb.org/mongo-driver/mongo
```

## Variables de entorno necesarias

- Agregar un archivo llamado `.env` en la raiz del proyecto con las siguientes variables de entorno:

```bash


# Setear variables
RABBITMQ_HOST=<rabbitmq-service-ip>
RABBITMQ_PORT=<rabbitmq-service-port>
RABBITMQ_USER=<rabbitmq-user>
RABBITMQ_PASS=<rabbitmq-pass>
RABBITMQ_QUEUE=<rabbitmq-queuename>
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

### Ejecutar suscriptor

```bash

#Para iniciar cliente
> go run suscriber.go

```

