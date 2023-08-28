package internal

import (
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

/*

THIS IS JUST A TEST PLACEHOLDER

*/


type Star struct {
    // X and Y are the coordinates of the center of the Star
    X float64
    Y float64

    // Points is the number of points on the Star
    Size float64

    // Radius is the radius of the Star
    Brightness float64
}

func NewStar(x float64, y float64, size float64, brightness float64) *Star {
    return &Star{
        X: x,
        Y: y,
        Size: size,
        Brightness: brightness,
    }
}

func (s *Star) Draw(gc *draw2dimg.GraphicContext) {
    gc.SetFillColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
    gc.SetStrokeColor(color.RGBA{0xff, 0xff, 0xff, 0xff})

    gc.BeginPath()
    draw2dkit.Circle(gc, s.X, s.Y, s.Size)
    gc.FillStroke()
    gc.Close()
}
