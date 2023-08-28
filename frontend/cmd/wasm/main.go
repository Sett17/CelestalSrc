//go:build js && wasm

package main

import (
	"celestralsrc/frontend/internal"
	"image/color"
	"math"
	"math/rand"
	"syscall/js"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"github.com/markfarnan/go-canvas/canvas"
)

var done chan struct{}

var cvs *canvas.Canvas2d
var width float64
var height float64

var stars []*internal.Star

func main() {
	cvs, _ = canvas.NewCanvas2d(false)
	cvs.Create(int(js.Global().Get("innerWidth").Float()*0.75), int(js.Global().Get("innerHeight").Float()*0.75))

	height = float64(cvs.Height())
	width = float64(cvs.Width())

    //just a test and has nothing to to with the end result!
    N := 400
    for i := 0; i < N; i++ {
        x := float64(cvs.Width()) * rand.Float64()
        y := float64(cvs.Height()) * rand.Float64()
        size := rand.Float64()*1
        brightness := 0.2 + rand.Float64()*0.8
        stars = append(stars, internal.NewStar(x, y, size, brightness))
    }

	cvs.Start(60, Render)

	//go doEvery(renderDelay, Render) // Kick off the Render function as go routine as it never returns
	<-done
}

var t float64 = 0
var step float64 = 0.01

func Render(gc *draw2dimg.GraphicContext) bool {
	gc.SetFillColor(color.RGBA{0xaa, 0x77, 0xff, 0x00}) //background with opacity gives cool effect; TODO change back
	gc.Clear()

    for _, star := range stars {
        star.Draw(gc)
    }

	gc.SetFillColor(color.RGBA{0xff, 0x00, 0xff, 0xdf})
	gc.SetStrokeColor(color.RGBA{0xff, 0x00, 0xff, 0xff})

	gc.BeginPath()
	draw2dkit.Circle(gc, width/2+(135*math.Sin(t)), height/2+(100*math.Cos(t)), 50)
	gc.FillStroke()
	gc.Close()

	t += step
	return true
}
