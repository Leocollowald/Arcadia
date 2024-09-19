package entity

import (
	"fmt"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Dealer struct {
	Name     string
	Position rl.Vector2
	Health   int
	Money    int
	Inv      []item.Item
	IsAlive  bool
	Worth    int
	Damage   int
	Loot     []item.Item
	Sprite   rl.Texture2D
}

func (d *Dealer) ToString() {
	fmt.Printf(`
	Joueur:
		Vie: %d,
		Argent: %d,
		Inventaire: %+v
	
	\n`, d.Health, d.Money, d.Inv)
}
