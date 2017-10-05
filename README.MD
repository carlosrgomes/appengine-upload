Instalar Go:
https://golang.org/dl/

Instalar SDK Google:

https://cloud.google.com/sdk/docs/quickstart-windows?hl=pt-br

Instalar Appengine SDK

gcloud components install app-engine-go


Configurar o GOROOT do Intellij:
C:\Program Files (x86)\Google\Cloud SDK\google-cloud-sdk\platform\google_appengine\goroot-1.8

Configurar GOPATH do S.O

C:\workspaces-go

Configurar GOROOT do S.O
C:\Go\

Configurar o Appgine SDK GO no S.O
Variavel: GOSDKAPPENGINE
Valor: C:\Program Files (x86)\Google\Cloud SDK\google-cloud-sdk\platform\google_appengine

Exportar o path do gopath
%GOPATH%\bin
%GOSDKAPPENGINE%\bin

Subir local:
Entrar na pasta e digitar o comando:
    dev_appserver.py app.yaml

Deploy no Google
    gcloud app deploy

