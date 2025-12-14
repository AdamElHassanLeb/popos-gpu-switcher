# Pop!_OS GPU Switcher

A small desktop application for switching GPU modes on Pop!_OS / System76 systems.

The application provides a simple GUI wrapper around `system76-power` and is designed with a clean separation between system-specific logic, UI, and presentation data. It is written in Go and distributed as a single self-contained executable.

---

## Overview

This project allows switching between the following GPU modes (when supported by hardware):

- `integrated`
- `nvidia`
- `hybrid`
- `compute`

GPU mode changes require a reboot, which the application can trigger after a successful switch.

The design intentionally keeps system interaction isolated so that additional operating systems or vendors can be supported later by adding new service implementations without changing the UI or application flow.

---

## Supported Systems

- Pop!_OS
- System76 hardware
- Intel, AMD, and NVIDIA configurations

Available GPU modes depend on the detected hardware. Unsupported modes fail gracefully.

---

## Project Structure

```text
popos-gpu-switcher/
├── cmd/
│   └── popos-gpu-switcher/
│       └── main.go
├── internal/
│   └── service/
│       ├── interface.go
│       ├── system76.go
│       └── run.go
├── apperrors/
│   └── errors.go
├── assets/
│   ├── embed.go
│   └── messages/
│       └── messages.json
├── assetmanager/
│   └── messages.go
├── views/
├── build/
├── go.mod
├── go.sum
└── README.md
