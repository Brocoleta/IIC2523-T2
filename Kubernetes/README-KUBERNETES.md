# Parte 1 - Kubernetes

## 1 - Crear un clúster Kubernetes

Para crear un cluster se utiliza el comando `minikube start`, con ello se levanta un nodo de kubernetes el cual se puede comprobar corriendo `kubectl get nodes`

![Tutorial 1](/images/kubernetes1.png "Tutorial 1")

![Tutorial 2](/images/kubernetes2.png "Tutorial 2")


## 2 - Desplegar una app

Para crear un deployment se utiliza `kubectl create deployment name --image=appimagelocation`, luego podemos comprobar que el deployment se haya realizado con `kubectl get deployments`.

Si queremos acceder a red privada de los clusters debemos crear un proxy con `kubectl proxy`.

![Tutorial 3](/images/kubernetes3.png "Tutorial 3")

![Tutorial 4](/images/kubernetes4.png "Tutorial 4")


## 3 - Troubleshoot Kubernetes (get, describe, logs y exec).

Para obtener los pods que estan corriendo utlizamos `kubectl get pods` y para poder obtener información sobre estos pods corremos `kubectl describe pods` que contiene info como la IP, los puertos utilizados, eventos, etc...

![Tutorial 5](/images/kubernetes5.png "Tutorial 5")

Para poder ver los logs de un pod se utiliza `kubectl logs POD_NAME`, que contiene todo lo que se ha enviado como STDOUT y para ejecutar comandos directamene en un pod podemos usar `kubectl exec POD_NAME -- COMMAND`

![Tutorial 6](/images/kubernetes6.png "Tutorial 6")


## 4 - Exponer una aplicación públicamente.

Para obtener los servicios en el cluster utilizamos `kubectl get services` y para poder crear un nuevo servicio `kubectl expose NAME --type="TYPE" --port PORT`

![Tutorial 7](/images/kubernetes7.png "Tutorial 7")

Para obtener las labels de un deployment podemos usar el comando de describe mencionado en la parte 3, también para crear labels se utiliza `kubectl label pods POD_NAME label=value` y podemos hacer llamados a comandos pidiendo exclusivamente los pods con X label.

![Tutorial 8](/images/kubernetes8.png "Tutorial 8")

![Tutorial 9](/images/kubernetes9.png "Tutorial 9")

Podemos eliminar servicios con `kubetl delete service -l label=value`, lo que elimina el servicio pero la app sigue corriendo dentro del cluster ya que eso esta ligado al deployment, service solo expone la aplicación.

![Tutorial 10](/images/kubernetes10.png "Tutorial 10")


## 5 - Escalar una aplicación

Podemos obtener el ReplicaSet creado por un deployment con `kubectl get rs`, esto muestra 2 valores importantes, DESIRED que son las replicas "esperadas" en nuestra app y CURRENT que son las que realmente estan corriendo en el momento.

Para cambiar la cantidad de réplicar se utiliza el comanto `kubectl scale deployments/DEPLOYNAME --replicas=N`

![Tutorial 11](/images/kubernetes11.png "Tutorial 11")

Automáticamente kubernetes aumenta la cantidad de pods y crea un load-balancer para distribuir la carga.

![Tutorial 12](/images/kubernetes12.png "Tutorial 12")


## 6 - Actualizar una aplicación

Para actualizar la imagen de una aplicación, se utiliza `kubectl set image deployments/NAME IMAGE`, esto automáticamente comienza un rolling update.

![Tutorial 13](/images/kubernetes13.png "Tutorial 13")

Podemos verificar el estado de un rolling update con `kubectl rollout status deployments/NAME`.

![Tutorial 14](/images/kubernetes14.png "Tutorial 14")

Para hacer un rollback de un deploy, se utiliza `kubectl rollout undo deployments/NAME` 

![Tutorial 15](/images/kubernetes15.png "Tutorial 15")
