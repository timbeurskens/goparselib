package goparselib

import (
	"io"
	"strings"
)

func Parse(str string, language Symbol) (Node, error) {
	reader := strings.NewReader(str)
	return parseReader(reader, 0, reader.Size(), language)
}

func parseReader(reader *strings.Reader, start, end int64, language Symbol) (Node, error) {
	// make sure function exits with reader at start
	defer func() {
		_, err := reader.Seek(start, io.SeekStart)
		if err != nil {
			panic(err)
		}
	}()

	switch language.(type) {
	case Union:
		l := language.(Union)
		found, best, lastErr := false, Node{}, error(nil)

		for i := range l {
			if child, err := parseReader(reader, start, end, l[i]); err == nil {
				found = true
				// longest match is accepted
				if child.Size > best.Size {
					best = child
				}
			} else if _, ok := err.(SyntaxError); !ok {
				// not a syntax error, so immediately fail
				return Node{}, err
			} else {
				lastErr = err
			}

			_, err := reader.Seek(start, io.SeekStart)
			if err != nil {
				return Node{}, InputError{
					Start: start,
					End:   end,
					Err:   err,
				}
			}
		}

		if !found {
			return Node{}, SyntaxError{
				At:       start,
				Expected: language,
				Got:      "",
				Err:      lastErr,
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
			if child, err := parseReader(reader, start+pos, end, l[i]); err != nil {
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

			_, err := reader.Seek(start+pos, io.SeekStart)
			if err != nil {
				return Node{}, InputError{
					Start: start,
					End:   end,
					Err:   err,
				}
			}
		}

		if start+pos > end {
			return Node{}, InputError{
				Start: start,
				End:   end,
			}
		}

		// log.Println(language, "->", children)

		return Node{
			Start:    start,
			Size:     pos,
			Children: children,
			Type:     language,
		}, nil
	case Terminal:
		t := language.(Terminal)
		loc := t.Reg.FindReaderIndex(reader)
		if loc == nil || loc[0] != 0 {
			return Node{}, SyntaxError{
				At:       start,
				Expected: language,
				Got:      "", // todo: add string
			}
		}
		return Node{
			Start: start,
			Size:  int64(loc[1]),
			Type:  language,
		}, nil
	case Reference:
		return parseReader(reader, start, end, *language.(Reference).R)
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
