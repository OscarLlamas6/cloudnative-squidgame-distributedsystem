# gRPC PubSub: USAC Squid Games - Distributed Cloud Native System

- API para guardar datos de un juego en bases de datos de Redis por medio de gRPC usando PubSub. 


## Instalaciones necesarias:

```bash
#Golang dependencies
> go get google.golang.org/grpc
> go get -u google.golang.org/grpc
> go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
> go get google.golang.org/protobuf/cmd/protoc-gen-go
> go get -u github.com/OscarLlamas6/grpc-helpers/protos/squidgame@3fee080cdaf278014e90fde74f6655a8b9513b2f
> go get github.com/joho/godotenv/cmd/godotenv

#PubSub
> go get go get -u cloud.google.com/go/pubsub


#Compilar archivo .proto
> protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative <.proto filename path>
```

## Variables de entorno necesarias

- Agregar un archivo llamado `.env` en la raiz de cada API con las siguientes variables de entorno:

```bash

# Variables de entorno para PubSub
PUBSUB_CLIENT_NAME="gRPC PubSub SQUID GAME Client :)"
PUBSUB_SERVER_NAME="gRPC PubSub SQUID GAMES Server :D"
PUBSUB_CLIENT_HOST=0.0.0.0:3039
PUBSUB_SERVER_HOST=0.0.0.0:6002

# Variables de entorno para bases de datos
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

# Imagen Cliente
> docker build -f .\Dockerfile.client -t oscarllamas6/grpc-pubsub-client:v1 .

# Imagen Servidor
> docker build -f .\Dockerfile.server -t oscarllamas6/grpc-pubsub-server:v1 .

```

### Ejecutar cliente y servidor

```bash

#Para iniciar cliente
> go run client/main.go

#Para iniciar servidor
> go run server/main.go
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