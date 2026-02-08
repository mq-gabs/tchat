package reader

import (
	"bufio"
	"os"
)

type Reader struct {
	scanner *bufio.Scanner
}

func New() *Reader {
	scanner := bufio.NewScanner(os.Stdin)
	return &Reader{
		scanner: scanner,
	}
}

func (r *Reader) Read() (string, error) {
	if !r.scanner.Scan() {
		return "", errCannotScanInput
	}

	input := r.scanner.Text()

	return input, nil
}
