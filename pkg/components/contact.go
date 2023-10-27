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
	})

	leftSide := container.NewVBox(
		label,
		contactType,
		contactName,
		contactPhoneNumber,
		contactEmail,
		createBtn,
	)
	content.Add(leftSide)
}
