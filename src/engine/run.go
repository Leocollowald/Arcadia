package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (engine *Engine) Run() {
	rl.SetTargetFPS(120)

	for engine.IsRunning {

		rl.BeginDrawing()

		switch engine.StateMenu {
		case HOME:
			engine.HomeRendering()
			engine.HomeLogic()

		case SETTINGS:
			engine.SettingsLogic()
			engine.SettingsRendering()

		case HISTORY:
			engine.HistoryRendering()
			engine.HistoryLogic()

		case PLAY:
			switch engine.StateEngine {
			case INGAME:
				engine.InGameRendering()
				engine.InGameLogic()

			case PAUSE:
				engine.PauseRendering()
				engine.PauseLogic()

			case INV:
				engine.RenderItems()
				engine.InvLogic()

			case DEALER:
				engine.updatedealer()
				engine.InvRendering()

			case GAMEOVER:
				engine.GameOverLogic()
				engine.GameOverRendering()

			case END:
				engine.EndLogic()
				engine.EndRendering()

			}
		}

		rl.EndDrawing()
	}
}
