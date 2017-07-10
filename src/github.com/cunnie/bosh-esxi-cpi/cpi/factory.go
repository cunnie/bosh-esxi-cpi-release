package cpi

import (
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	boshuuid "github.com/cloudfoundry/bosh-utils/uuid"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

type Factory struct {
	fs      boshsys.FileSystem
	uuidGen boshuuid.Generator
	opts    FactoryOpts
	logger  boshlog.Logger
}
