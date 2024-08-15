# K8S Deployment Example

Learn how to deploy backend services into Kubernetes.

```
docker build -t hello-service:v1 .
docker tag hello-service:v1 registry.digitalocean.com/do-container/hello-service:v1
docker push registry.digitalocean.com/do-container/hello-service:v1

Change deployment.yaml
```
registry.digitalocean.com/do-container/hello-service:v1

### Deploy services

```
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml


