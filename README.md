# Kubecenter


this is a web app for managing k8s resources

## push to /pull from harbor

```docker
docker build -t harbor.kubecenter.com/kubecenter/kubecenter:v1.1 .
docker push harbor.kubecenter.com/kubecenter/kubecenter:v1.1      
docker pull harbor.kubecenter.com/kubecenter/kubecenter:v1.1
docker run -d -p 7080:7080 harbor.kubecenter.com/kubecenter/kubecenter:v1.1
```


## push to gogs repo

```
git remote set-url --add origin http://VM_IP:3000/username/repository.git
git push
```


## config drone using docker-compose
```
volumes: 
  dronedata:
services:
  drone-server:
    image: drone/drone:2
    environment:
      DRONE_AGENTS_ENABLED: "true"
      DRONE_GOGS_SERVER: "http://ip:10880"
      DRONE_RPC_SECRET: "ib82ce97f2a4d0fb940bcd2983c4017bb"
      DRONE_SERVER_HOST: "ip:9080"
      DRONE_SERVER_PROTO: "http"
      DRONE_USER_CREATE: "username:kubecenter,machine:false,admin:true,token:55f24eb3d61ef6ac5e83d550178638dc"
    restart: always
    container_name: drone-server
    ports:
    - 9080:80
    - 9443:443
    volumes:
    - dronedata:/data
  drone-runner:
    image: drone/drone-runner-docker:1
    environment:
      DRONE_RPC_PROTO: "http"
      DRONE_RPC_HOST: "ip:9080"
      DRONE_RPC_SECRET: "ib82ce97f2a4d0fb940bcd2983c4017bb"
      DRONE_RUNNER_CAPACITY: "2"
      DRONE_RUNNER_NAME: "my-first-runner"
    ports:
    - 3000:3000
    restart: always
    container_name: drone-runner
    depends_on:
    - drone-server
    volumes:
    - /etc/docker/:/etc/docker
    - /var/run/docker.sock:/var/run/docker.sock
```