package main

var templateText string = `# This file is generated automatically. DO NOT EDIT.
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
      run "ifconfig {{.Device}} up"{{$dev := .Device}}
      {{range .Vlans}}
      run "ifconfig vlan{{.Id}} create"
      run "ifconfig vlan{{.Id}} {{.Address}} netmask {{.Netmask}} vlan {{.Id}} vlandev {{$dev}}"
      {{end}}
   }

   if ($fw_carp_init) {
      run "sleep 2"
   }
   if (! $fw_carp_up)
      set-state fw_slave
}

state fw_slave {
   init {
      {{range .Vlans}}run "ifconfig vlan{{.Id}} destroy"
      {{end}}
      run "ifconfig {{.Device}} down"
   }

   if ($fw_carp_init)
      run "sleep 2"
   if ($fw_carp_up)
      set-state fw_master
}
`
