package player

import (
	"azul3d.org/gfx.v1"
	"image"
)

const (
	LENGTH,HEIGHT int = 10,10
)

func DrawPlayer(user *Player, r gfx.Renderer) {
	screenHeight := r.Bounds().Dy()
	r.Clear(image.Rect(user.XPos - (LENGTH/2), screenHeight - user.YPos, user.XPos + (LENGTH/2), screenHeight - user.YPos - HEIGHT), gfx.Color{0,1,0,0})
	
}

