# Kubecenter


this is a web app for managing k8s resources

## push to /pull from harbor

```docker
docker build -t harbor.kubecenter.com/kubecenter/kubecenter:v1.1 .
docker push harbor.kubecenter.com/kubecenter/kubecenter:v1.1      
docker pull harbor.kubecenter.com/kubecenter/kubecenter:v1.1
docker run -d -p 7080:7080 harbor.kubecenter.com/kubecenter/kubecenter:v1.1
```

