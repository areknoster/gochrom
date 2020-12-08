package plane

import (
    "fyne.io/fyne"
    "fyne.io/fyne/widget"
    "github.com/areknoster/gochrom/pkg/gochrom"
)

type Plane struct {
    widget.BaseWidget
    renderer gochrom.Renderer
    size     fyne.Size
}

func NewPlane(renderer gochrom.Renderer, size fyne.Size) *Plane {
    p := &Plane{
        size:     size,
        renderer: renderer,
    }
    p.ExtendBaseWidget(p)
    return p
}

var (
    _ fyne.Widget   = &Plane{}
    _ gochrom.Plane = &Plane{}
)


func (p *Plane) CreateRenderer() fyne.WidgetRenderer {
    return NewPlaneRenderer(p.size, p.renderer.Render)
}
