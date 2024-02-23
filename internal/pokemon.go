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
