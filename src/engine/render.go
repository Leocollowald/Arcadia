package engine

import (
	"main/src/entity"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"

	"fmt"
)

func (e *Engine) Rendering() {
	rl.ClearBackground(rl.Blue)
}

func (e *Engine) HomeRendering() {
	rl.DrawTexture(e.Sprites["BACKGROUND"], 0, 0, rl.RayWhite)

}

func (e *Engine) HistoryRendering() {
	rl.DrawTexture(e.Sprites["HISTORY"], 0, 0, rl.RayWhite)
	rl.DrawText("[Enter] to Continue", int32(rl.GetScreenWidth())-rl.MeasureText("[Enter] to Continue", 32), int32(rl.GetScreenHeight())/1-35, 30, rl.RayWhite)
}

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)

	rl.BeginMode2D(e.Camera) // On commence le rendu camera

	e.RenderMap()

	e.RenderMonsters()
	e.RenderPlayer()
	e.Displaydealer()
	e.Displaychatuto()

	rl.EndMode2D() // On finit le rendu camera

	// Ecriture fixe (car pas affectée par le mode camera)
	rl.DrawText(fmt.Sprintf("Money : %d $", int32(e.Player.Money)), int32(rl.GetScreenWidth())-rl.MeasureText("Money :", 515), int32(rl.GetScreenHeight())/50, 30, rl.RayWhite)     // Print the money
	rl.DrawText(fmt.Sprintf("Health : %d", int32(e.Player.Health)), int32(rl.GetScreenWidth())-rl.MeasureText("Health :", 500), int32(rl.GetScreenHeight())/18, 30, rl.RayWhite)    // Print the health
	rl.DrawText(fmt.Sprintf("Stamina : %d", int32(e.Player.Stamina)), int32(rl.GetScreenWidth())-rl.MeasureText("Stamina :", 443), int32(rl.GetScreenHeight())/11, 30, rl.RayWhite) // Print the stamina
	rl.DrawText(fmt.Sprintf("Damage : %d", int32(e.Player.Damage)), int32(rl.GetScreenWidth())-rl.MeasureText("Damage :", 453), int32(rl.GetScreenHeight())/8, 30, rl.RayWhite)     // Print the damage
	rl.DrawText(fmt.Sprintf("Speed : %d", int32(e.Player.Speed)), int32(rl.GetScreenWidth())-rl.MeasureText("Speed :", 529), int32(rl.GetScreenHeight())/6, 30, rl.RayWhite)        // Print the speed

}

func (e *Engine) PauseRendering() {
	rl.DrawTexture(e.Sprites["BACKGROUNDPAUSE"], 0, 0, rl.RayWhite)

}

func (e *Engine) EndRendering() {
	rl.DrawTexture(e.Sprites["END"], 0, 0, rl.RayWhite)
	rl.DrawText("[Enter] to Continue", int32(rl.GetScreenWidth())-rl.MeasureText("[Enter] to Continue", 32), int32(rl.GetScreenHeight())/1-35, 30, rl.RayWhite)
	e.Player.Position.X = 615
	e.Player.Position.Y = 1600
}

func (e *Engine) InvRendering() {
	rl.DrawTexture(e.Sprites["SHOP"], 0, 0, rl.RayWhite)
	if len(e.Dealer.Inv) == 0 {
		rl.DrawText("Votre inventaire est vide.", 100, 150, 20, rl.RayWhite)

	} else {
		for i, item := range e.Dealer.Inv {
			rl.DrawTexturePro(
				item.Sprite,
				rl.NewRectangle(0, 0, 64, 64),
				rl.NewRectangle(100, float32(200+i*150), 125, 125),
				rl.Vector2{X: 0, Y: 0},
				0,
				rl.White)
		}
		rl.DrawText(fmt.Sprintf("Argent : %d $", e.Player.Money), 40, 20, 30, rl.RayWhite)
		rl.DrawText("[TAB] to exit Shop", 1620, 1040, 30, rl.RayWhite)
		rl.DrawText("Sword \n 150 $", 900, 710, 30, rl.RayWhite)
		rl.DrawText("Armor \n 200 $", 1230, 710, 30, rl.RayWhite)
		rl.DrawText("Speed \nPotion\n 100 $", 740, 710, 30, rl.RayWhite)
		rl.DrawText("Heal \nPotion\n 70 $", 1070, 710, 30, rl.RayWhite)
	}
	boutonXarmor := int32(1200)
	boutonYarmor := int32(500)
	boutonLargeurarmor := int32(150)
	boutonHauteurarmor := int32(170)
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(boutonXarmor), float32(boutonYarmor), float32(boutonLargeurarmor), float32(boutonHauteurarmor))) && e.Player.Money >= e.Dealer.Inv[1].Price {
			e.buyItem(0)
			e.Player.Health += 100
			e.Player.Armor++
		}
	}
	boutonXsword := int32(920)
	boutonYsword := int32(350)
	boutonLargeursword := int32(60)
	boutonHauteursword := int32(350)
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(boutonXsword), float32(boutonYsword), float32(boutonLargeursword), float32(boutonHauteursword))) && e.Player.Money >= e.Dealer.Inv[0].Price {
			e.buyItem(1)
			e.Player.Damage += 50
			e.Player.Sword++
		}
	}
	boutonXahealpotion := int32(1070)
	boutonYahealpotion := int32(550)
	boutonLargeurahealpotion := int32(60)
	boutonHauteurahealpotion := int32(140)
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(boutonXahealpotion), float32(boutonYahealpotion), float32(boutonLargeurahealpotion), float32(boutonHauteurahealpotion))) && e.Player.Money >= e.Dealer.Inv[2].Price {
			e.buyItem(2)
			e.Player.HealPotion++
		}
	}
	boutonXaspeedpotion := int32(760)
	boutonYaspeedpotion := int32(560)
	boutonLargeuraspeedpotion := int32(60)
	boutonHauteuraspeedpotion := int32(140)
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(boutonXaspeedpotion), float32(boutonYaspeedpotion), float32(boutonLargeuraspeedpotion), float32(boutonHauteuraspeedpotion))) && e.Player.Money >= e.Dealer.Inv[3].Price {
			e.buyItem(3)
			e.Player.Speed += 1
			e.Player.SpeedPotion++
		}
	}
}

