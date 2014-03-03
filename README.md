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

```
# This file is generated automatically. DO NOT EDIT.
init-state auto
fw_carp_up = "carp0.link.up"
fw_carp_init = "carp0.link.unknown"

state auto {
   if ($fw_carp_init)
      run "sleep 10"
   if ($fw_carp_up)
      set-state fw_master
   if (! $fw_carp_up)
      set-state fw_slave
}

state fw_master {
   init {
      run "ifconfig em0 up"
      
      run "ifconfig vlan1 create"
      run "ifconfig vlan1 10.0.0.1 netmask 255.255.255.0 vlan 1 vlandev em0"
      
      run "ifconfig vlan2 create"
      run "ifconfig vlan2 10.0.1.1 netmask 255.255.255.0 vlan 2 vlandev em0"
      
   }

   if ($fw_carp_init) {
      run "sleep 2"
   }
   if (! $fw_carp_up)
      set-state fw_slave
}

state fw_slave {
   init {
      run "ifconfig vlan1 destroy"
      run "ifconfig vlan2 destroy"
      
      run "ifconfig em0 down"
   }

   if ($fw_carp_init)
      run "sleep 2"
   if ($fw_carp_up)
      set-state fw_master
}
```

Of course, you can also build a dependency-free binary and deploy it across
all of your server:

`$ go build`

License
---
MIT
