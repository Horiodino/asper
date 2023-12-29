package vm

import "github.com/horiodino/asper/internal/logger"

// TODO add use valid  argument for all methods

type ModifyOperations interface {
	ModifyVM(instanceid logger.Asperstring) *modifyinstance
	MoveVM(instanceid logger.Asperstring) *modifyclient
	ResizeVM(instanceid logger.Asperstring) *modifyclient
	RebootVM(instanceid logger.Asperstring) *modifyclient
	StopVM(instanceid logger.Asperstring) *modifyclient
	StartVM(instanceid logger.Asperstring) *modifyclient
	RebuildVM(instanceid logger.Asperstring) *modifyclient
	ReinstallVM(instanceid logger.Asperstring) *modifyclient
	ResetVM(instanceid logger.Asperstring) *modifyclient
	CloneVM(instanceid logger.Asperstring) *modifyclient
	BackupVM(instanceid logger.Asperstring) *modifyclient
	RestoreVM(instanceid logger.Asperstring) *modifyclient
	ConvertVM(instanceid logger.Asperstring) *modifyclient
}
