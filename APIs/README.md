# APIs: USAC Squid Games - Distributed Cloud Native System

- APIs para guardar datos de un juego en bases de datos de Redis por medio de gRPC usando Kafka, RabbitMQ y PubSub. 


## Instalaciones necesarias:

```bash
#Golang dependencies
> go get google.golang.org/grpc
> go get -u google.golang.org/grpc
> go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
> go get google.golang.org/protobuf/cmd/protoc-gen-go
> go get -u github.com/OscarLlamas6/grpc-helpers/protos/squidgame@006352d75d8e1e7877b99cd456cb422f565ad504
> go get github.com/joho/godotenv/cmd/godotenv

#PubSub
> go get go get -u cloud.google.com/go/pubsub

#RabbitMQ
> go get github.com/streadway/amqp

#Kafka
> go get github.com/google/uuid
> go get github.com/segmentio/kafka-go

#Compilar archivo .proto
> protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative <.proto filename path>
```

## Variables de entorno necesarias

- Agregar un archivo llamado `.env` en la raiz de cada API con las siguientes variables de entorno:

```bash

# Variables de entorno para Kafka
KAFKA_CLIENT_NAME="gRPC Kafka SQUID GAME Client :)"
KAFKA_SERVER_NAME="gRPC Kafka SQUID GAMES Server :D"
KAFKA_CLIENT_HOST=0.0.0.0:3037
KAFKA_SERVER_HOST=0.0.0.0:6000
# Variables de entorno para RabbitMQ
RABBIT_CLIENT_NAME="gRPC RabbitMQ SQUID GAME Client :)"
RABBIT_SERVER_NAME="gRPC RabbitMQ SQUID GAMES Server :D"
RABBIT_CLIENT_HOST=0.0.0.0:3038
RABBIT_SERVER_HOST=0.0.0.0:6001
# Variables de entorno para PubSub
PUBSUB_CLIENT_NAME="gRPC PubSub SQUID GAME Client :)"
PUBSUB_SERVER_NAME="gRPC PubSub SQUID GAMES Server :D"
PUBSUB_CLIENT_HOST=0.0.0.0:3039
PUBSUB_SERVER_HOST=0.0.0.0:6002

#los valores de las ips y puertos son editables

```

### Ejecutar cliente y servidor

```bash

#Para iniciar cliente
> go run client/client.go

#Para iniciar servidor
> go run server/server.go
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