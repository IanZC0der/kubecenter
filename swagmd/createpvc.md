create a persistent volume claim, the request body should specify the detailed configs of 
the persistent volume claim. the requested capacity should be more than the target persistent volume capacity


an example:


```json
{
  "name": "test",
  "namespace": "test",
  "labels": [
    {
      "key": "pvc",
      "value": "pvc01"
    }
  ],
  "selector": [
    {
      "key": "pv",
      "value": "test-pv"
    }
  ],
  "accessModes": [
    "ReadWriteOnce"
  ],
  "capacity": 100
}
```