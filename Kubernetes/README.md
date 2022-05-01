# Kubernetes: USAC Squid Games - Distributed Cloud Native System


# Contenido
- [Variables de entorno en archivos YAML](#variables-de-entorno-en-archivos-yaml)
- [Setear gcloud y kubectl](#setear-gcloud-y-kubectl) 
- [Apache Kafka con Strimzi](#apache-kafka-con-strimzi)
- [Linkerd](#linkerd)
- [Chaos Mesh](#chaos-mesh)


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

# Correr comando para iniciar configuración
> gcloud init

#Se mostraraá un mensaje como el siguiente, darle "Y" para aceptar y loggearnos en GCP

Network diagnostic detects and fixes local network connection issues.
Checking network connection...done.
Reachability Check passed.
Network diagnostic passed (1/1 checks passed).

You must log in to continue. Would you like to log in (Y/n)? Y

# Se mostrará una lista de proyectos, escogemos el proyecto o creamos uno nuevo.

You are logged in as: [<your-account-email>].

Pick cloud project to use:
 [1] ayd2g15
 [2] basic-perigee-325800
 [3] sopes-proyecto2-329600
 [4] Create a new project
Please enter numeric choice or text value (must exactly match list item): 2

# Configuramos una región y zona predeterminada, ejemplo us-central1-a

> gcloud config set compute/zone us-central1-a

# Verificamos que haya sido configurada correctamente

> gcloud config list compute/zone

# Instalamos kubectl

> gcloud components install kubectl

# Creamos cluster kubernetes

> gcloud container clusters create squidgames --num-nodes=3 --no-enable-ip-alias

# Recuperando credenciales para Kubectl
> gcloud container clusters get-credentials k8s-demo --zone=us-central1-c

# Permisos necesarios para el ingress controlers
> kubectl create clusterrolebinding cluster-admin-binding --clusterrole cluster-admin --user $(gcloud config get-value account)  


# Creamos reglas de firewall para los puertos
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

# Levantar todos los servicios
> kubectl apply -f .\pubsub.yaml -f .\kafka.yaml -f .\rabbitmq.yaml -f .\Ingress-Error-TrafficSpliting\config-error.yaml -f .\Ingress-Error-TrafficSpliting\ingress.yaml -f .\Ingress-Error-TrafficSpliting\traffic-splitting.yaml

# Levantar servicios y dashboard
> kubectl apply -f .\pubsub.yaml -f .\kafka.yaml -f .\rabbitmq.yaml -f .\dashboard.yaml -f .\Ingress-Error-TrafficSpliting\config-error.yaml -f .\Ingress-Error-TrafficSpliting\ingress-completo.yaml -f .\Ingress-Error-TrafficSpliting\traffic-splitting.yaml

# Borrar todos los servicios
> kubectl delete -f .\pubsub.yaml -f .\kafka.yaml -f .\rabbitmq.yaml -f .\dashboard.yaml -f .\Ingress-Error-TrafficSpliting\config-error.yaml -f .\Ingress-Error-TrafficSpliting\ingress.yaml -f .\Ingress-Error-TrafficSpliting\traffic-splitting.yaml

# Borrar servicios y dashboard
> kubectl delete -f .\pubsub.yaml -f .\kafka.yaml -f .\rabbitmq.yaml -f .\dashboard.yaml -f .\Ingress-Error-TrafficSpliting\config-error.yaml -f .\Ingress-Error-TrafficSpliting\ingress-completo.yaml -f .\Ingress-Error-TrafficSpliting\traffic-splitting.yaml

 ```

 # Apache Kafka con Strimzi

- Instalar y setear Kafka con Strimzi
```bash
# Instalando Strimzi. Cambiar <namespace> por el nombre correcto.
> kubectl apply -f 'https://strimzi.io/install/latest?namespace=<namespace>'

# Chequear que el pod de Strimzi esté corriendo. Cambiar <namespace> por el nombre correcto.
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

Para más ejemplos e información de recurso de Kafka visitar [https://operatorhub.io/operator/strimzi-kafka-operator](https://operatorhub.io/operator/strimzi-kafka-operator)


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

# Linkerd

- Comandos para instalar y setear Linkerd

```bash

# Instalando Helm
> sudo curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
> sudo chmod 700 get_helm.sh
> sudo bash ./get_helm.sh

# Instalando Ingress-Controller
> kubectl create ns nginx-ingress
> helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx 
> helm repo update 
> helm install nginx-ingress ingress-nginx/ingress-nginx -n nginx-ingress

# Obtener IP publica del ingress-controller
> kubectl get services -n nginx-ingress

# Instalando Docker (Opcional para construir imagenes)
> sudo apt-get install docker.io
> sudo usermod -aG docker developer

# Desintalar kubectl (por si es necesario hacer downgrade)
# Si se instalo desde curl
> sudo rm /usr/local/bin/kubectl
# Si se instalo como un componente de gcloud
> gcloud components remove kubectl
# Si se instalo con snap
> snap remove kubectl
# Si se desea borrar la configuracion anterior, borrar config en ~/.kube

# Instalando version especifica de kubectl
> curl -LO https://dl.k8s.io/release/v1.20.0/bin/linux/amd64/kubectl
> sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

# Instalando Linkerd
> curl -fsL https://run.linkerd.io/install | sh
> nano ~/.bashrc <- export PATH=$PATH:/home/YOUR_USER/.linkerd2/bin
> linkerd check
> linkerd install | kubectl apply -f -
> linkerd check
> linkerd viz install | kubectl apply -f -
> linkerd check

# Abrir dashboard de Linkerd
> linkerd viz dashboard

# Inyectando Linkerd en Ingress-Controller
> kubectl get deployment nginx-ingress-ingress-nginx-controller -n nginx-ingress  -o yaml | linkerd inject - | kubectl apply -f -
```

- Ejemplo Ingress

```yaml
# Ingress Controller
apiVersion: networking.k8s.io/v1 
kind: Ingress 
metadata: 
  name: minimal-ingress
  annotations: 
    kubernetes.io/ingress.class: nginx 
    nginx.ingress.kubernetes.io/rewrite-target: / 
    nginx.ingress.kubernetes.io/service-upstream: "true"
  namespace: squidgames
spec: 
  rules: 
  - host: ${IC_HOST}
    http: 
      paths: 
      - path: /
        pathType: Prefix 
        backend: 
          service: 
            name: grpc-pubsub-client-service 
            port: 
              number: 3039 
```

- Ejemplo Error-Injector

```yaml
# Config Map Error-Injector
apiVersion: v1
kind: ConfigMap
metadata:
  name: error-injector
  namespace: squidgames
data:
 nginx.conf: |-
    events {}
    http {
      server {
        listen 8080;
          location / {
            return 500;
          }
      }
    }

---

# Deployment Error-Injector
apiVersion: apps/v1
kind: Deployment
metadata:
  name: error-injector
  namespace: squidgames
  labels:
    app: error-injector
spec:
  selector:
    matchLabels:
      app: error-injector
  replicas: 1
  template:
    metadata:
      labels:
        app: error-injector
    spec:
      containers:
        - name: nginx
          image: nginx:alpine
          volumeMounts:
            - name: nginx-config
              mountPath: /etc/nginx/nginx.conf
              subPath: nginx.conf
      volumes:
        - name: nginx-config
          configMap:
            name: error-injector
---

# Service Error-Injector
apiVersion: v1
kind: Service
metadata:
  name: error-injector-service
  namespace: squidgames
spec:
  type: ClusterIP
  ports:
  - port: 3039 
    protocol: TCP
    targetPort: 8080
  selector:
    app: error-injector
```

- Ejemplo de TrafficSpliting con Linkerd

```yaml
# Traffic Split
apiVersion: split.smi-spec.io/v1alpha1
kind: TrafficSplit
metadata:
  name: service-split
  namespace: squidgames
spec:
  service: grpc-pubsub-client-service
  backends:
  - service: grpc-pubsub-client-service
    weight: 300m
  - service: grpc-rabbit-client-service
    weight: 300m
  - service: grpc-kafka-client-service
    weight: 300m
  - service: error-injector-service
    weight: 100m


```

# Chaos Mesh

- Comandos para instalar y setear Chaos Mesh

```bash
# Verificar las versiones del cliente y el servidor, para poder instalar chaos mesh
#la version del cliente debe ser igual o una version menos que la del servidor, ejemplo
> Kubectl
Client Version: version.Info{Major:"1", Minor:"20", GitVersion:"v1.20.0", GitCommit:"af46c47ce925f4c4ad5cc8d1fca46c7b77d13b38", GitTreeState:"clean", BuildDate:"2020-12-08T17:59:43Z", GoVersion:"go1.15.5", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"20+", GitVersion:"v1.20.10-gke.301", GitCommit:"17ad7bd6afa01033d7bd3f02ce5de56f940a915d", GitTreeState:"clean", BuildDate:"2021-08-24T05:18:54Z", GoVersion:"go1.15.15b5", Compiler:"gc", Platform:"linux/amd64"}

# En este caso el Cliente y Servidor son versiones 1.20 por lo cual, si podemos instalar Chaos Mesh

# Instalar Chaos Mesh
> sudo curl -sSL https://mirrors.chaos-mesh.org/v2.0.3/install.sh | bash

# Verificar instalación de Chaos Mesh
> kubectl get pod -n chaos-testing

# Abrir dashboard de Chaos Mesh
> kubectl port-forward -n chaos-testing svc/chaos-dashboard 2333:2333

# Desinstalar Chaos Mesh
> sudo curl -sSL https://mirrors.chaos-mesh.org/v0.9.1/install.sh | bash -s -- --template | kubectl delete -f -

```

- Ejemplo de un experimento en Chaos Mesh

```yaml
kind: Schedule
apiVersion: chaos-mesh.org/v1alpha1
metadata:
  name: experiment1
  namespace: chaos-testing
#  annotations:
#    experiment.chaos-mesh.org/pause: 'false'
spec:
  schedule: '@every 15s'
  startingDeadlineSeconds: null
  concurrencyPolicy: Forbid
  historyLimit: 1
  type: PodChaos
  podChaos:
    selector:
      namespaces:
        - squidgames
      labelSelectors:
        app: grpc-pubsub-client
    mode: one
    action: pod-kill
    duration: 2m
    gracePeriod: 0
```

- Comandos para aplicar y ver efectos del experimento

```bash
# Aplicando experimento para matar un pod
> kubectl apply -f chaos_mesh/pod-experiment.yaml

# Ver efectos del experimento en tiempo real
> kubectl get pods -n squidgame -w

```
