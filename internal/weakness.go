package pokemon

import (
	"encoding/json"
	"net/http"
)

type Coverage struct {
	ReceivedDmg []TypeCoverage
	DealtDmg    []TypeCoverage
}

type TypeCoverage struct {
	TypeName      string
	DamageCounter int8
}

type TypeRelations struct {
	// TODO
}

func (p *Pokemon) GetCoverage() (c Coverage, err error) {
	reqUrl := p.PokemonTypes[0].Name.Url
	resp, err := http.Get(reqUrl)
	if err != nil {
		return c, err
	}

	//TODO

	//err = json.NewDecoder(resp.Body).Decode()
	//if err != nil {
	//	return c, err
	//}

	return c, nil
}
