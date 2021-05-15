package goparselib

// LoadFunction is the function type for the initial language plugin entry
type LoadFunction = func() Plugin

type Plugin interface {
	Root() Symbol
	Layout() []Symbol
}

func MakePlugin(root Symbol, layout []Symbol) Plugin {
	return PluginStruct{
		RootSymbol:    root,
		LayoutSymbols: layout,
	}
}

type PluginStruct struct {
	RootSymbol    Symbol
	LayoutSymbols []Symbol
}

func (p PluginStruct) Layout() []Symbol {
	return p.LayoutSymbols
}

func (p PluginStruct) Root() Symbol {
	return p.RootSymbol
}
