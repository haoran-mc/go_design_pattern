package memento

type originator struct {
	state string
}

func (e *originator) setState(state string) {
	e.state = state
}

func (e *originator) getState() string {
	return e.state
}

func (e *originator) createMemento() *memento {
	return &memento{state: e.state}
}

func (e *originator) restorememento(m *memento) {
	e.state = m.getSavedState()
}
