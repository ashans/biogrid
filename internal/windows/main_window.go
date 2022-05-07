package windows

import (
	"biogrid/internal/algo"
	"biogrid/internal/components/form"
	"biogrid/internal/components/gridview"
	"biogrid/internal/entities"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func NewWindow(a fyne.App, config entities.AlignConfig) fyne.Window {
	mainWindow := a.NewWindow("BioGrid")
	mainWindow.Resize(fyne.NewSize(600, 400))
	p := layout.NewPaddedLayout()

	center := container.NewCenter()

	aligner := algo.NewAligner()

	f := form.NewForm(config, func(c entities.AlignConfig) {
		grid := gridview.NewGridView(c, aligner.Align(c))

		if len(center.Objects) != 0 {
			center.Objects = []fyne.CanvasObject{}
		}
		center.Add(grid)
	})

	mainWindow.SetContent(container.NewVBox(
		container.New(p, f),
		center,
	))

	return mainWindow
}
