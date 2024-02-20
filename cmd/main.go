package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/geavenx/pokeweak/internal"
)

func main() {
	app := &cli.App{
		Name:  "PokéWeak",
		Usage: "Your handy CLI pokédex!",
		Action: func(cCtx *cli.Context) error {
			err := pokemon.GetPokemon(cCtx)
			if err != nil {
				return err
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("You are dumb.", err)
	}
}
