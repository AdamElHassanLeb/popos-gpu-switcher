package service

import (
	"context"
	"sync"
)

type GpuModeService interface {
	// AvailableModes returns the supported modes for this machine.
	// Typical values: integrated, nvidia, hybrid, compute.
	AvailableModes(ctx context.Context) (*Modes, error)

	// CurrentMode returns the currently selected graphics mode.
	CurrentMode(ctx context.Context) (string, error)

	// SwitchMode switches to the requested mode. Requires privilege.
	// Caller handles reboot prompting.
	SwitchMode(ctx context.Context, mode string) error

	//Reboot, reboots machine on demand
	Reboot(ctx context.Context) error
}

type Mode struct {
	Name string
}

type Modes struct {
	mu      sync.Mutex
	modeMap map[string]Mode
}
