package goparselib

import (
	"io"
	"reflect"
	"strings"
)

func Parse(str string, language Symbol) (bool, Node) {
	reader := strings.NewReader(str)
	return parseReader(reader, 0, reader.Size(), language)
}

func parseReader(reader *strings.Reader, start, end int64, language Symbol) (bool, Node) {
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
		found, best := false, Node{}

		for i := range l {
			if ok, child := parseReader(reader, start, end, l[i]); ok {
				found = true
				// longest match is accepted
				if child.Size > best.Size {
					best = child
				}
			}

			_, err := reader.Seek(start, io.SeekStart)
			if err != nil {
				panic(err)
			}
		}

		// if found {
		// 	log.Println(language, "->", best.Type, best.Size)
		// }

		return found, Node{
			Start:    best.Start,
			Size:     best.Size,
			Contents: "",
			Type:     language,
			Children: []Node{best},
		}
	case Concat:
		l := language.(Concat)

		pos := int64(0)
		children := make([]Node, 0)

		for i := range l {
			if ok, child := parseReader(reader, start+pos, end, l[i]); !ok {
				return false, Node{}
			} else {
				pos += child.Size
				children = append(children, child)
			}

			_, err := reader.Seek(start+pos, io.SeekStart)
			if err != nil {
				panic(err)
			}
		}

		if start+pos > end {
			panic("bound exceeded")
		}

		// log.Println(language, "->", children)

		return true, Node{
			Start:    start,
			Size:     pos,
			Children: children,
			Type:     language,
		}
	case Terminal:
		t := language.(Terminal)
		loc := t.Reg.FindReaderIndex(reader)
		if loc == nil || loc[0] != 0 {
			return false, Node{}
		}
		return true, Node{
			Start: start,
			Size:  int64(loc[1]),
			Type:  language,
		}
	case Reference:
		return parseReader(reader, start, end, *language.(Reference).R)
	case nil:
		// nil always matches with a zero length submatch
		return true, Node{
			Start:    start,
			Size:     0,
			Contents: "",
			Type:     nil,
			Children: nil,
		}
	default:
		panic(reflect.TypeOf(language))
		return false, Node{}
	}
}
