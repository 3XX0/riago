Riago
=====

Riago is a Riak client for Go.

Supported Operations
--------------------

- KV Get, Put, Del, GetBucket, SetBucket, ListBuckets, ListKeys
- 2i: Index
- MR: MapRed
- Search: SearchQuery
- Yokozuna: YokozunaIndexGet, YokozunaIndexPut, YokozunaIndexDelete, YokozunaSchemaGet, YokozunaSchemaPut

Example
-------

```go
package main

import "github.com/3XX0/riago"
import "github.com/3XX0/pooly"

func main() {
    conf := new(pooly.ServiceConfig)
    conf.Driver = riago.NewDriver()

    s := pooly.NewService("riak", conf)
    defer s.Close()

    s.Add("10.0.0.254:8087")

    c, err := s.GetConn()
    if err != nil {
            panic(err)
    }
    info, err := riago.RConn(c).ServerInfo()
    if err != nil {
            panic(err)
    }

    println(info.String())

    if err := c.Release(err, pooly.HostUp); err != nil {
            panic(err)
    }
}
```
