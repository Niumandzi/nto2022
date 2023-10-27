package components

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/niumandzi/nto2022/internal/usecase"
	"github.com/niumandzi/nto2022/model"
)

func SetContactScreen(content *fyne.Container, c usecase.ContactUseCase) {
	leftSide := CreateContact(content, c)
	rightSide := AllContacts(content, c)

	screen := container.NewHBox(leftSide, rightSide)

	content.Add(screen)
}

func ContactCard(contact model.Contact) string {
	var contactTypeData string

	switch contact.ContactType {
	case "private_client":
		contactTypeData = "Приватное лицо"
	case "legal_client":
		contactTypeData = "Легальное лицо"
	case "worker":
		contactTypeData = "Управляющий"
	}

	contactCard := contactTypeData + "\n" + contact.Name + "\n" + contact.Number

	return contactCard
}

func AllContacts(content *fyne.Container, c usecase.ContactUseCase) *widget.List {
	ctx := context.Background()
	contacts, err := c.GetAllContacts(ctx)
	if err != nil {
		NewModalWindow(content, err.Error())
	}
	rightSide := widget.NewList(
		func() int {
			return len(contacts)
		},
		func() fyne.CanvasObject {
			wdg := widget.NewLabel("")
			wdg.Resize(fyne.NewSize(500, 500))
			return wdg // представление контакта
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			label.Resize(fyne.NewSize(500, 500))
			contact := contacts[id]
			label.SetText(ContactCard(contact))
		},
	)
	rightSide.Resize(fyne.NewSize(500, 500))
	return rightSide
}

func CreateContact(content *fyne.Container, c usecase.ContactUseCase) *fyne.Container {
	label := widget.NewLabel("Создание контакта")

	contactType := widget.NewSelect([]string{"Приватное лицо", "Легальное лицо", "Управляющий"}, func(s string) {})

	contactName := widget.NewEntry()
	contactName.SetPlaceHolder("Имя")

	contactPhoneNumber := widget.NewEntry()
	contactPhoneNumber.SetPlaceHolder("Номер телефона")

	contactEmail := widget.NewEntry()
	contactEmail.SetPlaceHolder("Email")

	createBtn := widget.NewButton("Создать", func() {
		ctx := context.Background()

		var contactTypeData string

		switch contactType.Selected {
		case "Приватное лицо":
			contactTypeData = "private_client"
		case "Легальное лицо":
			contactTypeData = "legal_client"
		case "Управляющий":
			contactTypeData = "worker"
		}
		contactData := model.Contact{
			ContactType: contactTypeData,
			Name:        contactName.Text,
			Number:      contactPhoneNumber.Text,
			Email:       contactEmail.Text,
		}
		_, err := c.CreateContact(ctx, contactData)
		if err != nil {
			NewModalWindow(content, err.Error())
		}
		SetContactScreen(content, c)
	})

	leftSide := container.NewVBox(
		label,
		contactType,
		contactName,
		contactPhoneNumber,
		contactEmail,
		createBtn,
	)

	return leftSide
}
