package parsereader

import (
	"io"
	"unicode/utf8"
)

type Reader interface {
	io.RuneReader
	io.Seeker
}

type Wrapper struct {
	io.ReadSeeker
}

func (r2 Wrapper) ReadRune() (r rune, size int, err error) {
	var buffer [utf8.UTFMax]byte
	i := 0

	for !utf8.FullRune(buffer[:i]) || i == 0 {
		_, err = r2.Read(buffer[i : i+1])
		if err != nil {
			return
		}
		i++
	}

	r, size = utf8.DecodeRune(buffer[:i])
	return
}

func FromReadSeeker(seeker io.ReadSeeker) Reader {
	return Wrapper{seeker}
}
