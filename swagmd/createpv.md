create a persistent volume, the request body should specify the detailed configs of 
the persistent volume.


an example:


```json
{
  "name": "test",
  "namespace": "test",
  "labels": [
    {
      "key": "pv",
      "value": "test-pv"
    }
  ],
  "capacity": 100,
  "accessModes": [
    "ReadWriteOnce"
  ],
  "reClaimPolicy": "Recycle",
  "volumeSource": {
    "type": "nfs",
    "nfsVolumeSource": {
      "nfsPath": "",
      "nfsServer": "",
      "nfsReadyOnly": false
    }
  }
}
```