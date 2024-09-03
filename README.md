### How to Run Docker Container 
```shell
# Run this at the spot where Dockerfile is located
docker build -t <name:version> . 
```

* Case `CMD ["sh", "-c", "sleep infinity"]`
```shell
# Start the container, process will run forever.
docker run -d --name <container-name> <name:version>

# Run tests
docker exec <container-name> sh -c "go test -v"
```

* Case `CMD ["sh", "-c", "go test -v"]`
```shell
docker run <name:version>
```


### How to Push Images to DockerHub (Public)
* Assumption, you are logged in already
```shell
docker tag <name>:<version> <dockerId>/<name>:<version>

docker push <dockerId>/<name>:<version>
```

### How to Run Pod in Kubernetes
* Assumption: You already have a cluster setup and can execute commands with `kubectl`

* Make a YAML file, using the one from the Documentation is a good start:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
```

* Deploy your Pod
```shell
kubectl apply -f <yamlFile>
```

* Case `CMD ["sh", "-c", "sleep infinity"]`
```shell
# Run tests in pod
kubectl exec -it $(kubectl get pod -l app=<appName> -o jsonpath="{.items[0].metadata.name}") -- sh -c "go test -v"
```
