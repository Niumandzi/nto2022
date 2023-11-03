package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NavigationBar() fyne.CanvasObject {
	// Дерево навигации
	treeData := map[string][]string{
		"":         {"отели", "контакты"},
		"контакты": {"управляющие", "физические лица", "юридические лица"},
	}

	navTree := widget.NewTreeWithStrings(treeData)
	navTree.OnSelected = func(id string) {
		// Обработка выбранного элемента
		// Например, переключение контента на основе выбранного раздела
	}

	return container.NewVScroll(navTree)
}
