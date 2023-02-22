# Hexagonal Architecture | Fiber + MongoDB

- Docker-compose para levantar servicio de MongoDB

## Variables de entorno necesarias

## Linux
```bash

# Set
> export MONGO_INITDB_ROOT_USERNAME=<mongo-user>
> export MONGO_INITDB_ROOT_PASSWORD=<mongo-pass>
> export MONGO_INITDB_DATABASE=admin
> export dbUser=<mongo-user>
> export dbPwd=<mongo-pass>

# Echo (just for check if exists)
> echo $MONGO_INITDB_ROOT_USERNAME
```

## Windows
```powershell

# Set
> $env:MONGO_INITDB_ROOT_USERNAME='<mongo-user>'
> $env:MONGO_INITDB_ROOT_PASSWORD='<mongo-pass>'
> $env:MONGO_INITDB_DATABASE='admin'
> $env:dbUser='<mongo-user>'
> $env:dbPwd='<mongo-pass>'

# Echo (just for check if exists)
> $env:MONGO_INITDB_ROOT_USERNAME
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




