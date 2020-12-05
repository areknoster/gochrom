package plane

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"github.com/sirupsen/logrus"
)

type PlaneRenderer struct {
	raster *canvas.Raster
	plane  *Plane
}

func NewPlaneRenderer(plane *Plane) *PlaneRenderer {
	pr := &PlaneRenderer{plane: plane}
	pr.raster = canvas.NewRaster(plane.renderer.Render)
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
	return pr.plane.Size()
}

func (pr *PlaneRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{pr.raster}
}

func (pr PlaneRenderer) Refresh() {
	logrus.Debug("plane: PlaneRenderer is refreshing")
	canvas.Refresh(pr.raster)
}

var _ fyne.WidgetRenderer = &PlaneRenderer{}

