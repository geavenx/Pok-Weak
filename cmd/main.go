package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/geavenx/pokeweak/internal"
)

func main() {

	app := &cli.App{
		Name:  "PokéWeak",
		Usage: "Your handy CLI pokédex!",

		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "types",
				Aliases: []string{"t"},
				Usage:   "Output the type of the pokémon",
			},
			&cli.BoolFlag{
				Name:    "damage",
				Aliases: []string{"d"},
				Usage:   "Output the damage relations of the pokémon",
			},
		},

		Action: func(cCtx *cli.Context) error {
			p, err := pokemon.GetPokemon(cCtx)
			if err != nil {
				return err
			}

			if cCtx.Bool("types") {
				fmt.Printf("\n%v\n", p.Type())
			}

			if cCtx.Bool("damage") {
				dmgRelations, err := p.GetDamageRelations()
				if err != nil {
					return err
				}

				fmt.Println("\nATTACKING")
				fmt.Println(dmgRelations.NoEffectOn())
				fmt.Println(dmgRelations.NotVeryEffOn())
				fmt.Println(dmgRelations.SuperEffOn())

				fmt.Println("\nDEFENDING")
				fmt.Println(dmgRelations.ImmuneTo())
				fmt.Println(dmgRelations.NotVeryEffTo())
				fmt.Println(dmgRelations.SuperEffTo())
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("You are dumb.\n", err)
	}
}
