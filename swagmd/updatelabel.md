updates the labels of a node, request body should specify the node name and the labels. old labels will be deleted.

an example of the request body:

```json
{
  "name": "ubuntu-s-2vcpu-4gb-sfo3-02",
  "labels": [
    {
      "key": "test",
      "value": "app"
    }
  ]
}
```