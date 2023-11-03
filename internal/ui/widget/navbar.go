package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/niumandzi/nto2022/internal/ui/page"
	"github.com/niumandzi/nto2022/internal/usecase"
)

func NavigationBarWidget(mainContent *fyne.Container, cases *usecase.UseCases) *widget.Tree {
	// Дерево навигации
	treeData := map[string][]string{
		"": {"отели", "контакты"}}

	navTree := widget.NewTreeWithStrings(treeData)
	navTree.OnSelected = func(id string) {
		var content fyne.CanvasObject
		// Обработка выбранного элемента
		switch id {
		case "отели":
			//content = ui.ShowHotels() // Предположим, что ui.ShowHotels возвращает fyne.CanvasObject
		case "контакты":
			content = page.ShowContacts(cases)
		default:
			content = widget.NewLabel("Выберите категорию")
		}

		mainContent.Objects = []fyne.CanvasObject{content}
		mainContent.Refresh() // Обновляем содержимое контейнера
	}

	return navTree
}
