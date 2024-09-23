package engine

import (
	"main/src/entity"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type menu int

const (
	HOME     menu = iota
	SETTINGS menu = iota
	PLAY     menu = iota
	HISTORY  menu = iota
)

type engine int

const (
	INGAME   engine = iota
	PAUSE    engine = iota
	GAMEOVER engine = iota
	INV      engine = iota
	DEALER   engine = iota
)

type Engine struct {
	Player                  entity.Player
	Monsters                []entity.Monster
	Item                    []item.Item
	Dealer                  entity.Dealer
	Chatuto                	entity.Chatuto
	InitialMonsterPositions []rl.Vector2
	InitialMonsterHealths   []int

	Music       rl.Music
	MusicVolume float32

	Sprites map[string]rl.Texture2D

	Camera rl.Camera2D

	MapJSON MapJSON

	IsRunning   bool
	StateMenu   menu
	StateEngine engine
}
