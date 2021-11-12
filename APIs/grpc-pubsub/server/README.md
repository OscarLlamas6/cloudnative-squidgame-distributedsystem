# gRPC PubSub: USAC Squid Games - Distributed Cloud Native System

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

# pubsub:

> go get -u cloud.google.com/go/pubsub

#Compilar archivo .proto
> protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative <.proto filename path>
```

## Variables de entorno necesarias

- Agregar un archivo llamado `.env` en la raiz de cada API con las siguientes variables de entorno:

```bash

# Variables de entorno para PubSub
PUBSUB_CLIENT_NAME="gRPC PubSub SQUID GAME Client :)"
PUBSUB_SERVER_NAME="gRPC PubSub SQUID GAMES Server :D"
PUBSUB_CLIENT_HOST=localhost
PUBSUB_CLIENT_PORT=3039
PUBSUB_SERVER_HOST=localhost
PUBSUB_SERVER_PORT=6002
TOPIC_NAME=projects/<projet-id>/topics/<topic-name>
PUBSUB_KEY_PATH=<relative path to private key json file>
SUB_NAME=projects/<project-id>/subscriptions/<sub-name>
PUBSUB_PROJECT=<project-id>
GOLANG_TOPIC=<topic-name>
GOLANG_SUB=<sub-name>
#los valores de las ips y puertos son editables

```

### Crear Docker images

```bash

# Imagen Docker
> docker build -t oscarllamas6/grpc-pubsub-server:v1 .


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