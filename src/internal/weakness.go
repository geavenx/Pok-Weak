package pokemon

import (
	"encoding/json"
	"net/http"
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
	var d *DamageRelations

	for i := range p.PokemonTypes {
		resp, err := http.Get(p.PokemonTypes[i].Name.Url)
		if err != nil {
			return nil, err
		}

		err = json.NewDecoder(resp.Body).Decode(&d)
		if err != nil {
			return nil, err
		}

		resp, err = http.Get(p.PokemonTypes[0].Name.Url)
		if err != nil {
			return nil, err
		}

		err = json.NewDecoder(resp.Body).Decode(&D)
		if err != nil {
			return nil, err
		}

	}
	return D, nil
}

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
