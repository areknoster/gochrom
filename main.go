package main

import (
    "fyne.io/fyne"
    "fyne.io/fyne/app"
    "fyne.io/fyne/layout"
    "github.com/areknoster/gochrom/pkg/plane"
    "github.com/areknoster/gochrom/pkg/state"
    "github.com/areknoster/gochrom/pkg/render"
    "github.com/sirupsen/logrus"
)

type Config struct{
    title string
    canvasSize fyne.Size
}

func main(){
    logrus.SetLevel(logrus.ErrorLevel)
    cfg := Config{
        title:      "GoFill",
        canvasSize: fyne.Size{700, 700},
    }

    fyneApp := app.New()
    window :=fyneApp.NewWindow(cfg.title)

    ss := state.NewStateStorage()

    CIArenderer := render.NewCIA(ss)

    CIAPlane, CIASetMode := plane.NewPlane(CIArenderer, cfg.canvasSize)
    ss.AddRefresh(CIAPlane.Refresh)
    CIASetMode(plane.NewNullMode())

    SPDPlane, SPDSetMode := plane.NewPlane(CIArenderer, cfg.canvasSize)
    ss.AddRefresh(SPDPlane.Refresh)
    SPDSetMode(plane.NewSPDEditMode(ss))


    container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), CIAPlane, SPDPlane)
    window.SetContent(container)
    window.SetFixedSize(true)
    window.ShowAndRun()


}

