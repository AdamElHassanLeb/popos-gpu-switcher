package main

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"

	"github.com/kbinani/screenshot"
)

func StartUI(ctx context.Context) {
	a := app.New()

	d, ok := a.Driver().(desktop.Driver)
	if !ok {
		// Fallback: normal window
		w := a.NewWindow("Pop!_OS GPU Switcher")
		w.SetContent(widget.NewLabel("Desktop driver not available."))
		w.ShowAndRun()
		return
	}

	w := d.CreateSplashWindow() // borderless
	w.SetPadded(false)
	w.SetContent(container.NewVBox(
		widget.NewLabel("GPU Mode Selector"),
		widget.NewButton("Close", func() { w.Close() }),
	))

	// Borderless windows have no OS controls, so you must provide your own close button.
	// Resize to 60% of screen without hardcoding:
	if screenshot.NumActiveDisplays() > 0 {
		bounds := screenshot.GetDisplayBounds(0) // primary display (usually)
		w.Resize(fyne.NewSize(
			float32(bounds.Dx())*0.6,
			float32(bounds.Dy())*0.6,
		))
	} else {
		// Fallback if screen size can't be determined
		w.Resize(fyne.NewSize(900, 540))
	}

	w.CenterOnScreen()
	w.ShowAndRun()
}
