package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/niumandzi/nto2022/internal/ui/widgets"
	"github.com/niumandzi/nto2022/internal/usecase"
	"github.com/niumandzi/nto2022/model"
)

func Contacts(cases *usecase.UseCases, window fyne.Window) fyne.CanvasObject {
	contactListContainer := container.NewStack()

	contactTypes := map[string]string{
		"Все":              "all",
		"Менеджеры":        "worker",
		"Физические лица":  "private_client",
		"Юридические лица": "legal_client",
	}

	updateContactList := func(contactType string) {
		contacts, err := cases.Contact.GetContactsByType(contactType)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		grid := container.New(layout.NewGridLayoutWithColumns(3))
		for _, contact := range contacts {
			card := createContactCard(contact)
			grid.Add(card)
		}

		contactListContainer.Objects = []fyne.CanvasObject{container.NewVScroll(grid)}
		contactListContainer.Refresh() // Обновляем контейнер с контактами
	}

	// Создаем выпадающее меню и передаем в него функцию updateContactList как callback
	contactTypeSelect := widgets.TypeSelectWidget(contactTypes, updateContactList)

	// Создаем контейнер для выпадающего меню, чтобы разместить его в верхней части экрана
	toolbar := container.NewBorder(nil, nil, nil, contactTypeSelect)

	// Начальное обновление списка контактов
	updateContactList("all")

	// Создаем и возвращаем главный контейнер
	mainContainer := container.NewBorder(toolbar, nil, nil, nil, contactListContainer)

	return mainContainer
}

func createContactCard(contact model.Contact) fyne.CanvasObject {
	label := widget.NewLabel(card(contact))
	label.Wrapping = fyne.TextWrapWord // Перенос текста по словам
	return widget.NewCard("", "", label)
}

func card(contact model.Contact) string {
	var status string
	switch contact.ContactType {
	case "worker":
		status = "Менеджер"
	case "private_client":
		status = "Физическое лицо"
	case "legal_client":
		status = "Юридическое лицо"
	}

	return "Статус: " + status + "\n" +
		"Имя: " + contact.Name + "\n" +
		"Номер: " + contact.Number + "\n" +
		"Email: " + contact.Email
}
