package render

import (
    "fyne.io/fyne"
    "github.com/areknoster/gochrom/pkg/geom2d"
    "github.com/areknoster/gochrom/pkg/gochrom"
    "github.com/areknoster/gochrom/pkg/normde"
    "github.com/areknoster/gochrom/pkg/spectrum"
    "github.com/fogleman/gg"
    "golang.org/x/image/colornames"
    "image"
)

type SPD struct{
    storage gochrom.StateStorage
    background image.Image
}

func NewSPD(storage gochrom.StateStorage) *SPD {
    img, err := gg.LoadPNG("resources/spectrum.png")
    if err != nil {
        panic(err)
    }
    return &SPD{
        storage: storage,
        background: img,
    }
}

var _ gochrom.Renderer = &SPD{}

var(
    GraphMin = geom2d.Point{0.1883, 0.1700}
    GraphMax = geom2d.Point{0.8550, 0.5466}
)

func (s *SPD) Render(w, h int) image.Image {
    dc := gg.NewContext(w,h)
    resized := normde.ResizeImage(s.background, fyne.Size{w, h})
    dc.SetColor(colornames.White)
    dc.Clear()
    dc.DrawImage(resized, 0, 0)
    dc.InvertY()
    dc.Scale(float64(w), float64(h))
    //draw spectrum points
    state := s.storage.Get()
    lambdas := state.Spectrum.Lambdas
    xRange := GraphMax.X - GraphMin.X
    yRange := GraphMax.Y - GraphMin.Y

    for i, lambda := range lambdas {
        dc.LineTo(GraphMin.X + float64(i) * xRange / float64(spectrum.SpectrumRange), GraphMin.Y + lambda * yRange)
    }
    dc.SetColor(colornames.Black)
    dc.SetLineWidth(10)
    dc.Stroke()

    return dc.Image()
}

