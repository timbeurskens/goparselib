package goparselib

import (
	"io"
	"os"
	"strings"

	"github.com/timbeurskens/goparselib/parsereader"
)

func ParseString(str string, language Symbol) (Node, error) {
	reader := strings.NewReader(str)
	return ParseInput(parsereader.FromReadSeeker(reader), language)
}

func loc2string(start int64, n int, reader parsereader.Reader) (str string, err error) {
	var r rune

	_, err = reader.Seek(start, io.SeekStart)
	if err != nil {
		return
	}

	builder := &strings.Builder{}
	builder.Grow(n)

	for i := 0; i < n; i++ {
		r, _, err = reader.ReadRune()
		if err != nil {
			return
		}
		_, err = builder.WriteRune(r)
		if err != nil {
			return
		}
	}

	str = builder.String()
	return
}

func ParseFile(filename string, language Symbol) (Node, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Node{}, err
	}

	defer file.Close()

	return ParseInput(parsereader.FromReadSeeker(file), language)
}

func ParseInput(reader parsereader.Reader, language Symbol) (Node, error) {
	return parse(reader, 0, language)
}

func parse(reader parsereader.Reader, start int64, language Symbol) (Node, error) {
	switch language.(type) {
	case Union:
		l := language.(Union)
		found, best, firstErr := false, Node{}, error(nil)

		for i := range l {
			if child, err := parse(reader, start, l[i]); err == nil {
				found = true
				// longest match is accepted
				if child.Size > best.Size {
					best = child
				}
			} else if _, ok := err.(SyntaxError); !ok {
				// not a syntax error, so immediately fail
				return Node{}, err
			} else if firstErr == nil {
				firstErr = err
			}
		}

		if !found {
			return Node{}, SyntaxError{
				At:       start,
				Expected: language,
				Got:      "",
				Err:      firstErr,
			}
		}

		return Node{
			Start:    best.Start,
			Size:     best.Size,
			Contents: "",
			Type:     language,
			Children: []Node{best},
		}, nil
	case Concat:
		l := language.(Concat)

		pos := int64(0)
		children := make([]Node, 0)

		for i := range l {
			if child, err := parse(reader, start+pos, l[i]); err != nil {
				return Node{}, SyntaxError{
					At:       start,
					Expected: language,
					Got:      "",
					Err:      err,
				}
			} else {
				pos += child.Size
				children = append(children, child)
			}
		}

		return Node{
			Start:    start,
			Size:     pos,
			Children: children,
			Type:     language,
		}, nil
	case Terminal:
		t := language.(Terminal)

		// make sure the reader is positioned at start
		_, err := reader.Seek(start, io.SeekStart)
		if err != nil {
			return Node{}, InputError{
				At:  start,
				Err: err,
			}
		}

		loc := t.Reg.FindReaderIndex(reader)
		if loc == nil || loc[0] != 0 {

			// find failing rune
			reader.Seek(start, io.SeekStart)
			r, _, _ := reader.ReadRune()

			return Node{}, SyntaxError{
				At:       start,
				Expected: language,
				Got:      string([]rune{r}), // todo: add string
			}
		}

		// extract string contents
		str, err := loc2string(start, loc[1], reader)
		if err != nil {
			return Node{}, InputError{
				At:  start,
				Err: err,
			}
		}

		return Node{
			Start:    start,
			Size:     int64(loc[1]),
			Type:     language,
			Contents: str,
		}, nil
	case Reference:
		return parse(reader, start, *language.(Reference).R)
	case nil:
		// nil always matches with a zero length submatch
		return Node{
			Start:    start,
			Size:     0,
			Contents: "",
			Type:     nil,
			Children: nil,
		}, nil
	default:
		return Node{}, ParserError{language}
	}
}
