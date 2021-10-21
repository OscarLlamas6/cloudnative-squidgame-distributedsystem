# PubSub Worker: USAC Squid Games - Distributed Cloud Native System

- Suscriptor de Google PubSub implementado en Golang


## Instalaciones necesarias:

```bash
#Golang dependencies
> go get -u cloud.google.com/go/pubsub
> go get github.com/joho/godotenv
```

## Variables de entorno necesarias

- Agregar un archivo llamado `.env` en la raiz del proyecto con las siguientes variables de entorno:

```bash


# Variables de entorno para PubSub
TOPIC_NAME=projects/<projet-id>/topics/<topic-name>
PUBSUB_KEY_PATH=<relative path to private key json file>
SUB_NAME=projects/<project-id>/subscriptions/<sub-name>
PUBSUB_PROJECT=<project-id>
GOLANG_TOPIC=<topic-name>
GOLANG_SUB=<sub-name>

# En la raiz del proyecto ubicar el json de la private key 

```

### Ejecutar cliente y servidor

```bash

#Para iniciar cliente
> go run suscriber.go

```

### Generar una "Private Key" para PubSub en Google Cloud Platform

    1) Go to GCP and click on hamburger menu
    2) IAM & Admin -> Service Accounts (click)
    3) Create "Service Account"
    4) Fill name, id, and description of new service account
    5) Select role: Pub/Sub -> Pub/Sub Admin, then click "continue" and "DONE"
    6) Click on new service account and go to "KEYS"
    7) Click on "Add Key"->"Create new key", stay with json and click "CREATE"