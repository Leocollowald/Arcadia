package engine

import (
	"fmt"
	"main/src/entity"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// 0 si l'histoire n'a pas encore été lu, 1 si l'histoire a été lu
var ReadHistory = 0

func (e *Engine) HomeLogic() {

	if rl.IsKeyPressed(rl.KeyP) {
		e.StateMenu = SETTINGS
	}

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/fairy-lands-fantasy-music-in-a-magical-forest-fantasy.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)

	boutonXplay := int32(450)
	boutonYplay := int32(830)
	boutonLargeurplay := int32(200)
	boutonHauteurplay := int32(150)
	//Menus
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && ReadHistory == 1 {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(boutonXplay), float32(boutonYplay), float32(boutonLargeurplay), float32(boutonHauteurplay))) {
			e.StateMenu = PLAY
			e.StateEngine = INGAME
			rl.StopMusicStream(e.Music)
		}
	}
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && ReadHistory == 0 {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(boutonXplay), float32(boutonYplay), float32(boutonLargeurplay), float32(boutonHauteurplay))) {
			e.StateMenu = HISTORY
			ReadHistory = 1
		}
	}
	boutonXexit := int32(1270)
	boutonYexit := int32(830)
	boutonLargeurexit := int32(200)
	boutonHauteurexit := int32(150)
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(boutonXexit), float32(boutonYexit), float32(boutonLargeurexit), float32(boutonHauteurexit))) {
			e.IsRunning = false
		}
	}
}

func (e *Engine) PauseLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = INGAME
	}
	boutonXresume := int32(730)
	boutonYresume := int32(1000)
	boutonLargeurresume := int32(170)
	boutonHauteurresume := int32(60)
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(boutonXresume), float32(boutonYresume), float32(boutonLargeurresume), float32(boutonHauteurresume))) {
			e.StateEngine = INGAME
		}
	}
	boutonXmenu := int32(1200)
	boutonYmenu := int32(1000)
	boutonLargeuremenu := int32(170)
	boutonHauteuremenu := int32(60)
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(boutonXmenu), float32(boutonYmenu), float32(boutonLargeuremenu), float32(boutonHauteuremenu))) {
			e.StateMenu = HOME
			rl.StopMusicStream(e.Music)
		}
	}
	//Musique
	rl.UpdateMusicStream(e.Music)
}

// Explication de l'histoire
func (e *Engine) HistoryLogic() {
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.StateMenu = PLAY
		e.StateEngine = INGAME
		rl.StopMusicStream(e.Music)
	}
}

func (e *Engine) SettingsLogic() {
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.StateMenu = HOME
		rl.StopMusicStream(e.Music)
	}
}

var Stamina = false

func (e *Engine) InGameLogic() {

	// Dealer logic
	e.dealerCollisions()

	// Mouvement
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		e.Player.Position.Y -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		e.Player.Position.Y += e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		e.Player.Position.X -= e.Player.Speed
		e.Player.IsAlive = false
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		e.Player.Position.X += e.Player.Speed
		e.Player.IsAlive = true
	}

	// Mouvement Dash
	if rl.IsKeyDown(rl.KeyW) && rl.IsKeyPressed(rl.KeyLeftShift) {
		if e.Player.Stamina >= 25 {
			e.Player.Position.Y -= e.Player.Speed + 30
			e.Player.Stamina -= 25
		}
	} else if e.Player.Stamina < 100 && !Stamina {
		Stamina = true
		go func() {
			for e.Player.Stamina < 100 {
				e.Player.Stamina += 1
				time.Sleep(70 * time.Millisecond)
			}
			Stamina = false
		}()
	}
	if rl.IsKeyDown(rl.KeyS) && rl.IsKeyPressed(rl.KeyLeftShift) {
		if e.Player.Stamina >= 25 {
			e.Player.Position.Y += e.Player.Speed + 30
			e.Player.Stamina -= 25
		}
	}
	if rl.IsKeyDown(rl.KeyA) && rl.IsKeyPressed(rl.KeyLeftShift) {
		if e.Player.Stamina >= 25 {
			e.Player.Position.X -= e.Player.Speed + 30
			e.Player.IsAlive = false
			e.Player.Stamina -= 25
		}
	}
	if rl.IsKeyDown(rl.KeyD) && rl.IsKeyPressed(rl.KeyLeftShift) {
		if e.Player.Stamina >= 25 {
			e.Player.Position.X += e.Player.Speed + 10
			e.Player.IsAlive = true
			e.Player.Stamina -= 25
		}
	}

	// Camera
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X + 70, Y: e.Player.Position.Y + 70}
	e.Camera.Offset = rl.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2}

	// Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = PAUSE
	}

	e.CheckCollisions()

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {

		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)
}

