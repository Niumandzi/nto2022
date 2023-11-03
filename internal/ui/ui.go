package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func SetupNavigation(nav *widget.Tree, content *fyne.Container) {
	nav.OnSelected = func(id string) {
		// Обработка выбранного элемента
		// Обновляем основной контент в зависимости от выбранного элемента
		switch id {
		case "отели":
			// content.Objects = []fyne.CanvasObject{...}
		case "контакты":
			// content.Objects = []fyne.CanvasObject{...}
		}
		content.Refresh() // Обновляем содержимое контейнера
	}
}
