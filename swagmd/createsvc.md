create/update a k8s service. the request body should specify the detailed configs of 
the service to be created/updated. labels of the pod should be specified in the selector

an example:

```json

{
  "name": "test-svc",
  "namespace": "test",
  "labels": [
    {
      "key": "test",
      "value": "svc"
    }
  ],
  "selector": [
    {
      "key": "app",
      "value": "test"
    }
  ],
  "ports": [
    {
      "name": "http",
      "targetPort": 80,
      "port": 80,
      "nodePort": 0
    }
  ]
}
```