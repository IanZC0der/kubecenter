create a role. the configurations of the role should be specified in the request
body. if the namespace is empty, a cluster role will be created

an example of request body:

```json
{
  "name": "test",
  "namespace": "test",
  "labels": [
    {
      "key": "role",
      "value": "test"
    }
  ],
  "rules": [
    {
      "apiGroups": [""],
      "verbs": ["get","list"],
      "resources": ["pods"],
      "resourceNames": ["web"],
      "nonResourceURLs": []
    }
  ]
}
```