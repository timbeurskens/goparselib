package ast

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/timbeurskens/goparselib"
	"github.com/timbeurskens/goparselib/examples/ini"
)

type StringProperty struct {
	PropertyName string `json:"name"`
	Value        string `json:"value"`
}

func (s StringProperty) Name() string {
	return s.PropertyName
}

func (s StringProperty) String() string {
	return fmt.Sprintf("%s=%s", s.PropertyName, s.Value)
}

type FloatProperty struct {
	PropertyName string  `json:"name"`
	Value        float64 `json:"value"`
}

func (f FloatProperty) Name() string {
	return f.PropertyName
}

func (f FloatProperty) String() string {
	return fmt.Sprintf("%s=[f]%f", f.PropertyName, f.Value)
}

type Property interface {
	Name() string
	fmt.Stringer
}

type File struct {
	Properties []Property            `json:"global"`
	Sections   map[string][]Property `json:"sections"`
}

func LoadPropertyName(node goparselib.Node) (string, error) {
	return node.Children[0].Contents, nil
}

func LoadFloatProperty(node goparselib.Node) (FloatProperty, error) {
	name, err := LoadPropertyName(node.Children[0])
	if err != nil {
		return FloatProperty{}, err
	}

	valueStr := node.Children[1].Contents
	value, err := strconv.ParseFloat(valueStr, 64)

	return FloatProperty{
		PropertyName: name,
		Value:        value,
	}, nil
}

func LoadStringProperty(node goparselib.Node) (StringProperty, error) {
	name, err := LoadPropertyName(node.Children[0])
	if err != nil {
		return StringProperty{}, err
	}

	value := node.Children[1].Contents

	return StringProperty{
		PropertyName: name,
		Value:        value,
	}, nil
}

func LoadProperty(node goparselib.Node) (Property, error) {
	if reflect.DeepEqual(node.Children[0].Type, ini.StringProperty) {
		return LoadStringProperty(node.Children[0])
	} else if reflect.DeepEqual(node.Children[0].Type, ini.FloatProperty) {
		return LoadFloatProperty(node.Children[0])
	} else {
		return nil, fmt.Errorf("node %s is neither a string nor float", node.Children[0].Type)
	}
}

func LoadSection(node goparselib.Node) (string, []Property, error) {
	name, err := LoadPropertyName(node.Children[0])
	if err != nil {
		return "", nil, err
	}

	properties, err := LoadProperties(node.Children[1])
	if err != nil {
		return "", nil, err
	}

	return name, properties, nil
}

func LoadSections(node goparselib.Node) (map[string][]Property, error) {
	nodes, err := node.CollapseType(ini.Section)
	if err != nil {
		return nil, err
	}

	sections := make(map[string][]Property)

	for i := range nodes {
		key, properties, err := LoadSection(nodes[i])
		if err != nil {
			return nil, err
		}
		sections[key] = properties
	}

	return sections, nil
}

func LoadProperties(node goparselib.Node) ([]Property, error) {
	nodes, err := node.CollapseType(ini.Property)
	if err != nil {
		return nil, err
	}

	properties := make([]Property, len(nodes))

	for i := range nodes {
		property, err := LoadProperty(nodes[i])
		if err != nil {
			return nil, err
		}
		properties[i] = property
	}

	return properties, nil
}

func LoadFile(node goparselib.Node) (File, error) {
	properties, err := LoadProperties(node.Children[0])
	if err != nil {
		return File{}, err
	}

	sections, err := LoadSections(node.Children[1])
	if err != nil {
		return File{}, err
	}

	return File{
		Properties: properties,
		Sections:   sections,
	}, nil
}
