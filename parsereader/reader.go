package parsereader

import (
	"io"
	"unicode/utf8"
)

// Reader implements a io.RuneReader and io.Seeker which are needed for Parser objects
type Reader interface {
	io.RuneReader
	io.Seeker
}

// Wrapper implements a Reader by utilizing a io.ReadSeeker
type Wrapper struct {
	io.ReadSeeker
}

// ReadRune returns the next rune from the io.ReadSeeker
func (r2 Wrapper) ReadRune() (r rune, size int, err error) {
	// allocate the maximum space for a UTF8 encoded rune
	var buffer [utf8.UTFMax]byte
	i := 0

	// as long as we have not read any token, or the data is not a full UTF8 rune, do
	for !utf8.FullRune(buffer[:i]) || i == 0 {
		// read a single byte from the reader
		_, err = r2.Read(buffer[i : i+1])
		if err != nil {
			return
		}
		// increase the byte counter
		i++
	}

	// return the UTF8 decoded rune
	r, size = utf8.DecodeRune(buffer[:i])
	return
}

// FromReadSeeker returns a Reader object from an io.ReadSeeker object
func FromReadSeeker(seeker io.ReadSeeker) Reader {
	return Wrapper{seeker}
}
