package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/niumandzi/nto2022/internal/usecase"
)

var menuItems = map[string][]string{
	"":         {"отели", "контакты"},
	"контакты": {"управляющие", "физические лица", "юридические лица"},
}

func NavigationBar(content *fyne.Container, usecase usecase.ContactUseCase) fyne.CanvasObject {
	tree := widget.NewTreeWithStrings(menuItems)
	tree.OnSelected = func(id string) {
		updateContent(content, id, usecase)
	}
	return container.NewBorder(nil, nil, nil, nil, tree)
}

func updateContent(content *fyne.Container, id string, usecase usecase.ContactUseCase) {
	content.Objects = nil // очистить текущий контент
	switch id {
	case "отели":
		content.Add(widget.NewLabel("Отображение информации об отелях"))
	case "контакты":
		SetContactScreen(content, usecase)

		// ... (и так далее для других случаев)
	}
	content.Refresh() // обновить контейнер
}
