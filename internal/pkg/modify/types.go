package modify

import "github.com/horiodino/asper/internal/logger"

// TODO add use valid  argument for all methods

type ModifyOperations interface {
	ModifyFirewall(instanceid logger.Asperstring)
	ModifyNetwork(instanceid logger.Asperstring)
	ModifyDisk(instanceid logger.Asperstring)
	ModifySSH(instanceid logger.Asperstring)
}
