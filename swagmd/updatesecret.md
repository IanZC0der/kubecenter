create/update a secret. the configurations of the secret should be specified in the request
body.

an example of request body:


```json
{
  "name": "test",
  "namespace": "test",
  "type": "Opaque",
  "labels": [
    {
      "key": "cm",
      "value": "test"
    }
  ],
  "data": [
    {
      "key": "testKey",
      "value": "dGVzdFZhbHVl"
    }
  ]
}
```