package cpi

import "github.com/cunnie/bosh-esxi-cpi/apiv1"

type FactoryOpts struct {
	ESXi  ESXiOpts
	Agent apiv1.AgentOptions
}

type ESXiOpts struct {
}
