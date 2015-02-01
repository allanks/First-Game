package level

import (
	"player"
)

type Platform struct {
	startPos, height, length int
}

type Level struct {
	platforms []Platform
	playerStartXPos, playerStartYPos int
}

func CreateLevel() (*Level) {
	
	l := new(Level)
	
	l.platforms = append(l.platforms, Platform{0,10,100})
	l.platforms = append(l.platforms, Platform{120,10,100})
	l.platforms = append(l.platforms, Platform{240,10,560})
	l.platforms = append(l.platforms, Platform{100, 50, 80})
	l.platforms = append(l.platforms, Platform{200, 50, 80})
	l.platforms = append(l.platforms, Platform{100, 90, 60})
	l.platforms = append(l.platforms, Platform{100, 130, 40})
	l.platforms = append(l.platforms, Platform{100, 170, 20})
	l.playerStartXPos = 10
	l.playerStartYPos = 20
	
	return l
}

func (l *Level) CreatePlayer() (*player.Player) {
	p := player.CreatePlayer(l.playerStartXPos, l.playerStartYPos)
	return p
}

func (p *Platform) isPlayerOnPlatform(user *player.Player) bool {
	if user.YPos == p.height && user.XPos >= p.startPos && user.XPos <= (p.startPos+p.length) {
		return true
	}
	return false
}

func (l *Level) IsPlayerOnPlatform(user *player.Player) bool {
	if user.YPos < 0 {
		user.XPos = l.playerStartXPos
		user.YPos = l.playerStartYPos
		return true
	}
	for _,p := range l.platforms {
		if p.isPlayerOnPlatform(user) {
			return true
		}
	}
	return false
}

