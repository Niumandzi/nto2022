package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/niumandzi/nto2022/internal/usecase"
	"github.com/niumandzi/nto2022/model"
)

func ShowContacts(cases *usecase.UseCases) fyne.CanvasObject {
	contacts, err := cases.Contact.GetAllContacts()
	if err != nil {
		// обработка ошибки
		return widget.NewLabel("Ошибка загрузки контактов")
	}

	// Установите размер карточки контакта

	grid := container.New(layout.NewGridLayoutWithColumns(3))

	for _, contact := range contacts {
		card := createContactCard(contact)
		grid.Add(card)
	}

	return container.NewVScroll(grid)
}

// createContactCard создает виджет для отображения информации о контакте
func createContactCard(contact model.Contact) fyne.CanvasObject {
	label := widget.NewLabel(сontactCard(contact))
	label.Wrapping = fyne.TextWrapWord // Перенос текста по словам
	return widget.NewCard("", "", label)
}

func сontactCard(contact model.Contact) string {
	return "Статус: " + contact.ContactType + "\n" +
		"Имя: " + contact.Name + "\n" +
		"Номер: " + contact.Number + "\n" +
		"Email: " + contact.Email // Если вы хотите добавить Email
}
