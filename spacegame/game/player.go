package game

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Player the user controls
var (
	Player player
)

type player struct {
	object
	lives int
}

// Controls the player
type Controls interface {
	IncreaseRspd()
	DecreaseRspd()
}

// CreatePlayer at coordinates
func CreatePlayer(x float64, y float64) {
	Player.image, _, err = ebitenutil.NewImageFromFile("assets/player.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}

	Player.xPos = x
	Player.yPos = y
	Player.xSpd = 0
	Player.ySpd = 0
	Player.rSpd = 1
	Player.lives = 3

	All = append(All, Player.object)
}

func (plr *player) IncreaseRspd() {
	plr.rSpd += 0.1
}

func (plr *player) DecreaseRspd() {
	plr.rSpd -= 0.1
}
