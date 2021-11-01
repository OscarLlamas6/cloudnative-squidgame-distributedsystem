# Kubernetes: USAC Squid Games - Distributed Cloud Native System


# Contenido
- [Variables de entorno en archivos YAML](#variables-de-entorno-en-archivos-yaml)
- [Setear gcloud y kubectl](#setear-gcloud-y-kubectl) 
- [Apache Kafka con Strimzi](#apache-kafka-con-strimzi)


# Variables de entorno en archivos YAML

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
>  envsubst < deployment.yaml | kubectl apply -f -

# Creando un archivo nuevo con el resultado de sustituir las variables de entorno
> envsubst < deployment.yaml > nuevo_deployment.yaml
```

# Setear gcloud y kubectl

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


# Extra para WSL Windows
> cp /mnt/c/Users/oscar/.kube/config ~/.kube

 ```


 # Comandos kubectl

 ```bash

# Levantar servicios Pubsub
>  kubectl apply -f pubsub.yaml

# Levantar servicios RabbitMQ
>  kubectl apply -f rabbitmq.yaml

# Levantar servicios Kafka
>  kubectl apply -f kafka.yaml

# Levantar todos loser servicios
> kubectl delete -f .\pubsub.yaml -f .\kafka.yaml -f .\rabbitmq.yaml -f .\Ingress-Error-TrafficSpliting\config-error.yaml -f .\Ingress-Error-TrafficSpliting\ingress.yaml -f .\Ingress-Error-TrafficSpliting\traffic-splitting.yaml

 ```

 # Apache Kafka con Strimzi

- Instalar y setear Kafa con Strimzi
```bash
# Instalando Strimzi. Cambiar <namespace> por el nombre correcto.
> kubectl apply -f 'https://strimzi.io/install/latest?namespace=<namespace>'

# Chequear que el pod de Strimzi este corriendo. Cambiar <namespace> por el nombre correcto.
> kubectl get pods -n <namespace>
```

- Definiendo YAMLs para crear cluster, topics o cualquier otro recurso de Kafka deseado.

```yaml
# Ejemplo de un cluster simple
apiVersion: kafka.strimzi.io/v1beta2
kind: Kafka
metadata:
  name: squidgames-cluster
  namespace: squidgames
spec:
  kafka:
    version: 3.0.0
    replicas: 2
    listeners:
      - name: plain
        port: 9092
        type: internal
        tls: false
      - name: tls
        port: 9093
        type: internal
        tls: true
    config:
      offsets.topic.replication.factor: 2
      transaction.state.log.replication.factor: 2
      transaction.state.log.min.isr: 2
      log.message.format.version: '3.0'
      inter.broker.protocol.version: '3.0'
    storage:
      type: ephemeral
  zookeeper:
    replicas: 2
    storage:
      type: ephemeral
  entityOperator:
    topicOperator: {}
    userOperator: {}

```

```yaml
# Ejemplo de un Topic en Kafka
apiVersion: kafka.strimzi.io/v1beta2
kind: KafkaTopic
metadata:
  name: squidgames
  namespace: squidgames
  labels:
    strimzi.io/cluster: squidgames-cluster
spec:
  partitions: 10
  replicas: 2
  config:
    retention.ms: 604800000
    segment.bytes: 1073741824
```

Para más ejemplos e información de recurso de Kafka visitar [https://operatorhub.io/operator/strimzi-kafka-operator](https://operatorhub.io/operator/strimzi-kafka-operator=)


- Aplicar manifiestos YAMLs para crear objetos.

```bash
# Creando cluster definido en archivo strimzi.yaml
> kubectl apply -f strimzi.yaml

# Listar pods con watcher
> kubectl get pods -n squidgames -w

# Revisar servicios de Kafka en K8s
> kubectl get kafka -n squidgames

# Obtener info detallada de servicos de Kafka
> kubectl get kafka -n squidgames -o yaml

# Listar topics de Kafka 
> kubectl get kafkatopic -n squidgames

```