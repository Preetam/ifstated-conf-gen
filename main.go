package main

import (
	"encoding/json"
	"os"
	"text/template"
)

type Vlan struct {
	Id      int
	Address string
	Netmask string
}

type VlanDev struct {
	Device string
	Vlans  []Vlan
}

var config VlanDev

func main() {
	d := json.NewDecoder(os.Stdin)
	d.Decode(&config)
	t, _ := template.Must(template.New("config"), nil).Parse(templateText)
	t.Execute(os.Stdout, config)
}