// Inv check
func (e *Engine) InvLogic() {

	if rl.IsKeyPressed(rl.KeyI) {
		if len(e.Player.Inventory) > 0 {
			item := e.Player.Inventory[0]
			fmt.Printf("Vous avez utilisé %s.\n", item.Name)
			e.Player.Inventory = e.Player.Inventory[1:]
		} else {
			fmt.Println("Je n'ai pas d'items a échanger")
		}

	}
}

func (e *Engine) GameOverLogic() {
	if rl.IsKeyDown(rl.KeyEnter) {
		e.StateMenu = PLAY
		e.StateEngine = INGAME
		rl.StopMusicStream(e.Music)
		e.Player.Position.X = 615
		e.Player.Position.Y = 1600
		e.Player.Health = 100
		e.Player.Stamina = 100
	}
}

func (e *Engine) CheckCollisions() {

	e.MonsterCollisions()
}

var Dead = false
var Attack = false

func (e *Engine) MonsterCollisions() {

	// Enlever les monstres morts
	var newMonsters []entity.Monster
	for _, monster := range e.Monsters {
		if monster.Health > 0 {
			newMonsters = append(newMonsters, monster)
		} else {
			rl.UnloadTexture(monster.Sprite)
		}
	}

	e.Monsters = newMonsters

	if e.Player.Health <= 0 {
		e.StateEngine = GAMEOVER
	}
	if e.StateEngine == GAMEOVER {
		// Réinitialiser les mobs
		for i := range e.Monsters {
			e.Monsters[i].Health = e.InitialMonsterHealths[i]
			e.Monsters[i].Position = e.InitialMonsterPositions[i]
		}
	}

	if e.StateEngine == INGAME {
		fps := rl.GetFPS()
		rl.DrawText(fmt.Sprintf("FPS: %d", fps), 1800, 10, 24, rl.RayWhite)
	}

	for i, monster := range e.Monsters {

		// porté aggro
		if monster.Position.X > e.Player.Position.X-150 &&
			monster.Position.X < e.Player.Position.X+150 &&
			monster.Position.Y > e.Player.Position.Y-150 &&
			monster.Position.Y < e.Player.Position.Y+150 {

			e.NormalTalk(monster, fmt.Sprintf("%d HP", monster.Health))
			if e.Player.Position.X < e.Monsters[i].Position.X+30 && e.Monsters[i].Health > 0 {
				e.Monsters[i].Position.X -= 2
			}
			if e.Player.Position.Y < e.Monsters[i].Position.Y+30 && e.Monsters[i].Health > 0 {
				e.Monsters[i].Position.Y -= 2
			}
			if e.Player.Position.Y > e.Monsters[i].Position.Y-30 && e.Monsters[i].Health > 0 {
				e.Monsters[i].Position.Y += 2
			}
			if e.Player.Position.X > e.Monsters[i].Position.X-30 && e.Monsters[i].Health > 0 {
				e.Monsters[i].Position.X += 2
			}

			// porté attaque Monstre
			if monster.Position.X > e.Player.Position.X-35 &&
				monster.Position.X < e.Player.Position.X+35 &&
				monster.Position.Y > e.Player.Position.Y-35 &&
				monster.Position.Y < e.Player.Position.Y+35 {
				//Damage to Player
				if monster.Health > 0 && !Dead {
					Dead = true
					go func() {
						for e.Monsters[i].Health > 0 && e.Player.Health > 0 {
							// Vérifier la distance entre le joueur et le monstre
							distance := math.Sqrt(math.Pow(float64(e.Player.Position.X-e.Monsters[i].Position.X), 2) + math.Pow(float64(e.Player.Position.Y-e.Monsters[i].Position.Y), 2))
							if distance > 100 {
								// Arrêter l'attaque si le joueur est trop loin
								Dead = false
								return
							}
							e.Player.Health -= monster.Damage
							time.Sleep(1 * time.Second)
						}
						Dead = false
					}()
				}
				if e.Player.Position.Y < e.Monsters[i].Position.Y+30 && e.Monsters[i].Health > 0 {
					e.Monsters[i].Position.Y -= 3
				}
				if e.Player.Position.Y > e.Monsters[i].Position.Y-30 && e.Monsters[i].Health > 0 {
					e.Monsters[i].Position.Y += 3
				}
				if e.Player.Position.X > e.Monsters[i].Position.X-30 && e.Monsters[i].Health > 0 {
					e.Monsters[i].Position.X += 3
				}

				if monster.Position.X > e.Player.Position.X-31 &&
					monster.Position.X < e.Player.Position.X+31 &&
					monster.Position.Y > e.Player.Position.Y-31 &&
					monster.Position.Y < e.Player.Position.Y+31 {
				}

			}
		}

		//Porté attaque joueur
		if monster.Position.X > e.Player.Position.X-100 &&
			monster.Position.X < e.Player.Position.X+100 &&
			monster.Position.Y > e.Player.Position.Y-100 &&
			monster.Position.Y < e.Player.Position.Y+100 {
			if e.Monsters[i].Health <= 0 && e.Monsters[i].IsAlive {
				e.Monsters[i].IsAlive = false
				e.Player.Money += e.Monsters[i].Worth
			}
			if rl.IsKeyPressed(rl.KeyE) && e.Monsters[i].Health > 0 && !Attack {
				Attack = true
				go func() {
					for e.Monsters[i].Health > 0 {
						if rl.IsKeyUp(rl.KeyE) {
							Attack = false
							return
						}
						e.Monsters[i].Health -= e.Player.Damage
						rl.PlaySound(e.Player.SwordSound)
						e.Player.Stamina -= 15
						// Knockback
						if e.Player.Position.X > e.Monsters[i].Position.X {
							e.Monsters[i].Position.X -= 30
						}
						if e.Player.Position.X < e.Monsters[i].Position.X {
							e.Monsters[i].Position.X += 30
						}
						time.Sleep(1 * time.Second)
					}
					Attack = false
				}()
			}
		}
	}
}

