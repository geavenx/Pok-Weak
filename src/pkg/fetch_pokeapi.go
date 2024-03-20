package pokemon

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Api_data struct {
	Name string `json:"name"`
}

type Gen_data struct {
	Pokemon_species []NameAndUrl `json:"pokemon_species"`
}

// FetchAllPokeAPI() function fetch all pokemon from gen 1 to 9 and cache them locally
func FetchAllPokeAPI() {
	err := FetchGens(1, 9)
	if err != nil {
		fmt.Println("Error fetching generations!!!!!")
		panic(err)
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	file := fmt.Sprintf("%s/assets/cache/pokemons.txt", dir)

	f, err := os.Open(file)
	if err != nil {
		fmt.Println("gotcha")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	done := make(chan bool)

	fmt.Println("Fetching Pokemon API and caching locally...")

	for scanner.Scan() {
		go CachePokemonConc(scanner.Text(), done)
	}

	for scanner.Scan() {
		<-done
	}

	fmt.Println("Done fetching PokeAPI and caching locally!")
}

// CachePokemonConc FETCH POKEMON FROM API CONCURRENTLY and cache locally
func CachePokemonConc(pokemonName string, done chan bool) {

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)

		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}

		file, err := os.Create(fmt.Sprintf("%s/assets/cache/pokemons/%s.json", dir, pokemonName))
		if err != nil {
			panic(err)
		}
		defer file.Close()

		bufWriter := bufio.NewWriter(file)

		_, err = bufWriter.WriteString(bodyString)
		if err != nil {
			panic(err)
		}

		// Flushing the buffer to write data to the file
		err = bufWriter.Flush()
		if err != nil {
			panic(err)
		}

	}
}

// CachePokemon receive response body from PokeAPI and cache the response locally
// object 1 -> pokemon
// object 2 -> types
func CachePokemon(object int, name string, resBody []byte) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("im here now: ", err)
	}

	var objectName string

	if object == 0 {
		objectName = "pokemons"
	} else if object == 1 {
		objectName = "types"
	} else {
		objectName = "unknown"
	}
	if objectName == "unknown" {
		return
	}

	file, err := os.Create(fmt.Sprintf("%s/assets/cache/%s/%s.json", dir, objectName, name))
	if err != nil {
		fmt.Printf("error trying to create cache file: %s", err)
	}
	defer file.Close()

	_, err = file.Write(resBody)
	if err != nil {
		fmt.Printf("error trying to write api response to cache file: %s", err)
	}
	file.Sync()
}

// This function will fetch all Pokemon names from a range of generations
// and save them locally in a pokemons.txt file.
func FetchGens(startGen int, endGen int) (err error) {
	if startGen < 1 || endGen > 9 {
		fmt.Println("Invalid generation range. Please enter a range between 1 and 9.")
		return nil
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	file, err := os.Create(fmt.Sprintf("%s/assets/cache/pokemons.txt", dir))
	if err != nil {
		return err
	}
	defer file.Close()

	for i := startGen; i <= endGen; i++ {
		var gp *Gen_data
		url := fmt.Sprintf("https://pokeapi.co/api/v2/generation/%v", i)
		resp, err := http.Get(url)
		if err != nil {
			return err
		}

		err = json.NewDecoder(resp.Body).Decode(&gp)
		if err != nil {
			return err
		}

		for _, p := range gp.Pokemon_species {
			_, err = file.WriteString(fmt.Sprintf("%s\n", p.Name))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
