package engine

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Layer struct {
	Data    []int   `json:"data"`
	Height  int     `json:"height"`
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Opacity float32 `json:"opacity"`
	Type    string  `json:"type"`
	Visible bool    `json:"visible"`
	Width   int     `json:"width"`
	X       int     `json:"x"`
	Y       int     `json:"y"`
}

type TileSet struct {
	Columns     int    `json:"columns"`
	FirstGid    int    `json:"firstgid"`
	Image       string `json:"image"`
	ImageHeight int    `json:"imageheight"`
	ImageWidth  int    `json:"imagewidth"`
	Margin      int    `json:"margin"`
	Name        string `json:"name"`
	Spacing     int    `json:"spacing"`
	TileCount   int    `json:"tilecount"`
	TileHeight  int    `json:"tileheight"`
	TileWidth   int    `json:"tilewidth"`
}

type MapJSON struct {
	CompressionLevel int       `json:"compressionLevel"`
	Height           int       `json:"height"`
	Infinite         bool      `json:"infinite"`
	Layers           []Layer   `json:"layers"`
	NextLayerID      int       `json:"nextlayerid"`
	NextObjectID     int       `json:"nextobjectid"`
	Orientation      string    `json:"orientation"`
	RenderOrder      string    `json:"renderorder"`
	TiledVersion     string    `json:"tiledversion"`
	TileHeight       int       `json:"tileheight"`
	TileSets         []TileSet `json:"tilesets"`
	TileWidth        int       `json:"tilewidth"`
	Type             string    `json:"type"`
	Version          string    `json:"version"`
	Width            int       `json:"width"`
}

func (e *Engine) InitMap(mapFile string) {
	/*
		Naive & slow map loader, render all layers everywhere each frame:
		- Parse JSON
		- Load required textures
		- For each layer
			- For each tile
				- Find closest TileSet GID to select correct texture
				- Get X and Y coordinates of the tile
				- Draw tile
				- Move to next position (line, or column)
	*/
	file, err := os.ReadFile(mapFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	json.Unmarshal(file, &e.MapJSON)

	//Load all required textures from TileSets
	for _, TileSet := range e.MapJSON.TileSets {
		path := path.Dir(mapFile) + "/"
		e.Sprites[TileSet.Name] = rl.LoadTexture(path + TileSet.Image)
	}
}

func (e *Engine) RenderMap() {
	/*
		Naive & slow map loader, render all layers everywhere each frame:
		- Parse JSON
		- Load required textures
		- For each layer
			- For each tile
				- Find closest TileSet GID to select correct texture
				- Get X and Y coordinates of the tile
				- Draw tile
				- Move to next position (line, or column)
	*/

	// Prepare source and destination rectangle (only X and Y will change on both)
	srcRectangle := rl.Rectangle{X: 0, Y: 0, Width: float32(e.MapJSON.TileHeight), Height: float32(e.MapJSON.TileHeight)}
	destRectangle := rl.Rectangle{X: 0, Y: 0, Width: float32(e.MapJSON.TileWidth), Height: float32(e.MapJSON.TileWidth)}
	column_counter := 0

	for _, Layer := range e.MapJSON.Layers {
		for _, tile := range Layer.Data {
			if tile != 0 {
				wantedTileSet := e.MapJSON.TileSets[0]
				for _, TileSet := range e.MapJSON.TileSets { // Get correct texture
					if TileSet.FirstGid < tile {
						wantedTileSet = TileSet
					}
				}

				index := tile - wantedTileSet.FirstGid

				srcRectangle.X = float32(index)
				srcRectangle.Y = 0

				if index > wantedTileSet.Columns { // If Tile number exceeds columns (overflow), adjust, find X and Y coordinates
					srcRectangle.X = float32(index % wantedTileSet.Columns)
					srcRectangle.Y = float32(index / wantedTileSet.Columns)
				}

				srcRectangle.X *= float32(e.MapJSON.TileWidth)
				srcRectangle.Y *= float32(e.MapJSON.TileHeight)

				rl.DrawTexturePro(
					e.Sprites[wantedTileSet.Name],
					srcRectangle,
					destRectangle,
					rl.Vector2{X: 0, Y: 0},
					0,
					rl.White,
				)
			}

			// After each draw, move to the right. When at max width, new line (like a typewriter)
			destRectangle.X += 32
			column_counter += 1
			if column_counter >= e.MapJSON.Width {
				destRectangle.X = 0
				destRectangle.Y += 32
				column_counter = 0
			}

		}
		destRectangle.X, destRectangle.Y, column_counter = 0, 0, 0
	}
}

func (e *Engine) CheckCollisionstiles() {

	// Définir la hitbox du joueur
	playerHitbox := rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 32, 32)
	tileSize := 32
	mapWidth := 200
	newX, newY := e.Player.Position.X, e.Player.Position.Y

	collisionLayerName := "Collision Layer"

	for _, layer := range e.MapJSON.Layers {
		if layer.Name == collisionLayerName {
			for index, tile := range layer.Data {
				if tile != 0 {

					tileX := (index % mapWidth) * tileSize -60
					tileY := (index / mapWidth) * tileSize -60

					tileHitbox := rl.Rectangle{
						X:      float32(tileX),
						Y:      float32(tileY),
						Width:  float32(tileSize),
						Height: float32(tileSize),
					}

					// Vérification de la collision horizontale (axe X) d'abord
					updatedPlayerHitboxX := rl.NewRectangle(newX, e.Player.Position.Y, 32, 32)
					if CheckCollisionenvironnement(updatedPlayerHitboxX, tileHitbox) {
						if newX < tileHitbox.X {
							newX = tileHitbox.X - playerHitbox.Width // Collision à gauche
						} else if newX > tileHitbox.X {
							newX = tileHitbox.X + tileHitbox.Width // Collision à droite
						}
					}

					// Vérification de la collision verticale (axe Y) en utilisant la nouvelle position X
					updatedPlayerHitboxY := rl.NewRectangle(newX, newY, 32, 32)
					if CheckCollisionenvironnement(updatedPlayerHitboxY, tileHitbox) {
						if newY < tileHitbox.Y {
							newY = tileHitbox.Y - playerHitbox.Height // Collision au-dessus
						} else if newY > tileHitbox.Y {
							newY = tileHitbox.Y + tileHitbox.Height // Collision en dessous
						}
					}
				}
			}
		}
	}

	// Mettre à jour la position du joueur après avoir vérifié les collisions
	e.Player.Position.X = newX
	e.Player.Position.Y = newY
}

// Fonction pour vérifier si deux rectangles se chevauchent
func CheckCollisionenvironnement(playerHitbox, tileHitbox rl.Rectangle) bool {
    playerRight := playerHitbox.X + playerHitbox.Width
    playerBottom := playerHitbox.Y + playerHitbox.Height
    tileRight := tileHitbox.X + tileHitbox.Width
    tileBottom := tileHitbox.Y + tileHitbox.Height

    // Vérification des chevauchements sur les deux axes
    horizontalOverlap := playerHitbox.X < tileRight && playerRight > tileHitbox.X
    verticalOverlap := playerHitbox.Y < tileBottom && playerBottom > tileHitbox.Y

    // Retourne vrai si les deux axes se chevauchent
    return horizontalOverlap && verticalOverlap
}
