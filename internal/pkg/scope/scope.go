package scope

// Scope represents a scope in a context of some block
type Scope struct {
	Parent  *Scope
	Closure map[any]any
}

func (s *Scope) Put(key any, value any) {
	s.Closure[key] = value
}

func (s *Scope) Get(key any) (any, bool) {
	val, ok := s.Closure[key]
	return val, ok
}

func Push(currScope **Scope) {
	newScope := &Scope{
		Parent:  *currScope,
		Closure: make(map[any]any),
	}
	*currScope = newScope
}

func Pop(currScope **Scope) {
	// take the address of parent scope before nulling out an old one
	toReplaceWith := (*currScope).Parent
	*currScope = nil
	*currScope = toReplaceWith
}
