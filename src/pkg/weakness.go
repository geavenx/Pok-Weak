package pokemon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/geavenx/pokeweak/src/utils"
)

type DamageRelations struct {
	Relations TypeRelations `json:"damage_relations"`
}

type TypeRelations struct {
	NoEffectOn   []NameAndUrl `json:"no_damage_to"`
	NotVeryEffOn []NameAndUrl `json:"half_damage_to"`
	SuperEffOn   []NameAndUrl `json:"double_damage_to"`
	ImmuneTo     []NameAndUrl `json:"no_damage_from"`
	NotVeryEffTo []NameAndUrl `json:"half_damage_from"`
	SuperEffTo   []NameAndUrl `json:"double_damage_from"`
}

func (p *Pokemon) GetDamageRelations() (*DamageRelations, error) {
	var D *DamageRelations

	for i := range p.PokemonTypes {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}
		// if cached file exists, use it

		fExists := utils.FileExists(fmt.Sprintf("%s/assets/cache/types/%s.json", dir, p.Name()))
		if fExists {
			fmt.Println("file exists at least")
			file := fmt.Sprintf("%s/assets/cache/types/%s.json", dir, p.Name())

			f, err := os.Open(file)
			if err != nil {
				fmt.Printf("error trying to open cache file: %s", err)
			}
			defer f.Close()

			err = json.NewDecoder(f).Decode(&D)
			if err != nil {
				return nil, err
			}

			// if not, fetch from API
		} else {
			resp, err := http.Get(p.PokemonTypes[i].Name.Url)
			if err != nil {
				return nil, err
			}
			fmt.Println("I am a request!")

			err = json.NewDecoder(resp.Body).Decode(&D)
			if err != nil {
				return nil, err
			}
			if resp.StatusCode == http.StatusOK {
				b, err := json.Marshal(&D)
				if err != nil {
					fmt.Printf("error on json.Marshal (line 65): %s", err)
				}
				CachePokemon(1, p.PokemonName, b)
			}
		}

	}
	return D, nil
}

// METHODS
//

func (d *DamageRelations) NoEffectOn() (s string) {
	s = "No effect on: "

	if len(d.Relations.NoEffectOn) == 0 {
		return "Has effect on all types!"
	}

	for i := range d.Relations.NoEffectOn {
		if i == len(d.Relations.NoEffectOn)-1 {
			s = s + d.Relations.NoEffectOn[i].Name
		} else {
			s = s + d.Relations.NoEffectOn[i].Name + ", "
		}
	}

	return s
}

func (d *DamageRelations) NotVeryEffOn() (s string) {
	s = "Not very effective on: "

	if len(d.Relations.NotVeryEffOn) == 0 {
		return
	}

	for i := range d.Relations.NotVeryEffOn {
		if i == len(d.Relations.NotVeryEffOn)-1 {
			s = s + d.Relations.NotVeryEffOn[i].Name
		} else {
			s = s + d.Relations.NotVeryEffOn[i].Name + ", "
		}
	}

	return s
}

func (d *DamageRelations) SuperEffOn() (s string) {
	s = "Super effective on: "

	if len(d.Relations.NotVeryEffOn) == 0 {
		return
	}

	for i := range d.Relations.SuperEffOn {
		if i == len(d.Relations.SuperEffOn)-1 {
			s = s + d.Relations.SuperEffOn[i].Name
		} else {
			s = s + d.Relations.SuperEffOn[i].Name + ", "
		}
	}

	return s
}

func (d *DamageRelations) ImmuneTo() (s string) {
	s = "Immune to: "

	if len(d.Relations.ImmuneTo) == 0 {
		return "Immune to nothing!"
	}

	for i := range d.Relations.ImmuneTo {
		if i == len(d.Relations.ImmuneTo)-1 {
			s = s + d.Relations.ImmuneTo[i].Name
		} else {
			s = s + d.Relations.ImmuneTo[i].Name + ", "
		}
	}

	return s
}

func (d *DamageRelations) NotVeryEffTo() (s string) {
	s = "Not very effective to: "

	if len(d.Relations.NotVeryEffTo) == 0 {
		return
	}

	for i := range d.Relations.NotVeryEffTo {
		if i == len(d.Relations.NotVeryEffTo)-1 {
			s = s + d.Relations.NotVeryEffTo[i].Name
		} else {
			s = s + d.Relations.NotVeryEffTo[i].Name + ", "
		}
	}

	return s
}

func (d *DamageRelations) SuperEffTo() (s string) {
	s = "Super effective to: "

	if len(d.Relations.SuperEffTo) == 0 {
		return
	}

	for i := range d.Relations.SuperEffTo {
		if i == len(d.Relations.SuperEffTo)-1 {
			s = s + d.Relations.SuperEffTo[i].Name
		} else {
			s = s + d.Relations.SuperEffTo[i].Name + ", "
		}
	}

	return s
}