func (e *Engine) SettingsRendering() {
	rl.DrawTexture(e.Sprites["SETTINGS"], 0, 0, rl.RayWhite)
	rl.DrawText("Z : Avancer\nQ : Gauche\nD : Droite\nS : Reculer\nSHIFT + Z,Q,S,D : Dash\nE : Attaquer\nL : Ouvrir le Shop\nF : Utiliser Potion de Heal", 720, 300, 30, rl.Black)
	rl.DrawText("Crédits :\nLéo, Thomas, Maxime, Ilies ", 720, 750, 30, rl.Black)
	rl.DrawText("[Echap] to Exit", 1600, 1000, 30, rl.Black)
}

func (e *Engine) RenderItems() {
	if !e.Player.Alive {
		return // Ne pas afficher les items si le joueur est mort
	}

	// Parcourir l'inventaire du joueur et afficher les items en haut à gauche
	for i, item := range e.Player.Inventory {
		rl.DrawText(item.Name, 10, int32(100+i*100), 20, rl.RayWhite)
	}
}

func (e *Engine) GameOverRendering() {
	rl.DrawTexture(e.Sprites["DEAD"], 0, 0, rl.RayWhite)
	rl.DrawText("[Enter] to Respawn", int32(rl.GetScreenWidth())-rl.MeasureText("[Enter] to Respawn", 32), int32(rl.GetScreenHeight())/1-35, 30, rl.RayWhite)

}

func (e *Engine) RenderPlayer() {

	if e.Player.IsAlive {
		rl.DrawTexturePro(
			e.Player.Sprite, //normal
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)

	} else {
		rl.DrawTexturePro(
			e.Player.Sprite, // invertion horizontal
			rl.NewRectangle(0, 0, -100, 100),
			rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}

}

func (e *Engine) RenderMonsters() {
	for _, monster := range e.Monsters {
		rl.DrawTexturePro(
			monster.Sprite,
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(monster.Position.X, monster.Position.Y, 150, 150),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}
}

func (e *Engine) RenderDialog(m entity.Monster, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(m.Position.X)+70,
		int32(m.Position.Y)+50,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}

func (e *Engine) RenderDialogDealer(d entity.Dealer, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(d.Position.X),
		int32(d.Position.Y)+50,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}

func (e *Engine) RenderDialogChatuto(chatuto entity.Chatuto, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(chatuto.Position.X)-150,
		int32(chatuto.Position.Y)+50,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}

func (e *Engine) Normalexplanation(m entity.Dealer, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(m.Position.X),
		int32(m.Position.Y)+50,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}
func (e *Engine) Displaydealer() {
	rl.DrawTexturePro(
		e.Dealer.Sprite, //normal
		rl.NewRectangle(0, 0, 100, 100),
		rl.NewRectangle(e.Dealer.Position.X, e.Dealer.Position.Y, 150, 150),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)

	for _, item := range e.Dealer.Inv {
		text := item.Name
		rl.DrawText(text, 100, int32(150+50), 20, rl.RayWhite)

	}
}

func (e *Engine) Displaychatuto() {
	rl.DrawTexturePro(
		e.Chatuto.Sprite, //normal
		rl.NewRectangle(0, 0, 100, 100),
		rl.NewRectangle(e.Chatuto.Position.X, e.Chatuto.Position.Y, 150, 150),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)
}

func (e *Engine) Ciphertalk(s string) string {
	var s2 string
	for _, char := range s {
		char = char + 2
		s2 = s2 + string(char)
	}
	return s2
}

func (e *Engine) Robottalk(s string) string {
	var s2 string
	var l string
	for _, char := range s {
		x := int(char)
		l = ""
		for x/2 != 0 {
			r := x % 2
			l = strconv.Itoa(r) + l
			x = x / 2
		}
		l = "1" + l
		for len(l) < 8 {
			l = "0" + l
		}
		s2 = s2 + l
	}
	return s2
}
