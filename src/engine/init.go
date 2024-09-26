package engine

import (
	"fmt"
	"main/src/entity"
	"main/src/item"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
)

func (e *Engine) Init() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "After Death")

	// Mode plein ecran
	rl.ToggleFullscreen()

	rl.EndDrawing()

	// Initialisation des variables de l'engine
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)

	// Initialisation des composants du jeu
	e.InitEntities()
	e.InitCamera()
	e.InitMusic()
	e.InitDealer()
	e.InitChatuto()
	e.InitItem()
	e.InitMap("textures/map/tilesets/map.json")

}

func (e *Engine) InitEntities() {

	e.Player = entity.Player{
		Position:    rl.Vector2{X: 615, Y: 1600},
		Health:      100,
		Money:       1000,
		Damage:      10,
		Speed:       4,
		Stamina:     100,
		Inventory:   []item.Item{},
		Sword:       0,
		Armor:       0,
		HealPotion:  0,
		SpeedPotion: 0,

		IsAlive: true,

		Sprite: e.Player.Sprite,
	}

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 1500, Y: 1600},
		Health:   20,
		Damage:   5,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/slime/Slime-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 2000, Y: 1600},
		Health:   20,
		Damage:   5,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/slime/Slime-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 2100, Y: 1650},
		Health:   20,
		Damage:   5,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/slime/Slime-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 2100, Y: 1550},
		Health:   20,
		Damage:   5,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/slime/Slime-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 2300, Y: 1150},
		Health:   40,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 2350, Y: 1150},
		Health:   40,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 2400, Y: 1150},
		Health:   40,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 2750, Y: 850},
		Health:   40,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 2750, Y: 900},
		Health:   40,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 2825, Y: 825},
		Health:   40,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 2825, Y: 875},
		Health:   40,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 2825, Y: 925},
		Health:   40,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 3200, Y: 550},
		Health:   40,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 3200, Y: 600},
		Health:   40,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 3275, Y: 525},
		Health:   70,
		Damage:   20,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Armored Orc/Armored Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 3275, Y: 575},
		Health:   70,
		Damage:   20,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Armored Orc/Armored Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 3275, Y: 625},
		Health:   70,
		Damage:   20,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Armored Orc/Armored Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 3550, Y: 500},
		Health:   70,
		Damage:   20,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Armored Orc/Armored Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 3550, Y: 550},
		Health:   70,
		Damage:   20,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Armored Orc/Armored Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 3600, Y: 475},
		Health:   40,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 3600, Y: 525},
		Health:   40,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 3600, Y: 575},
		Health:   40,
		Damage:   15,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 3650, Y: 500},
		Health:   70,
		Damage:   20,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Armored Orc/Armored Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 3650, Y: 550},
		Health:   70,
		Damage:   20,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Armored Orc/Armored Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 4000, Y: 550},
		Health:   150,
		Damage:   30,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Elite Orc/Elite Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 4000, Y: 500},
		Health:   150,
		Damage:   30,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Elite Orc/Elite Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 4700, Y: 700},
		Health:   70,
		Damage:   20,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Armored Orc/Armored Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 4700, Y: 650},
		Health:   70,
		Damage:   20,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Armored Orc/Armored Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 4750, Y: 720},
		Health:   150,
		Damage:   30,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Elite Orc/Elite Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 4750, Y: 675},
		Health:   150,
		Damage:   30,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Elite Orc/Elite Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 4750, Y: 625},
		Health:   150,
		Damage:   30,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Elite Orc/Elite Orc.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 5400, Y: 1100},
		Health:   150,
		Damage:   30,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Elite Orc/Elite Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 5450, Y: 1100},
		Health:   150,
		Damage:   30,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Elite Orc/Elite Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 5500, Y: 1100},
		Health:   150,
		Damage:   30,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Elite Orc/Elite Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 5450, Y: 1150},
		Health:   150,
		Damage:   30,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Elite Orc/Elite Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 5500, Y: 1150},
		Health:   150,
		Damage:   30,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Elite Orc/Elite Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 5550, Y: 1150},
		Health:   150,
		Damage:   30,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Elite Orc/Elite Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "eric",
		Position: rl.Vector2{X: 2950, Y: 1600},
		Health:   500,
		Damage:   50,
		Loot:     []item.Item{},
		Worth:    1,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Orc rider/Orc rider-Idle.png"),
	})

	e.InitialMonsterPositions = make([]rl.Vector2, len(e.Monsters))
	e.InitialMonsterHealths = make([]int, len(e.Monsters))
	for i := range e.Monsters {
		e.InitialMonsterPositions[i] = e.Monsters[i].Position
		e.InitialMonsterHealths[i] = e.Monsters[i].Health
	}
	minWorth := 1
	maxWorth := 10
	for i := range e.Monsters {
		e.Monsters[i].Worth = rand.Intn(maxWorth-minWorth+1) + minWorth
	}
}

func (e *Engine) InitCamera() {
	e.Camera = rl.NewCamera2D( //Camera vide, a changer dans chaque logique de scene
		rl.NewVector2(0, 0),
		rl.NewVector2(0, 0),
		0.0,
		2.0,
	)
}

func (e *Engine) InitMusic() {
	rl.InitAudioDevice()

	e.Music = rl.LoadMusicStream("sounds/music/fairy-lands-fantasy-music-in-a-magical-forest-fantasy.mp3")

	rl.PlayMusicStream(e.Music)
	rl.LoadSound("sounds/music/weapswrd-epee.wav")
}

func (e *Engine) InitDealer() {

	e.Dealer = entity.Dealer{
		Inv:      []item.Item{},
		Name:     "yannis",
		Position: rl.NewVector2(1000, 1500),
		Sprite:   rl.LoadTexture("textures/entities/dealer/Swordsman-Idle.png"),
	}
}
func (e *Engine) InitItem() {
	e.Dealer.Inv = append(e.Player.Inventory, item.Item{
		Name:         "Armor",
		Price:        200,
		Sprite:       rl.LoadTexture("textures/shild/shild.png"),
		IsConsumable: true,
		IsEquippable: true,
	})
	e.Dealer.Inv = append(e.Dealer.Inv, item.Item{
		Name:         "Sword",
		Price:        150,
		Sprite:       rl.LoadTexture("textures/sword/sword.png"),
		IsConsumable: true,
		IsEquippable: true,
	})

	e.Dealer.Inv = append(e.Dealer.Inv, item.Item{
		Name:         "Heal Potion",
		Price:        70,
		Sprite:       rl.LoadTexture("textures/potion/PotionYellow.png"),
		IsConsumable: true,
		IsEquippable: true,
	})
	e.Dealer.Inv = append(e.Dealer.Inv, item.Item{
		Name:         "Speed Potion",
		Price:        100,
		Sprite:       rl.LoadTexture("textures/potion/PotionYellow.png"),
		IsConsumable: true,
		IsEquippable: true,
	})
	fmt.Println(e.Player.Inventory)
}

func (e *Engine) InitChatuto() {
	e.Chatuto = entity.Chatuto{
		Name:     "chatuto",
		Position: rl.Vector2{X: 700, Y: 1600},
		Sprite:   rl.LoadTexture("textures/entities/chatuto/Wizard-Idle.png"),
	}
}
