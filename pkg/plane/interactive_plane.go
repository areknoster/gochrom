package plane

import (
    "fyne.io/fyne"
    "fyne.io/fyne/widget"
    "github.com/areknoster/gochrom/pkg/gochrom"
    "github.com/areknoster/gochrom/pkg/normde"
    "github.com/areknoster/gochrom/pkg/raster"
    "github.com/sirupsen/logrus"
)

type InteractivePlane struct {
    widget.BaseWidget
    renderer gochrom.Renderer
    size     fyne.Size
    mode gochrom.InteractiveMode
}

func NewInteractivePlane(renderer gochrom.Renderer, size fyne.Size) (*InteractivePlane, func(mode gochrom.InteractiveMode)) {
    ip := &InteractivePlane{
        renderer: renderer,
        size: size,
    }
    ip.ExtendBaseWidget(ip)
    return ip, func(mode gochrom.InteractiveMode) {
        ip.mode = mode
    }
}

var (
    _ fyne.Widget   = &InteractivePlane{}
    _ gochrom.Plane = &InteractivePlane{}
    _ fyne.Draggable = &InteractivePlane{}
    _ fyne.Tappable  = &InteractivePlane{}
)

func (p *InteractivePlane) Tapped(event *fyne.PointEvent) {
    normPt := normde.NormPoint2D(
        raster.Pixel{
            X: event.Position.X,
            Y: event.Position.Y,
        },
        p.size.Width, p.size.Height).
        InvertY()
    logrus.Debugf("Tapped: %v", normPt)
    p.mode.HandleClick(normPt)
}

func (p *InteractivePlane) Dragged(event *fyne.DragEvent) {
    start := normde.NormPoint2D(
        raster.Pixel{
            X: event.Position.X - event.DraggedX,
            Y: event.Position.Y - event.DraggedY,
        },
        p.size.Width,
        p.size.Height).
        InvertY()

    vec := normde.NormVector2D(
        raster.Pixel{X: event.DraggedX, Y: event.DraggedY},
        p.size.Width, p.size.Height).
        InvertY()

    logrus.Debugf("Drag: start: %v, vec: %v", start, vec)
    p.mode.HandleDrag(start, vec)
}

func (p *InteractivePlane) DragEnd() {
    logrus.Debugf("Drag finished")
    p.mode.HandleDragEnd()
}

func (p *InteractivePlane) CreateRenderer() fyne.WidgetRenderer {
    return NewPlaneRenderer(p.size, p.renderer.Render)

}