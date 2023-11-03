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
		case "Главная":
			// content.Objects = []fyne.CanvasObject{...}
		case "Пользователи":
			// content.Objects = []fyne.CanvasObject{...}
		case "Пользователи/Пользователь 1":
			// content.Objects = []fyne.CanvasObject{...}
			// ... обрабатываем другие случаи
		}
		content.Refresh() // Обновляем содержимое контейнера
	}
}
