package widget

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ModalWindowWidget(content fyne.CanvasObject, s string) {
	var popUp *widget.PopUp
	mainWin := fyne.CurrentApp().Driver().CanvasForObject(content)

	modalContent := container.NewVBox(
		widget.NewLabel(fmt.Sprintf("Attention! %v", s)),
		widget.NewButton("Close", func() {
			popUp.Hide()
		}))
	popUp = widget.NewModalPopUp(modalContent, mainWin)
	popUp.Show()
}
