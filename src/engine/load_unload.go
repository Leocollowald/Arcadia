package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Load() {
	// Chargement des textures du personnage
	e.Player.Sprite = rl.LoadTexture("textures/entities/soldier/Soldier-Idle.png")
	e.Sprites["BACKGROUND"] = rl.LoadTexture("textures/img/ecran.png")
	e.Sprites["BACKGROUNDPAUSE"] = rl.LoadTexture("textures/img/pause.png")
	e.Sprites["HISTORY"] = rl.LoadTexture("textures/img/intro.png")
	e.Sprites["DEAD"] = rl.LoadTexture("textures/img/dead.png")
	e.Sprites["SHOP"] = rl.LoadTexture("textures/img/shop.png")
	e.Sprites["SETTINGS"] = rl.LoadTexture("textures/img/settings.png")
	e.Sprites["END"] = rl.LoadTexture("textures/img/END.png")
	e.Sprites["MINISWORD"] = rl.LoadTexture("textures/img/minisword.png")
	e.Sprites["MINIARMOR"] = rl.LoadTexture("textures/img/miniarmor.png")
	e.Sprites["MINIHEALPOTION"] = rl.LoadTexture("textures/img/minihealpotion.png")
	e.Sprites["MINISPEEDPOTION"] = rl.LoadTexture("textures/img/minipotionspeed.png")
	e.Music = rl.LoadMusicStream("sounds/music/fairy-lands-fantasy-music-in-a-magical-forest-fantasy.mp3")

	e.Player.SwordSound = rl.LoadSound("sounds/music/weapswrd-epee.wav")
}

func (e *Engine) Unload() {
	// On libère les textures chargées, le joueur, la map, les monstres, etc...
	rl.UnloadTexture(e.Player.Sprite)

	for _, sprite := range e.Sprites {
		rl.UnloadTexture(sprite)
	}

	for _, monster := range e.Monsters {
		rl.UnloadTexture(monster.Sprite)
	}
}
