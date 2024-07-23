create a service account. the configurations of the secret should be specified in the request
body.

an example of request body:

```json
{
  "name": "test",
  "namespace": "test",
  "labels": [
    {
      "key": "sa",
      "value": "test"
    }
  ]
}
```