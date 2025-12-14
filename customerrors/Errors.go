package errors

import "errors"

var ErrUnsupportedOS = errors.New("unsupported operating system")
var ErrInvalidGPUMode = errors.New("invalid gpu mode")
var ErrModeSwitchUnsuccesful = errors.New("could not switch gpu mode")
var ErrReboot = errors.New("could not reboot")
var ErrShutdown = errors.New("could not shutdown")
var ErrCurrentMode = errors.New("could not get current mode")
