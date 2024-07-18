create/update a statefulSet. the configs of the statefulSet should be specified 
in the request body. use the createService api to create a service first.

an example:
```json

{
    "name": "nginx-statefulset",
    "namespace": "test",
    "labels": [
      {
        "key": "app",
        "value": "nginx-statefulset"
      }
    ],
    "selector": [
      {
        "key": "app",
        "value": "nginx-statefulset-tp"
      }
    ],
    "replicas": 2,
    "volumeClaimTemplates": [
      {
        "name": "www",
        "labels": [
          {
            "key": "pvc",
            "value": "pvc01"
          }
        ],
        "accessModes": [
          "ReadWriteOnce"
        ],
        "capacity": 100,
        "storageClassName": "nfs-client"
      }
    ],
  "template": {
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
    "base": {
      "labels": [
        {
          "key": "app",
          "value": "nginx-statefulset-tp"
        }
      ],
      "restartPolicy": "Always"
    },
    "initContainers": [
      {
        "name": "busybox-init",
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
        "envs": [
          {
            "name": "foo",
            "value": "bar"
          }
        ],
        "volumeMounts": [
          {
            "mountName": "www",
            "mountPath": "/usr/share/nginx/html",
            "readOnly": false
          }
        ]
      },
      {
        "name": "busybox",
        "image": "busybox",
        "imagePullPolicy": "IfNotPresent",
        "tty": true
      }
    ]
  }
}
```