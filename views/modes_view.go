package views

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	services "github.com/AdamElHassanLeb/popos-gpu-switcher/internal/services"
)

const SelectedLanguage string = "en"

type listItem struct {
	Name string
	Desc string
}

func ListPage(parent fyne.Window, service services.GpuModeService) fyne.CanvasObject {
	_ = service // use later

	// Main background
	bg := canvas.NewLinearGradient(
		color.RGBA{123, 31, 162, 255},
		color.RGBA{236, 64, 122, 255},
		math.Pi/4,
	)

	//HERE CHAT
	//Here call AvailableModes(ctx context.Context) (*Modes, error)
	//then for each mode get the description based on selected language
	//call async with loading animation while it gets

	items := []listItem{
		{"Integrated", "Use Intel iGPU (battery friendly)"},
		{"Hybrid", "Auto/offload to dGPU when needed"},
		{"NVIDIA", "Force dGPU (best performance)"},
	}

	// Root stack: base page + optional overlay (bubble)
	var overlay *fyne.Container
	root := container.NewStack()
	root.Add(bg)

	// Main list UI
	selected := 0
	var list *widget.List

	list = widget.NewList(
		func() int { return len(items) },
		func() fyne.CanvasObject { return NewModeListItem() }, // from mode_list_item.go
		func(i widget.ListItemID, o fyne.CanvasObject) {
			row := o.(*ModeListItem)
			row.NameLabel.SetText(items[i].Name)
			row.DescLabel.SetText(items[i].Desc)
			row.SetSelected(i == selected)

			row.OnTap = func() {
				selected = i
				list.Refresh()
				showBubble(root, parent, &overlay, items[i], func() {
					// TODO apply using service
					// service.SetMode(...)
				})
			}
		},
	)
	list.HideSeparators = true

	centeredList := container.NewCenter(
		container.NewGridWrap(fyne.NewSize(620, 340), list),
	)
	root.Add(centeredList)

	return root
}

// showBubble renders a bubble overlay directly on top of the page (no dialog, no window chrome)
func showBubble(root *fyne.Container, parent fyne.Window, overlayPtr **fyne.Container, item listItem, onApply func()) {
	// Remove existing overlay if any
	if *overlayPtr != nil {
		root.Remove(*overlayPtr)
		*overlayPtr = nil
	}

	// Optional dim backdrop (set alpha to 0 if you truly want NO dim)
	backdrop := canvas.NewRectangle(color.RGBA{0, 0, 0, 110})

	// Bubble glass card
	card := canvas.NewRectangle(color.RGBA{255, 255, 255, 30})
	card.CornerRadius = 22
	card.StrokeWidth = 1
	card.StrokeColor = color.RGBA{255, 255, 255, 60}

	// Bubble background (same theme)
	bubbleBg := canvas.NewLinearGradient(
		color.RGBA{123, 31, 162, 255},
		color.RGBA{236, 64, 122, 255},
		math.Pi/4,
	)

	title := widget.NewLabelWithStyle(item.Name, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	desc := widget.NewLabel(item.Desc)
	desc.Alignment = fyne.TextAlignCenter
	desc.Wrapping = fyne.TextWrapWord

	// Button cards (match list item aesthetic)
	btnCard := func(alpha uint8) *canvas.Rectangle {
		r := canvas.NewRectangle(color.RGBA{255, 255, 255, alpha})
		r.CornerRadius = 14
		return r
	}

	var overlay *fyne.Container

	closeOverlay := func() {
		if overlay != nil {
			root.Remove(overlay)
			root.Refresh()
			*overlayPtr = nil
		}
	}

	applyBtn := widget.NewButton("Apply", func() {
		if onApply != nil {
			onApply()
		}
		closeOverlay()
	})
	cancelBtn := widget.NewButton("Cancel", func() { closeOverlay() })

	apply := container.NewStack(btnCard(60), container.NewPadded(applyBtn))
	cancel := container.NewStack(btnCard(40), container.NewPadded(cancelBtn))

	buttonRow := container.NewHBox(
		layout.NewSpacer(),
		container.NewGridWrap(fyne.NewSize(140, 44), cancel),
		container.NewGridWrap(fyne.NewSize(140, 44), apply),
		layout.NewSpacer(),
	)

	bubbleContent := container.NewPadded(container.NewVBox(
		title,
		desc,
		layout.NewSpacer(),
		buttonRow,
	))

	// Half-screen bubble size
	s := parent.Canvas().Size()
	bubble := container.NewGridWrap(
		fyne.NewSize(s.Width*0.5, s.Height*0.5),
		container.NewStack(bubbleBg, card, bubbleContent),
	)

	overlay = container.NewStack(
		backdrop,
		container.NewCenter(bubble),
	)

	*overlayPtr = overlay
	root.Add(overlay)
	root.Refresh()
}
