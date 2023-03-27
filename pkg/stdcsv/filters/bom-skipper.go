package filters

import (
	"bufio"
	"io"
)

func BOMSkipper(fin io.Reader) (io.Reader, error) {
	input := bufio.NewReader(fin)
	r, _, err := input.ReadRune()
	if err != nil {
		return nil, err
	}
	if r != '\uFEFF' {
		input.UnreadRune() // Not a BOM, restore rune back
	}

	return input, nil
}
