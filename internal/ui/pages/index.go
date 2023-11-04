package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/niumandzi/nto2022/internal/ui/pages/contact"
	"github.com/niumandzi/nto2022/internal/usecase"
)

func NavigationBar(mainContent *fyne.Container, cases *usecase.UseCases, window fyne.Window) *widget.Tree {
	// Дерево навигации
	treeData := map[string][]string{
		"": {"отели", "контакты"}}

	navTree := widget.NewTreeWithStrings(treeData)
	navTree.OnSelected = func(id string) {
		var content fyne.CanvasObject
		// Обработка выбранного элемента
		switch id {
		case "отели":
			//content = page.ShowHotels()
		case "контакты":
			content = contact.ShowContacts(cases, window)
		default:
			content = widget.NewLabel("Выберите категорию")
		}

		mainContent.Objects = []fyne.CanvasObject{content}
		mainContent.Refresh() // Обновляем содержимое контейнера
	}

	return navTree
}
