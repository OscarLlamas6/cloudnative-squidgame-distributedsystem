# Redis: USAC Squid Games - Distributed Cloud Native System

- Docker-compose para levantar servicio de Redis


## Variables de entorno necesarias


```bash
export REDISPASS=<redis-pass>
```

## Comandos

```bash
#Levantar servicio
> docker-compose up -d

#Detener servicio
> docker-compose down

#Acceder a redis-cli
> docker exec -it <redis_container_id> redis-cli

```




