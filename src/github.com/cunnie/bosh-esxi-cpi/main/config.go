package main

import (
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/cppforlife/bosh-docker-cpi-release/src/github.com/cppforlife/bosh-docker-cpi/cpi"
)

type Config struct {
	Actions cpi.FactoryOpts
}

func NewConfigFromPath(path string, fs boshsys.FileSystem) (config Config, err error) {
	config = Config{}
	err = nil
	return config, nil
}
