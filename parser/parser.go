package parser

import (
	"github.com/timbeurskens/goparselib"
	"io"
	"os"
	"strings"

	"github.com/timbeurskens/goparselib/parsereader"
)

func ParseString(str string, language goparselib.Symbol) (goparselib.Node, error) {
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

func ParseFile(filename string, language goparselib.Symbol) (goparselib.Node, error) {
	file, err := os.Open(filename)
	if err != nil {
		return goparselib.Node{}, err
	}

	defer file.Close()

	return ParseInput(parsereader.FromReadSeeker(file), language)
}

func ParseInput(reader parsereader.Reader, language goparselib.Symbol) (goparselib.Node, error) {
	return parse(reader, 0, language)
}

func parse(reader parsereader.Reader, start int64, language goparselib.Symbol) (goparselib.Node, error) {
	switch language.(type) {
	case goparselib.Union:
		l := language.(goparselib.Union)
		found, best, firstErr := false, goparselib.Node{}, error(nil)

		for i := range l {
			if child, err := parse(reader, start, l[i]); err == nil {
				found = true
				// longest match is accepted
				if child.Size > best.Size {
					best = child
				}
			} else if _, ok := err.(SyntaxError); !ok {
				// not a syntax error, so immediately fail
				return goparselib.Node{}, err
			} else if firstErr == nil {
				firstErr = err
			}
		}

		if !found {
			return goparselib.Node{}, SyntaxError{
				At:       start,
				Expected: language,
				Got:      "",
				Err:      firstErr,
			}
		}

		return goparselib.Node{
			Start:    best.Start,
			Size:     best.Size,
			Contents: "",
			Type:     language,
			Children: []goparselib.Node{best},
		}, nil
	case goparselib.Concat:
		l := language.(goparselib.Concat)

		pos := int64(0)
		children := make([]goparselib.Node, 0)

		for i := range l {
			if child, err := parse(reader, start+pos, l[i]); err != nil {
				return goparselib.Node{}, SyntaxError{
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

		return goparselib.Node{
			Start:    start,
			Size:     pos,
			Children: children,
			Type:     language,
		}, nil
	case goparselib.Terminal:
		t := language.(goparselib.Terminal)

		// make sure the reader is positioned at start
		_, err := reader.Seek(start, io.SeekStart)
		if err != nil {
			return goparselib.Node{}, InputError{
				At:  start,
				Err: err,
			}
		}

		loc := t.Reg.FindReaderIndex(reader)
		if loc == nil || loc[0] != 0 {

			// find failing rune
			reader.Seek(start, io.SeekStart)
			r, _, _ := reader.ReadRune()

			return goparselib.Node{}, SyntaxError{
				At:       start,
				Expected: language,
				Got:      string([]rune{r}), // todo: add string
			}
		}

		// extract string contents
		str, err := loc2string(start, loc[1], reader)
		if err != nil {
			return goparselib.Node{}, InputError{
				At:  start,
				Err: err,
			}
		}

		return goparselib.Node{
			Start:    start,
			Size:     int64(loc[1]),
			Type:     language,
			Contents: str,
		}, nil
	case goparselib.Reference:
		return parse(reader, start, *language.(goparselib.Reference).R)
	case nil:
		// nil always matches with a zero length submatch
		return goparselib.Node{
			Start:    start,
			Size:     0,
			Contents: "",
			Type:     nil,
			Children: nil,
		}, nil
	default:
		return goparselib.Node{}, ParserError{language}
	}
}