func (e *Engine) dealerCollisions() {

	if e.Dealer.Position.X > e.Player.Position.X-20 &&
		e.Dealer.Position.X < e.Player.Position.X+20 &&
		e.Dealer.Position.Y > e.Player.Position.Y-20 &&
		e.Dealer.Position.Y < e.Player.Position.Y+20 {

		if e.Dealer.Name == "yannis" {
			e.RenderDialogDealer(e.Dealer, "Bonjour, appuie sur [L] pour ouvrir le shop")
			if rl.IsKeyPressed(rl.KeyL) {
				e.updatedealer()
				e.StateEngine = DEALER
			}
		}
	}
}

func (e *Engine) NormalTalk(m entity.Monster, sentence string) {
	e.RenderDialog(m, sentence)
}
func (e *Engine) dealertalk(m entity.Dealer, sentence string) {
	e.Normalexplanation(m, sentence)
}

func (e *Engine) updatedealer() {
	if rl.IsKeyPressed(rl.KeyTab) {
		e.StateEngine = INGAME

	}

}

func (e *Engine) buyItem(index int) {
	//item := e.Dealer.Inv[index]
	if index < 0 || index >= len(e.Dealer.Inv) {
		fmt.Println("Index invalide")
		return
	}
}

/*if e.Player.Money >= item.Price {
	e.Player.Money -= item.Price
	e.Player.Inventory = append(e.Player.Inventory, item)
	rl.DrawText("Vous avez acheté %s pour %d pièces\n", item.Name, item.Price, 100, 100, 30, rl.RayWhite)
	fmt.Println(e.Player.Inventory)
} else {
	rl.DrawText("Pas assez d'argent !", 100, 100, 30, rl.RayWhite)
}
*/