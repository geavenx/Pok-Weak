package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

type Pokemon struct {
	Name string  `json:"name"`
	Type []Types `json:"types"`
}

type Types struct {
	Slot int      `json:"slot"`
	Type TypeInfo `json:"type"`
}

type TypeInfo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {
	app := &cli.App{
		Name:  "PokéWeak",
		Usage: "Your handy CLI typechart",
		Action: func(cCtx *cli.Context) error {
			resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", cCtx.Args().Get(0)))
			if err != nil {
				log.Fatal("How could you do this????", err)
			}

			var p Pokemon

			err = json.NewDecoder(resp.Body).Decode(&p)
			if err != nil {
				log.Fatal("You cant unmarshal:  ", err)
			}

			fmt.Printf("pokemon name: %s\n", p.Name)
			fmt.Printf("pokemon type(s): %s\n", p.Type[0].Type.Name)
			fmt.Printf("pokemon type(s): %s\n", p.Type[1].Type.Name)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("You are dumb.", err)
	}
}
