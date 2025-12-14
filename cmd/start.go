package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	services "github.com/AdamElHassanLeb/popos-gpu-switcher/internal/services"
	"github.com/AdamElHassanLeb/popos-gpu-switcher/views"
)

func StartUI(service services.GpuModeService, serviceStarted bool) {
	a := app.New()

	if !serviceStarted {
		w := a.NewWindow("Pop!_OS GPU Switcher")
		w.SetContent(widget.NewLabel("Service couldnt start"))
		w.ShowAndRun()
		return
	}

	d, ok := a.Driver().(desktop.Driver)
	if !ok {
		w := a.NewWindow("Pop!_OS GPU Switcher")
		w.SetContent(widget.NewLabel("Desktop driver not available."))
		w.ShowAndRun()
		return
	}

	// Borderless window
	w := d.CreateSplashWindow()
	w.SetPadded(false)

	w.SetContent(
		container.NewBorder(
			widget.NewLabel("GPU Mode Switcher"),            //top
			widget.NewButton("Close", func() { w.Close() }), // bottom
			nil, // left
			nil, // right
			views.ListPage(w, service),
		),
	)

	// Hardcoded size (adjust later)
	w.Resize(fyne.NewSize(900, 540))
	w.CenterOnScreen()

	w.ShowAndRun()
}
