# MongoDB: USAC Squid Games - Distributed Cloud Native System

- Docker-compose para levantar servicio de MongoDB


## Variables de entorno necesarias


```bash
export MONGO_INITDB_ROOT_USERNAME=<mongo-user>
export MONGO_INITDB_ROOT_PASSWORD=<mongo-pass>
export MONGO_INITDB_DATABASE=admin
export dbUser=<mongo-user>
export dbPwd=<mongo-pass>
```

## Comandos

```bash
#Levantar servicio
> docker-compose up -d

#Detener servicio
> docker-compose down

#Acceder a mongo-cli
> sudo docker exec -it <mongo_container_id> mongo

```




