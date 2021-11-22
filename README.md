## Publish docker image
```
$ docker login
$ sbt
> docker:publish
```

## Test kubejob
```
$ minikube start
$ kubejob --file job.yaml
```


