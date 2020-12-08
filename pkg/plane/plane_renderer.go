package plane

import (
	"image"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"github.com/sirupsen/logrus"
)

type PlaneRenderer struct {
	raster *canvas.Raster
	render func(w int, h int) image.Image
	size fyne.Size
}

func NewPlaneRenderer(size fyne.Size, render func(w int, h int) image.Image) *PlaneRenderer {
	pr := &PlaneRenderer{
		render: render,
		size: size,
	}
	pr.raster = canvas.NewRaster(render)
	return pr
}

func (pr *PlaneRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (pr *PlaneRenderer) Destroy() {
}

func (pr *PlaneRenderer) Layout(size fyne.Size) {
	pr.raster.Resize(size)
}

func (pr *PlaneRenderer) MinSize() fyne.Size {
	return pr.size
}

func (pr *PlaneRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{pr.raster}
}

func (pr *PlaneRenderer) Refresh() {
	logrus.Debug("plane: PlaneRenderer is refreshing")
	canvas.Refresh(pr.raster)
}

var _ fyne.WidgetRenderer = &PlaneRenderer{}

