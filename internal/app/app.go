package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/niumandzi/nto2022/internal/ui/widget"
	"github.com/niumandzi/nto2022/internal/usecase"
)

type GUI struct {
	App    fyne.App
	Window fyne.Window
	Cases  *usecase.UseCases
}

func NewGUI(cases *usecase.UseCases) *GUI {
	// Инициализация нового приложения Fyne
	a := app.New()

	// Создание главного окна приложения
	w := a.NewWindow("НТО 2022")
	w.Resize(fyne.NewSize(1200, 700)) // Установка стандартных размеров окна

	// Здесь могут быть вызовы функций, которые настраивают глобальные ресурсы,
	// такие как база данных, конфигурации и т.д.
	// setupGlobalResources()

	return &GUI{
		App:    a,
		Window: w,
		Cases:  cases,
	}
}
func (gui *GUI) SetupUI() {
	// Создаем пустой контентный контейнер
	mainContent := container.NewMax()

	// Создаем навигационную панель с функцией переключения контента
	navBar := widget.NavigationBarWidget(mainContent, gui.Cases)

	// Объединяем навигационную панель и основное содержимое
	splitContainer := container.NewHSplit(container.NewVScroll(navBar), mainContent)
	splitContainer.Offset = 0.15 // Задаем соотношение размеров панели и содержимого

	// Устанавливаем объединенный контейнер в качестве содержимого окна
	gui.Window.SetContent(splitContainer)
}

func (gui *GUI) Run() {
	// Устанавливаем содержимое окна
	gui.SetupUI()
	// Показываем окно и запускаем приложение
	gui.Window.ShowAndRun()
}
