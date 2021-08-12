package ini

import (
	"reflect"
	"strings"

	"github.com/timbeurskens/goparselib"
)

// Reader can parse ini files
type Reader struct {
	Model goparselib.Node
}

func New(input string) (Reader, error) {
	tree, err := goparselib.ParseString(input, Root)
	if err != nil {
		return Reader{}, err
	}
	ast, err := tree.Reduce(nil, parOpen, parClose, eq, goparselib.Blank, goparselib.EOL, goparselib.BlankOpt)
	if err != nil {
		return Reader{}, err
	}

	return Reader{
		Model: ast,
	}, nil
}

func (r Reader) Property(ident string) (string, error) {
	prop, err := r.Model.Find(func(r1 goparselib.Node) bool {
		if !reflect.DeepEqual(r1.Type, Property) {
			return false
		}
		_, err := r1.Find(func(r2 goparselib.Node) bool {
			if !reflect.DeepEqual(r2.Type, goparselib.Ident) {
				return false
			}
			return r2.Contents == ident
		})
		return err == nil
	})

	if err != nil {
		return "", err
	}

	v, err := prop.Find(func(r goparselib.Node) bool {
		return reflect.DeepEqual(r.Type, goparselib.Float)
	})

	if err != nil {
		return "", err
	}

	writer := &strings.Builder{}
	err = v.Output(writer)
	if err != nil {
		return "", err
	}

	return writer.String(), nil
}

func (r Reader) Section(ident string) (Reader, error) {
	node, err := r.Model.Find(func(n goparselib.Node) bool {
		if !reflect.DeepEqual(n.Type, Section) {
			return false
		}
		_, err := n.Find(func(n2 goparselib.Node) bool {
			if !reflect.DeepEqual(n2.Type, goparselib.Ident) {
				return false
			}
			return n2.Contents == ident
		})
		return err == nil
	})

	return Reader{node}, err
}
