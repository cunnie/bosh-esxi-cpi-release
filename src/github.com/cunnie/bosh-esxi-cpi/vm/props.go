package vm

import "github.com/vmware/govmomi/vim25/types"

//import "github.com/vmware/govmomi"

type VMProps struct {
	ExposedPorts []string `json:"ports"` // [6868/tcp]

	types.VirtualMachineConfigSummary
}

type NetProps struct {
	Name   string
	Driver string
}
