package cpi

import "github.com/cppforlife/bosh-cpi-go/apiv1"

type FactoryOpts struct {
	ESXi  ESXiOpts
	Agent apiv1.AgentOptions
}

type ESXiOpts struct {
}
