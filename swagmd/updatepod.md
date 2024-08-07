updates a Kubernetes pod with detailed configurations, the old pod (if exists) will be deleted, an example of pod config:

```json
{
  "base": {
    "name": "test",
    "namespace": "test",
    "labels": [
      {
        "key": "app",
        "value": "test"
      }
    ],
    "restartPolicy": "Always"
  },
  "volumes": [
    {
      "name": "cache-volume",
      "type": "emptyDir"
    }
  ],
  "netWorking": {
    "hostNetwork": true,
    "hostName": "test",
    "dnsPolicy": "Default",
    "dnsConfig": {
      "nameservers": [
        "8.8.8.8"
      ]
    },
    "hostAliases": [
      {
        "key": "64.23.172.139",
        "value": "foo.bar,foo2.bar"
      }
    ]
  },
  "initContainers": [
    {
      "name": "busybox",
      "image": "busybox",
      "imagePullPolicy": "IfNotPresent",
      "command": [
        "echo"
      ],
      "args": [
        "hello world"
      ]
    }
  ],
  "containers": [
    {
      "name": "nginx",
      "image": "nginx",
      "imagePullPolicy": "IfNotPresent",
      "privileged": true,
      "tty": true,
      "workingDir": "/test",
      "envs": [
        {
          "key": "foo",
          "value": "bar"
        }
      ],
      "startupProbe": {
        "enable": true,
        "type": "http",
        "httpGet": {
          "scheme": "HTTP",
          "host": "",
          "path": "/",
          "port": 80,
          "httpHeaders": [
            {
              "key": "foo",
              "value": "bar"
            }
          ]
        },
        "initialDelaySeconds": 10,
        "periodSeconds": 5,
        "timeoutSeconds": 10,
        "successThreshold": 1,
        "failureThreshold": 10
      },
      "livenessProbe": {
        "enable": true,
        "type": "tcp",
        "tcpSocket": {
          "host": "",
          "port": 80
        },
        "initialDelaySeconds": 10,
        "periodSeconds": 5,
        "timeoutSeconds": 10,
        "successThreshold": 1,
        "failureThreshold": 10
      },
      "readinessProbe": {
        "enable": true,
        "type": "exec",
        "exec": {
          "command": [
            "echo",
            "helloworld"
          ]
        },
        "initialDelaySeconds": 10,
        "periodSeconds": 5,
        "timeoutSeconds": 10,
        "successThreshold": 1,
        "failureThreshold": 10
      },
      "resources": {
        "enable": true,
        "memRequest": 128,
        "memLimit": 128,
        "cpuRequest": 100,
        "cpuLimit": 100
      },
      "volumeMounts": [
        {
          "mountName": "cache-volume",
          "mountPath": "/test",
          "readyOnly": false
        }
      ]
    }
  ]
}
```