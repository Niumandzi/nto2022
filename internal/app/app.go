package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/niumandzi/nto2022/internal/ui/widgets"
)

type GUI struct {
	App    fyne.App
	Window fyne.Window
}

func NewGUI() *GUI {
	// Инициализация нового приложения Fyne
	a := app.New()

	// Создание главного окна приложения
	w := a.NewWindow("НТО 2022")
	w.Resize(fyne.NewSize(800, 600)) // Установка стандартных размеров окна

	// Здесь могут быть вызовы функций, которые настраивают глобальные ресурсы,
	// такие как база данных, конфигурации и т.д.
	// setupGlobalResources()

	return &GUI{
		App:    a,
		Window: w,
	}
}
func (gui *GUI) SetupUI() {
	// Создаем навигационную панель
	navBar := widgets.NavigationBar()

	// Создаем контейнер для основного содержимого (заполнитель)
	content := container.NewVBox(
	// ... Ваш основной контент
	)

	// Объединяем навигационную панель и основное содержимое
	splitContainer := container.NewHSplit(navBar, content)
	splitContainer.Offset = 0.2 // Отношение ширины панели к общей ширине окна

	// Устанавливаем объединенный контейнер в качестве содержимого окна
	gui.Window.SetContent(splitContainer)
}

func (gui *GUI) Run() {
	// Устанавливаем содержимое окна
	gui.SetupUI()

	// Показываем окно и запускаем приложение
	gui.Window.ShowAndRun()
}
