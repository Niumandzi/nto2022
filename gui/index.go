package gui

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/niumandzi/nto2022/internal/usecase"
	"github.com/niumandzi/nto2022/pkg/components"
)

func Index(ctx context.Context, cases *usecase.UseCases) {
	a := app.New()
	w := a.NewWindow("НТО2022")
	w.SetMaster()

	content := container.NewStack()
	tutorial := container.NewBorder(nil, nil, nil, nil, content)

	split := container.NewHSplit(components.NavigationBar(content, cases.Contact), tutorial)
	//split := container.NewHSplit(widget.NewEntry(), widget.NewEntry())
	//fmt.Printf("%T", cases.Contact)
	split.Offset = 0
	w.SetContent(split)

	w.Resize(fyne.NewSize(1200, 700))
	w.ShowAndRun()
}
