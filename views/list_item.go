package views

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ModeListItem struct {
	widget.BaseWidget

	NameLabel *widget.Label
	DescLabel *widget.Label

	card     *canvas.Rectangle
	selected bool

	OnTap func()
}

func NewModeListItem() *ModeListItem {
	m := &ModeListItem{
		NameLabel: widget.NewLabel(""),
		DescLabel: widget.NewLabel(""),
	}

	m.NameLabel.TextStyle = fyne.TextStyle{Bold: true}
	m.DescLabel.Wrapping = fyne.TextWrapWord

	m.ExtendBaseWidget(m)
	return m
}

func (m *ModeListItem) SetSelected(v bool) {
	m.selected = v
	m.applyStyle()
}

func (m *ModeListItem) Tapped(*fyne.PointEvent) {
	if m.OnTap != nil {
		m.OnTap()
	}
}

func (m *ModeListItem) CreateRenderer() fyne.WidgetRenderer {
	m.card = canvas.NewRectangle(color.RGBA{255, 255, 255, 28})
	m.card.CornerRadius = 18

	row := container.NewVBox(m.NameLabel, m.DescLabel)
	root := container.NewMax(
		m.card,
		container.NewPadded(row),
	)

	m.applyStyle()
	return widget.NewSimpleRenderer(root)
}

func (m *ModeListItem) applyStyle() {
	if m.card == nil {
		return
	}

	// unselected: subtle glass
	if !m.selected {
		m.card.FillColor = color.RGBA{255, 255, 255, 28}
		m.card.Refresh()
		return
	}

	// selected: brighter highlight (still glassy)
	m.card.FillColor = color.RGBA{255, 255, 255, 65}
	m.card.Refresh()
}
