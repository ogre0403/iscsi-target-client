# iscsi-target-client

A http clinet wrapper for [iscsi-target-api](https://github.com/ogre0403/iscsi-target-api.git).

## Curl Example

```bash
$ curl -XPOST \
  -d '{"type":"tgtimg","name":"test.img","size":10,"unit":"MiB","group":"test"}' \
  --user admin:password \
  http://127.0.0.1:8811/createVol

$ curl -XPOST \
  -d '{"targetIQN":"iqn.2017-07.k8s.default:myclaim", "volume": {"type":"tgtimg","name":"test.img","group":"test"}}' \
  --user admin:password \
  http://127.0.0.1:8811/attachLun

$ curl -XDELETE \
  -d '{"targetIQN":"iqn.2017-07.k8s.default:myclaim"}' \
  --user admin:password \
  http://127.0.0.1:8811/deleteTar

$ curl -XDELETE \
  -d '{"type":"tgtimg","name":"test.img","group":"test"}' \
  --user admin:password \
  http://127.0.0.1:8811/deleteVol
```

## Json body

* Volume
    ```json
    {
      "type": "tgtimg", 
      "group": "test",
      "name": "test.img",
      "size": 10,
      "unit": "MiB",
      "thin": false
    }
    ```

* Target 
    ```json
    {
      "targetIQN": "iqn.2017-07.k8s.default:myclaim", 
      "volume": {
          "type": "tgtimg",
          "group": "test",
          "name": "test.img"
      }
    }
    ```