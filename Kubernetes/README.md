# Kubernetes: USAC Squid Games - Distributed Cloud Native System

Instalaciones necesarias:

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
    
 ```