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
#### Deployment
* Assumption: You already have a cluster setup and can execute commands with `kubectl`

* Make a YAML file, using the one from the Documentation is a good start:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx #This is the app name
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

#### Jobs
* Use the following YAML as reference
```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: pi #This is the job name
spec:
  template:
    spec:
      containers:
      - name: pi
        image: perl:5.34.0
        command: ["perl",  "-Mbignum=bpi", "-wle", "print bpi(2000)"]
      restartPolicy: Never
  backoffLimit: 4
```

* Deploy your Job
```shell
kubectl apply -f <yamlFile>
```
* Job will either Complete or Fail
* You can view logs of the Job like this

```shell
kubectl logs -l job-name=<jobName>
```


#### Possible Polling Techniques

```shell
kubectl get job ws01-test-job -n default -o jsonpath='{.status.conditions[0].type}'
```

```shell
kubectl logs -f --timestamps  \$(kubectl get pods -l job-name=ws01-test-job -o jsonpath='{.items[0].metadata.name}')
```
