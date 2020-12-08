package render

import (
    "fyne.io/fyne"
    "github.com/areknoster/gochrom/pkg/gochrom"
    "github.com/areknoster/gochrom/pkg/normde"
    "github.com/areknoster/gochrom/pkg/spectrum"
    "github.com/fogleman/gg"
    "golang.org/x/image/colornames"
    "golang.org/x/image/draw"
    "image"
)

type CIE struct {
    storage       gochrom.StateStorage
    isFramed      bool
    diagImg       image.Image
    framedDiagImg image.Image
}

func NewCIE(storage gochrom.StateStorage) *CIE {
    img, err := gg.LoadPNG("resources/CIE-diagram.png")
    //render framed image
    copy := image.NewRGBA(img.Bounds())
    draw.Draw(copy, copy.Bounds(), img, img.Bounds().Min, draw.Src)
    if err != nil {
        panic(err)
    }
    return &CIE{
        storage:       storage,
        diagImg:       img,
        framedDiagImg: createFramedImage(copy),
    }
}

func createFramedImage(img image.Image) image.Image {
    dc := gg.NewContext(img.Bounds().Max.X, img.Bounds().Max.Y)
    dc.Scale(float64(img.Bounds().Max.X), float64(img.Bounds().Max.Y))
    const thick = 0.05

    var x, y, w, h float64 = 0, 0, 1, 1

    //black
    dc.DrawRectangle(x, y, w, h)
    dc.SetColor(colornames.Black)
    dc.Fill()

    //white
    x, y, w, h = x+thick, y+thick, w-2*thick, h-2*thick
    dc.DrawRectangle(x, y, w, h)
    dc.SetColor(colornames.White)
    dc.Fill()

    //rgb
    x, y, w, h = x+thick, y+thick, w-2*thick, h-2*thick
    //r
    dc.DrawRectangle(x,y, w, thick)
    dc.SetColor(colornames.Red)
    dc.Fill()
    //g
    dc.DrawRectangle(x,y + thick, thick, h - thick)
    dc.SetColor(colornames.Green)
    dc.Fill()
    //b
    dc.DrawRectangle(x + w - thick,y + thick, thick, h - thick)
    dc.SetColor(colornames.Blue)
    dc.Fill()
    //gray
    dc.DrawRectangle(x,y + w - thick, w, thick)
    dc.SetColor(colornames.Gray)
    dc.Fill()

    x, y, w, h = x+thick, y+thick, w-2*thick, h-2*thick
    imgFill := gg.NewSurfacePattern(img, gg.RepeatNone)
    dc.DrawRectangle(x,y,w,h)
    dc.SetFillStyle(imgFill)
    dc.Fill()
    //dc.DrawRectangle(thickness, thickness, 1- 2 * thickness, 1 - 2 * thickness)
    //dc.SetColor(colornames.Black)
    //dc.Fill()

    //mask := dc.AsMask()

    return dc.Image()
}

var _ gochrom.Renderer = &CIE{}

func (c *CIE) ChangeImage() {
    c.isFramed = !c.isFramed
}

func (c *CIE) Render(w, h int) image.Image {
    dc := gg.NewContext(w, h)

    //set background color and image
    dc.SetColor(colornames.White)
    dc.Clear()
    var img image.Image
    if c.isFramed{
        img = c.framedDiagImg
    }else{
        img = c.diagImg
    }
    resized := normde.ResizeImage(img, fyne.Size{w, h})
    dc.DrawImage(resized, 0, 0)

    //add underlying CIE diagram shape
    dc.InvertY()
    dc.Scale(float64(w)*1.03, float64(h)*0.97)
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
