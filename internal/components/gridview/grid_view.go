package gridview

import (
	"biogrid/internal/entities"
	"biogrid/internal/util/gutil"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewGridView(config entities.AlignConfig, grid gutil.Grid) *fyne.Container {
	values := grid.Flatten()
	var objs []fyne.CanvasObject
	for _, v := range values {
		objs = append(objs, widget.NewLabel(v))
	}

	return container.NewGridWithColumns(len(config.Seq1)+2, objs...)
}
