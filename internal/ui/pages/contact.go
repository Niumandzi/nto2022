package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
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
			card := createContactCard(cases, contact, window)
			grid.Add(card)
		}

		contactListContainer.Objects = []fyne.CanvasObject{container.NewVScroll(grid)}
		contactListContainer.Refresh() // Обновляем контейнер с контактами
	}

	createButton := widgets.CreateButtonWidget("", window)
	// Создаем выпадающее меню и передаем в него функцию updateContactList как callback
	typeSelect := widgets.TypeSelectWidget(contactTypes, updateContactList)

	// Создаем контейнер для выпадающего меню, чтобы разместить его в верхней части экрана
	toolbar := container.NewBorder(nil, nil, typeSelect, createButton)

	// Начальное обновление списка контактов
	updateContactList("all")

	// Создаем и возвращаем главный контейнер
	mainContainer := container.NewBorder(toolbar, nil, nil, nil, contactListContainer)

	return mainContainer
}

func createContactCard(cases *usecase.UseCases, contact model.Contact, window fyne.Window) fyne.CanvasObject {
	cardText := card(contact)
	label := widget.NewLabel(cardText)
	label.Wrapping = fyne.TextWrapWord // Text wrapping

	// Create buttons using standard icons
	deleteButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		// Your code to handle deleting using contact.ID
	})
	editButton := widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), func() {
		// Your code to handle editing using contact.ID
	})

	// Make buttons small
	deleteButton.Importance = widget.LowImportance
	editButton.Importance = widget.LowImportance

	// Arrange buttons in the bottom right corner
	buttons := container.NewHBox(layout.NewSpacer(), editButton, deleteButton)

	// Create card with label and buttons
	card := widget.NewCard("", "", container.NewBorder(nil, buttons, nil, nil, label))

	return card
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
