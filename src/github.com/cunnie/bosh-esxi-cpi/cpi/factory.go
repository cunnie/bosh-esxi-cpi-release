package cpi

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	boshuuid "github.com/cloudfoundry/bosh-utils/uuid"
)

type Factory struct {
	fs      boshsys.FileSystem
	uuidGen boshuuid.Generator
	opts    FactoryOpts
	logger  boshlog.Logger
}
