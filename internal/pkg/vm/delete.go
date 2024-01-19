package vm

import (
	"github.com/horiodino/asper/internal/logger"
)

// takes slice
func (c *vm) DeleteVM(instanceids []string) (TerminationResult, error) {

	response, err := validateInput(instanceids)
	if err != nil {
		return TerminationResult{}, err
	}
	logger := logger.NewLogger()
	logger.LogInfo(response, input)

	return TerminationResult{}, nil
}
