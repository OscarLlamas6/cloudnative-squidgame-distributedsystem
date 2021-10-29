# gRPC RabbitMQ: USAC Squid Games - Distributed Cloud Native System

- Servidor gRPC. 


## Instalaciones necesarias:

```bash
#Golang dependencies
> go get google.golang.org/grpc
> go get -u google.golang.org/grpc
> go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
> go get google.golang.org/protobuf/cmd/protoc-gen-go
> go get -u github.com/OscarLlamas6/grpc-helpers/protos/squidgame@3fee080cdaf278014e90fde74f6655a8b9513b2f
> go get github.com/joho/godotenv/cmd/godotenv

#RabbitMQ
> go get github.com/streadway/amqp

#Compilar archivo .proto
> protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative <.proto filename path>
```

## Variables de entorno necesarias

- Agregar un archivo llamado `.env` en la raiz de cada API con las siguientes variables de entorno:

```bash

# Variables de entorno para RabbitMQ
RABBIT_CLIENT_NAME="gRPC RabbitMQ SQUID GAME Client :)"
RABBIT_SERVER_NAME="gRPC RabbitMQ SQUID GAMES Server :D"
RABBIT_CLIENT_HOST=localhost
RABBIT_CLIENT_PORT=3038
RABBIT_SERVER_HOST=localhost
RABBIT_SERVER_PORT=6001

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

#los valores de las ips y puertos son editables

```

### Crear Docker images

```bash

# Imagen Docker
> docker build -t oscarllamas6/grpc-rabbit-server:v1 .

```

### Ejecutar servidor

```bash

#Para iniciar servidor
> go run main.go
```


### Ejemplo de JSON de entrada para guardar un juego con gRPC

```json
{
    "gamenumber": "1",
    "gamename": "Red Light, Green Light",
    "players": 16,
    "rungames": 100,
    "concurrence": 5,
    "timeout": 180
}
```