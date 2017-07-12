package cpi

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/vmware/govmomi"
)

type FactoryOpts struct {
	ESXiClient *govmomi.Client
	Agent      apiv1.AgentOptions
}
