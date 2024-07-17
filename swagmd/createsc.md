create a storage class, the request body should specify the detailed configs of
the storage class to be created. note that this project use nfs as the default provisioner. 
nfs supports two types of reclaim policy: Retain and Delete


an example:

```json
{
  "name": "test",
  "labels": [],
  "provisioner": "cluster.local/nfs-subdir-external-provisioner",
  "mountOptions": ["nfsvers=4"],
  "volumeBindingMode": "Immediate",
  "reclaimPolicy": "Delete",
  "allowVolumeExpansion": false,
  "parameters": []
}
```