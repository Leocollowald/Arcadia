package entity

import (
	"fmt"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position  rl.Vector2
	Health    int
	Money     int
	Damage    int
	Speed     float32
	Inventory []item.Item
	Stamina   int
	Alive     bool

	SwordSound rl.Sound

	IsAlive bool

	Sprite rl.Texture2D
}

func (p *Player) Attack(m *Monster) {
	m.Health -= 1
}

func (p *Player) Spend(m *Dealer) {
	m.Money -= 1
	p.Inventory = append(p.Inventory, item.Item{})

}

func (p *Player) ToString() {
	fmt.Printf(`
	Joueur:
		Vie: %d,
		Argent: %d,
		Inventaire: %+v
	
	\n`, p.Health, p.Money, p.Inventory)
}
