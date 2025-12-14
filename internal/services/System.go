package service

import (
	"os/exec"

	CustomErrors "github.com/AdamElHassanLeb/popos-gpu-switcher/customerrors"
)

func isSystem76() bool {
	_, err := exec.LookPath("system76-power")
	return err == nil
}

func GetGPUService() (GpuModeService, error) {

	if isSystem76() {
		return &System76_GpuModeService{}, nil
	}

	return nil, CustomErrors.ErrUnsupportedOS
}
