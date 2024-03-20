package pokemon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/geavenx/pokeweak/src/utils"
	"github.com/urfave/cli/v2"
)

type PokemonInfo interface {
	Type() string
	Name() string
	GetDamageRealations()
	Attacking()
	Defending()
}

type Pokemon struct {
	PokemonName  string     `json:"name"`
	PokemonTypes []TypeList `json:"types"`
}

type TypeList struct {
	Slot int        `json:"slot"`
	Name NameAndUrl `json:"type"`
}

func GetPokemon(cCtx *cli.Context) (*Pokemon, error) {
	var p *Pokemon

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	// if cached file exists, use it
	fExists := utils.FileExists(fmt.Sprintf("%s/assets/cache/pokemons/%s.json", dir, cCtx.Args().Get(0)))
	if fExists {
		file := fmt.Sprintf("%s/assets/cache/pokemons/%s.json", dir, cCtx.Args().Get(0))

		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("error trying to open cache file: %s", err)
		}
		defer f.Close()

		err = json.NewDecoder(f).Decode(&p)
		if err != nil {
			return nil, err
		}

		// if not, fetch from API
	} else {
		resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", cCtx.Args().Get(0)))
		if err != nil {
			return nil, err
		}

		err = json.NewDecoder(resp.Body).Decode(&p)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode == http.StatusOK {
			b, err := json.Marshal(*p)
			if err != nil {
				fmt.Printf("error on json.Marshal (line 75): %s", err)
			}
			CachePokemon(0, p.Name(), b)
		}

	}

	err = p.PrintSprite()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Pokemon) Type() string {
	s := "Type(s): "

	for i, t := range p.PokemonTypes {
		if i == len(p.PokemonTypes)-1 {
			s = s + t.Name.Name
		} else {
			s = s + t.Name.Name + ", "
		}
	}
	return s
}

func (p *Pokemon) Name() string {
	return p.PokemonName
}
