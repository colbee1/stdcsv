package filters

import (
	"bufio"
	"io"
)

type trailingRemover struct {
	input *bufio.Reader
	buf   []byte
	pos   int
	line  []byte
	comma rune
}

func SkipTrailingComma(fin io.Reader, comma rune) io.Reader {
	return &trailingRemover{
		input: bufio.NewReader(fin),
		buf:   make([]byte, 65*1024),
		comma: comma,
	}
}

func (tr *trailingRemover) Read(bs []byte) (int, error) {
	if tr.pos >= len(tr.line) {
		if err := tr.readNextLine(); err != nil {
			return 0, err
		}
	}

	n := copy(bs, tr.line[tr.pos:])
	tr.pos += n

	return n, nil
}

func (tr *trailingRemover) readNextLine() error {
	tr.line = tr.line[:0] // reset line
	tr.pos = 0
	for {
		line, isPrefix, err := tr.input.ReadLine()
		if err != nil {
			return err
		}
		tr.line = append(tr.line, line...)
		if !isPrefix {
			break
		}
	}

	if len(tr.line) == 0 {
		return nil
	}

	if tr.comma != 0 {
		runes := []rune(string(tr.line))
		if runes[len(runes)-1] == tr.comma {
			tr.line = []byte(string(runes[:len(runes)-1]))
		}
	}

	tr.line = append(tr.line, '\n')

	return nil
}
