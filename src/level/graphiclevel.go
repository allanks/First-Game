package level

import (
	"azul3d.org/gfx.v1"
	"image"
)

const (
	platformWidth int = 10
)

func DrawLevel(l *Level, r gfx.Renderer) {
	screenHeight := r.Bounds().Dy()
	for _, p := range l.platforms {
		r.Clear(image.Rect(p.startPos, screenHeight - p.height, p.startPos + p.length, screenHeight - p.height + platformWidth), gfx.Color{0,0,1,0})
	}
}

