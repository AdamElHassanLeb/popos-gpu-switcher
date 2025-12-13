package service

import (
	"context"

	CustomErrors "github.com/AdamElHassanLeb/popos-gpu-switcher/internal/CustomErrors"
)

type System76_GpuModeService struct{}

func (s *System76_GpuModeService) AvailableModes(ctx context.Context) (*Modes, error) {

	modes := Modes{
		modeMap: map[string]Mode{
			"integrated": {
				Name:        "integrated",
				Description: "Use only the integrated GPU for lowest power usage",
			},
			"nvidia": {
				Name:        "nvidia",
				Description: "Use the NVIDIA GPU for maximum performance",
			},
			"hybrid": {
				Name:        "hybrid",
				Description: "Use integrated GPU with NVIDIA on-demand",
			},
			"compute": {
				Name:        "compute",
				Description: "Use NVIDIA GPU only for compute workloads",
			},
		},
	}

	return &modes, nil
}

func (s *System76_GpuModeService) CurrentMode(ctx context.Context) (string, error) {
	res, _ := run("system76-power", "graphics")
	return res, CustomErrors.ErrCurrentMode
}

func (s *System76_GpuModeService) SwitchMode(ctx context.Context, mode string) error {

	run("pkexec", "system76-power", "graphics", mode)

	return CustomErrors.ErrModeSwitchUnsuccesful
}

func (s *System76_GpuModeService) Reboot(ctx context.Context) error {
	run("pkexec", "reboot")
	return CustomErrors.ErrReboot
}
