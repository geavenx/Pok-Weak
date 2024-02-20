package pokemon

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/urfave/cli/v2"
)

type PokemonInfo interface {
	Type() string
	Name() string
	GetCoverage() []Coverage
}

type Pokemon struct {
	PokemonName  string  `json:"name"`
	PokemonTypes []Types `json:"types"`
}

type Types struct {
	Slot int      `json:"slot"`
	Name TypeInfo `json:"type"`
}

type TypeInfo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func GetPokemon(cCtx *cli.Context) (*Pokemon, error) {
	resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", cCtx.Args().Get(0)))
	if err != nil {
		return nil, err
	}

	var p *Pokemon

	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Pokemon) Type() []string {
	var types []string
	for _, t := range p.PokemonTypes {
		types = append(types, t.Name.Name)
	}

	return types
}

func (p *Pokemon) Name() string {
	return p.PokemonName
}
