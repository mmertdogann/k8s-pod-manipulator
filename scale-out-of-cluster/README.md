 # out-of-cluster-scaler

A simple pod scaler golang application that can scale pod replicas outside the cluster via manipulating the deployment


<br>

## Setup & Installtion

<br>

```bash
git clone <repo-url>
```
<br>

## Running The App

> Note: You can use the -kubeconfig option to use a different config file. By default this program picks up the default file used by kubectl (when KUBECONFIG environment variable is not set).

To run the app with different config file, type this command:

```bash
go run main.go -kubeconfig="{absolute-path-of-kubeconfig-file}" 
```

To run the app with default config file, type this command:

```bash
go run main.go
```

> ⚠️ There should exist a dummy pod that crated with deployment manifest inside the cluster. I used this [repository](https://github.com/mmertdogann/k8s-node) to create a pod that used a deployment named nodeapp-deployment

Open a browser and then go to `localhost:3000` to make sure that application is running

Use `localhost:3000/scale/{replica-count}` to scale the specific replicas

Getting basic info about kubernetes components:

```bash
kubectl get node
kubectl get pod
kubectl get service
kubectl get all
```