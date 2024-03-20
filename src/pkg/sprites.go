package pokemon

import (
	"fmt"
	"os"
	"path/filepath"
)

func (p *Pokemon) PrintSprite() error {
	// Read a simple text files that use unicode characters and ANSI color codes and print it to the terminal

	folder, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}

	file := fmt.Sprintf("%s/assets/sprites/regular/%s", folder, p.PokemonName)

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	b := make([]byte, 20000)
	n, err := f.Read(b)
	fmt.Printf("%s", b[:n])

	return nil
}
