 # in-cluster-scaler

A simple pod scaler golang application that can scale pod replicas inside the cluster via manipulating the deployment

<br>

## Setup & Installtion

<br>

```bash
git clone <repo-url>
```
<br>

> For acquiring the Docker image either pull the [image](https://hub.docker.com/r/mmertdogann/kubernetes-pod-scaler) or build a new image using Dockerfile via type this command:

```bash
docker build -t <image-name> .
```

⚠️ If you will use the newly created Docker image, you need to change the image property of the deployment manifest with the new `<image-name>`.

**Ingress part**:

In order to access the services from outside the cluster, we need to install Ingress Controller then create an ingress manifest. It allows us to reach services via a browser by just typing a URL and it also maps custom DNS names to IP addresses.

Install NGINX Ingress Controller for Docker Desktop from [here](https://kubernetes.github.io/ingress-nginx/deploy/)

To use custom DNS name to reach application via browser, we need to edit hosts, append `127.0.0.1 podmanagement.com` to hosts by using this command:

```bash
sudo nano /etc/hosts
```

<br>

## Running The App

To run the application, we need to create `deployment` and `ingress` from yaml files using these commands:

```bash
kubectl apply -f pod-scaler.deployment.yaml
kubectl apply -f pod-scaler.ingress.yaml
```

> ⚠️ There should exist a dummy pod that crated with deployment manifest inside the cluster. I used this [repository](https://github.com/mmertdogann/k8s-node) to create a pod that used a deployment named nodeapp-deployment

Open a browser and then go to `podmanagement.com` to make sure that application is running

`podmanagement.com/scale/{replica-count}` to scale the specific replicas

Getting basic info about kubernetes components:

```bash
kubectl get node
kubectl get pod
kubectl get service
kubectl get ingress
kubectl get all
```