create/update a config map. the configurations of the configmap should be specified in the request
body. 

an example of request body:


```json
{
  "name": "test",
  "namespace": "test",
  "labels": [
    {
      "key": "cm",
      "value": "test"
    },
    {
      "key": "cm2",
      "value": "test2"
    }
  ],
  "data": [
    {
      "key": "testKey",
      "value": "testValue"
    },
    {
      "key": "db_name",
      "value": "testdb"
    },
    {
      "key": "db_host",
      "value": "127.0.0.1"
    }
  ]
}
```