package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewCard(title string) *fyne.Container {
	return container.NewVBox(
		widget.NewLabel(title),
		widget.NewButton("Подробнее", func() {
			// Обработка нажатия кнопки
		}),
	)
}
