package service

import "context"

type GpuModeService interface {
	// AvailableModes returns the supported mode strings for this machine.
	// Typical values: integrated, nvidia, hybrid, compute.
	AvailableModes(ctx context.Context) ([]string, error)

	// CurrentMode returns the currently selected graphics mode.
	CurrentMode(ctx context.Context) (string, error)

	// SwitchMode switches to the requested mode. Requires privilege.
	// Caller handles reboot prompting.
	SwitchMode(ctx context.Context, mode string) error

	//Reboot, reboots machine on demand
	Reboot(ctx context.Context) error
}
