package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/geavenx/pokeweak/internal"
)

func main() {
	var typeQuery string

	app := &cli.App{
		Name:  "PokéWeak",
		Usage: "Your handy CLI pokédex!",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "types",
				Value:       "all",
				Usage:       "Output the type of the pokémon",
				Destination: &typeQuery,
			},
		},

		Action: func(cCtx *cli.Context) error {
			p, err := pokemon.GetPokemon(cCtx)
			if err != nil {
				return err
			}

			if typeQuery == "all" {
				fmt.Printf("Type(s): %v\n", p.Type())
			} else {
				fmt.Printf("Name: %v\n", p.Name())
				fmt.Printf("Type(s): %v\n", p.Type())
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("You are dumb.", err)
	}
}
