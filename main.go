package main

import (
	"log"
	"main/src/engine"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {

	var godmode bool

	// Vérifier si l'argument -godmode est passé
	for _, arg := range os.Args {
		if arg == "--godmode" {
			godmode = true
			break
		}
	}

	var e engine.Engine

	e.Init()

	if godmode {
		log.Println("Godmode activé")
		e.EnableGodMode()
	}

	e.Load()
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	e.Run()
	e.Unload()
	e.Close()
}
