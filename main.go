package main

import (
	"log"
	"main/src/engine"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {

	var h bool
	var godmode bool

	// Vérifier les arguments passés
	for _, arg := range os.Args {
		if arg == "-h" || arg == "-help" {
			h = true
		}
		if arg == "-godmode" {
			godmode = true
		}
	}

	// Si l'option d'aide est demandée, afficher l'aide et quitter
	if h {
		log.Println("-godmode")
		return
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

	// Exécuter l'engine
	e.Run()
	e.Unload()
	e.Close()
}
