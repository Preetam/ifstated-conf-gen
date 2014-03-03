ifstated-conf-gen
====

This is a really simple tool to generate `/etc/ifstated.conf` files on OpenBSD.

Usage
---
First, create a JSON with the necessary fields. An example is provided as `example.json`:

```json
{
  "Device": "em0",
  "Vlans": [
    {
      "Id": 1,
      "Address": "10.0.0.1",
      "Netmask": "255.255.255.0"
    },
    {
      "Id": 2,
      "Address": "10.0.1.1",
      "Netmask": "255.255.255.0"
    }
  ]
}
```

All you have to do is pipe the JSON configuration and the generated `ifstated.conf` file
will be printed to stdout.

`$ cat example.json | go run main.go template.go`

Of course, you can also build a dependency-free binary and deploy it across
all of your server.

`$ go build`

License
---
MIT
