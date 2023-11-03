package components

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/niumandzi/nto2022/internal/usecase"
	"github.com/niumandzi/nto2022/model"
)

func CreateContact(content *fyne.Container, c usecase.ContactUseCase) *fyne.Container {
	label := widget.NewLabel("Создание контакта")

	contactType := widget.NewSelect([]string{"Физическое лицо", "Юридическое лицо", "Управляющий"}, func(s string) {})

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
		case "Физическое лицо":
			contactTypeData = "private_client"
		case "Юридическое лицо":
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
		_, err := c.CreateContact(contactData)
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
