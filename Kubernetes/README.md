# Kubernetes: USAC Squid Games - Distributed Cloud Native System


## Variables de entorno

-   Crear un archivo variables.conf con las variables de entorno, ejemplo:

```bash
REDIS_HOST=<redis-service-ip>
REDIS_PORT=<redis-service-port>
REDIS_PASS=<redis-db-password>
```

-   Crear un archivo config.sh con el siguiente codigo:

```bash
while read line; do export "$line";
done < variables.conf
echo "done"
```

- Instalar envsubst

```bash
> apt-get install gettext-base
```

- Ejemplos de uso.

```bash

# Leyendo y seteando variables de entorno
> . ./config.sh

# Aplicando un manifiesto de Kubernetes luego de sustituir variables de entorno
>  source .env
>  envsubst < deployment.yaml | kubectl apply -f -

# Creando un archivo nuevo con el resultado de sustituir las variables de entorno
> source .env
> envsubst < deployment.yaml > nuevo_deployment.yaml
```

# Contenido
- [Windows](#windows) 
- [Linux](#linux)    

# Windows

 - Instalar gcloud, correr el siguiente comando en Windows Powerll y ejecutar instalador

 ```bash
> (New-Object Net.WebClient).DownloadFile("https://dl.google.com/dl/cloudsdk/channels/rapid/GoogleCloudSDKInstaller.exe", "$env:Temp\GoogleCloudSDKInstaller.exe")

& $env:Temp\GoogleCloudSDKInstaller.exe

#Correr comando para iniciar configuracion
> gcloud init

#Se mostrara un mensaje como el siguiente, darle "Y" para aceptar y loggearnos en gcp

Network diagnostic detects and fixes local network connection issues.
Checking network connection...done.
Reachability Check passed.
Network diagnostic passed (1/1 checks passed).

You must log in to continue. Would you like to log in (Y/n)? Y

# Se mostraran una lista de proyectos, escogemos el proyecto o creamos uno nuevo.

You are logged in as: [<your-account-email>].

Pick cloud project to use:
 [1] ayd2g15
 [2] basic-perigee-325800
 [3] sopes-proyecto2-329600
 [4] Create a new project
Please enter numeric choice or text value (must exactly match list item): 2

#Configuramos una region y zona predeterminada, ejemplo us-central1-a

> gcloud config set compute/zone us-central1-a

# Verificamos que haya sido configurada correctamente

> gcloud config list compute/zone

# Instalamos kubectl

> gcloud components install kubectl

# Creamos cluster kubernetes

> gcloud container clusters create squidgames --num-nodes=3 --tags=allin,allout --machine-type=n1-standard-2 --no-enable-network-policy

#Recuperando credenciales para Kubectl
> gcloud container clusters get-credentials k8s-demo --zone=us-central1-c

#Permisos necesarios para el ingress controlers
> kubectl create clusterrolebinding cluster-admin-binding --clusterrole cluster-admin --user $(gcloud config get-value account)  


#Creamos reglas de firewall para los puertos
> gcloud compute firewall-rules create fwrule-kubernetes --allow tcp:30000-32767 

 ```


 # Comandos kubectl

 ```bash

# Levantar servicios Pubsub
>  kubectl apply -f pubsub.yaml

 ```