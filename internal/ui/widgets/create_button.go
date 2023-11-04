package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateButtonWidget(id string, window fyne.Window) *widget.Button {
	button := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
	})

	return button
}
