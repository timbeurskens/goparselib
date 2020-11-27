package ini

import (
	"errors"
	"reflect"
	"strings"

	"goparselib"
)

// Reader can parse ini files
type Reader struct {
	Model goparselib.Node
}

func New(input string) (Reader, error) {
	ok, tree := goparselib.Parse(input, root)
	if !ok {
		return Reader{}, errors.New("syntax error")
	}
	tree.Populate(input)
	ast, err := tree.Reduce(nil, open, close, eof, eq, goparselib.Blank, goparselib.EOL, goparselib.BlankOpt)
	if err != nil {
		return Reader{}, err
	}

	return Reader{
		Model: ast,
	}, nil
}

func (r Reader) Property(ident string) (string, error) {
	prop, err := r.FindFirst(property, func(r1 Reader) bool {
		_, err := r1.FindFirst(goparselib.Ident, func(r2 Reader) bool {
			return r2.Model.Contents == ident
		})
		return err == nil
	})

	if err != nil {
		return "", err
	}

	v, err := prop.FindFirst(value, func(r Reader) bool {
		return true
	})

	if err != nil {
		return "", err
	}

	writer := &strings.Builder{}
	err = v.Model.Output(writer)
	if err != nil {
		return "", err
	}

	return writer.String(), nil
}

func (r Reader) FindFirst(s goparselib.Symbol, condition func(r Reader) bool) (Reader, error) {
	if reflect.DeepEqual(r.Model.Type, s) && condition(r) {
		return r, nil
	} else {

		if r.Model.Children == nil {
			return Reader{}, errors.New("not found")
		}

		for i := range r.Model.Children {
			r2 := Reader{r.Model.Children[i]}
			if r3, err := r2.FindFirst(s, condition); err == nil {
				return r3, nil
			}
		}

		return Reader{}, errors.New("not found")
	}
}

func (r Reader) Section(ident string) (Reader, error) {
	return r.FindFirst(section, func(r1 Reader) bool {
		_, err := r1.FindFirst(goparselib.Ident, func(r2 Reader) bool {
			return r2.Model.Contents == ident
		})
		return err == nil
	})
}
