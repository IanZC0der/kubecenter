updates the taints of a node, request body should specify the node name and the taints. old taints will be deleted.

an example of the request body:

```json
{
  "name": "ubuntu-s-2vcpu-4gb-sfo3-02",
  "taints": [
    {
      "key": "test2",
      "value": "app2",
      "effect": "NoSchedule"
    }
  ]
}
```