package interpreter

type Environment struct {
	intVars map[string]int64
	parent  *Environment
}

func (e *Environment) getInt(name string) int64 {
	value, ok := e.intVars[name]
	if !ok {
		if e.parent != nil {
			return e.parent.getInt(name)
		} else {
			panic("Unknown variable")
		}
	}
	return value
}
func (e *Environment) setInt(name string, value int64) {
	e.intVars[name] = value
}

func NewEnvironment(parent *Environment) *Environment {
	return &Environment{
		intVars: make(map[string]int64),
		parent:  parent,
	}
}
