# Go PubSub Worker: USAC Squid Games - Distributed Cloud Native System

- Suscriptor de Google PubSub implementado en Golang


## Instalaciones necesarias:

```bash
#Golang dependencies
> go get -u cloud.google.com/go/pubsub
> go get github.com/joho/godotenv
> go get github.com/go-redis/redis/v8
> go get github.com/google/uuid
> go get go.mongodb.org/mongo-driver/mongo

#Pubsub
> go get -u cloud.google.com/go/pubsub

```

## Variables de entorno necesarias

- Agregar un archivo llamado `.env` en la raiz del proyecto con las siguientes variables de entorno:

```bash


# Setear variables
TOPIC_NAME=projects/<projet-id>/topics/<topic-name>
PUBSUB_KEY_PATH=<relative path to private key json file>
SUB_NAME=projects/<project-id>/subscriptions/<sub-name>
PUBSUB_PROJECT=<project-id>
GOLANG_TOPIC=<topic-name>
GOLANG_SUB=<sub-name>
REDIS_HOST=<redis-service-ip>
REDIS_PORT=<redis-service-port>
REDUS_PASS=<redis-pass>
MONGO_HOST=<mongo-service-ip>
MONGO_PORT=<mongo-service-port>
MONGO_DB=<mongo-db-name>
MONGO_COL=<mongo-collection-name>
MONGO_USER=<mongo-user>
MONGO_PASS=<mongo-pass>

# En la raiz del proyecto ubicar el json de la private key 

```

### Crear Docker images

```bash

# Imagen Docker
> docker build -t oscarllamas6/go-pubsub-worker:v1 .

```

### Ejecutar suscriptor

```bash

#Para iniciar suscriber
> go run main.go

```

### Generar una "Private Key" para PubSub en Google Cloud Platform

    1) Go to GCP and click on hamburger menu
    2) IAM & Admin -> Service Accounts (click)
    3) Create "Service Account"
    4) Fill name, id, and description of new service account
    5) Select role: Pub/Sub -> Pub/Sub Admin, then click "continue" and "DONE"
    6) Click on new service account and go to "KEYS"
    7) Click on "Add Key"->"Create new key", stay with json and click "CREATE"