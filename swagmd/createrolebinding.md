create a rolebinding. the configurations of the rolebinding should be specified in the request
body. if the namespace is empty, a cluster rolebinding will be created. service account and the role specified in
the request body should be created before creating the role binding

an example of request body:

```json
{
  "name": "test-rb",
  "namespace": "test",
  "labels": [
    {
      "key": "name",
      "value": "rb"
    }
  ],
  "roleRef": "test",
  "subjects": [
    {
      "name": "test",
      "namespace": "test"
    }
  ]
}
```