package main

import (
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/cunnie/bosh-esxi-cpi/cpi"
)

type Config struct {
	Actions cpi.FactoryOpts
}

func NewConfigFromPath(path string, fs boshsys.FileSystem) (config Config, err error) {
	config = Config{}
	err = nil
	return config, nil
}
