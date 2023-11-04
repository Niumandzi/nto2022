package contact

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/niumandzi/nto2022/internal/ui/widgets"
	"github.com/niumandzi/nto2022/internal/usecase"
	"github.com/niumandzi/nto2022/model"
)

func CreateContact(cases *usecase.UseCases, window fyne.Window) {
	var (
		contactType string
		name        string
		number      string
		email       string
	)

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Имя")

	numberEntry := widget.NewEntry()
	numberEntry.SetPlaceHolder("Номер телефона")

	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Email")

	typeSelect := widgets.TypeSelectWidget(map[string]string{
		"Менеджер":         "worker",
		"Физическиое лицо": "private_client",
		"Юридическое лицо": "legal_client",
	}, func(selected string) {
		contactType = selected
	})

	// Форма для ввода информации о контакте
	form := []*widget.FormItem{
		widget.NewFormItem("", typeSelect),
		widget.NewFormItem("", nameEntry),
		widget.NewFormItem("", numberEntry),
		widget.NewFormItem("", emailEntry),
	}

	// Создание диалогового окна
	dialog.ShowForm("Создать контакт", "Создать", "Отмена", form, func(b bool) {
		if !b {
			return // Пользователь отменил операцию
		}
		// Сбор информации из формы
		name = nameEntry.Text
		number = numberEntry.Text
		email = emailEntry.Text

		// Создание контакта
		newContact := model.Contact{
			ContactType: contactType,
			Name:        name,
			Number:      number,
			Email:       email,
		}

		// Передача контакта в usecase
		_, err := cases.Contact.CreateContact(newContact)
		if err != nil {
			dialog.ShowError(err, window)
		} else {
			dialog.ShowInformation("Контакт создан", "Контакт успешно создан!", window)
		}
	}, window)
}
