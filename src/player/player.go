package player

import (
)

const(
	LEFT,RIGHT,UP,DOWN,STATIONARY,MAX_JUMP int = -1,1,1,-1,0,50
)

type Player struct {
	XPos, YPos, jumpHeight int
}

func CreatePlayer(startPos, startHeight int) (*Player) {
	p := new(Player)
	p.XPos = startPos
	p.YPos = startHeight
	p.jumpHeight = MAX_JUMP
	return p
}

func (user *Player) MoveX(direction int) {
	user.XPos = user.XPos + direction
}
func (user *Player) Jump() {
	user.jumpHeight = 0
}
func (user *Player) MoveY() {
	if user.jumpHeight < MAX_JUMP {
		user.YPos = user.YPos + 1
		user.jumpHeight = user.jumpHeight + 1
	} else {
		user.YPos = user.YPos - 1
	}
}