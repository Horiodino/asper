package modify

import "github.com/horiodino/asper/internal/logger"

// TODO add use valid  argument for all methods

type ModifyOperations interface {
	ModifyFirewall(instanceid logger.Asperstring) *modifyclient
	ModifyNetwork(instanceid logger.Asperstring) *modifyclient
	ModifyDisk(instanceid logger.Asperstring) *modifyclient
	ModifySSH(instanceid logger.Asperstring) *modifyclient
}
