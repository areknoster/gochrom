package render

import (
    "fyne.io/fyne"
    "github.com/areknoster/gochrom/pkg/gochrom"
    "github.com/areknoster/gochrom/pkg/normde"
    "github.com/areknoster/gochrom/pkg/spectrum"
    "github.com/fogleman/gg"
    "golang.org/x/image/colornames"
    "image"
)

type CIE struct {
    storage gochrom.StateStorage
    diagImg image.Image
}

func NewCIE(storage gochrom.StateStorage) *CIE {
    img, err := gg.LoadPNG("resources/CIE-diagram.png")
    if err != nil {
        panic(err)
    }
    return &CIE{
        storage: storage,
        diagImg: img,
    }
}

var _ gochrom.Renderer = &CIE{}

func (c *CIE) Render(w, h int) image.Image {
    dc := gg.NewContext(w, h)

    //set background color and image
    dc.SetColor(colornames.White)
    dc.Clear()
    resized := normde.ResizeImage(c.diagImg, fyne.Size{w, h})
    dc.DrawImage(resized, 0, 0)

    //add underlying CIE diagram shape
    dc.InvertY()
    dc.Scale(float64(w) * 1.03, float64(h) * 0.97)
    dc.Translate(0.115, 0.1)
    for _, pt := range spectrum.MatchingDataPoints {
       cp := pt.ToXYZ().Normalize()
       dc.LineTo(cp.X, cp.Y)
    }
    dc.ClosePath()
    dc.SetColor(colornames.Firebrick)
    dc.SetLineWidth(5)
    dc.Stroke()

    //add point the spectrum aims at
    spectrumPoint := c.storage.Get().Spectrum.ToXYZ().Normalize()
    _ = spectrumPoint
    dc.DrawCircle(spectrumPoint.X, spectrumPoint.Y, 0.01)
    dc.SetColor(colornames.Black)
    dc.Stroke()

    //err := dc.LoadFontFace("resources/Roboto-Regular.ttf", 50)
    //if err != nil{
    //    panic(err)
    //}
    //dc.DrawStringAnchored("hello", 250, 250, 0.5, 0.5)

    //dc.SavePNG("image.png")
    return dc.Image()
}
