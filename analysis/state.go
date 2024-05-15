package analysis

type State struct {
    Document map[string]string
}

func NewState() State {
    return State{Document: map[string]string{}}
}

func (s *State) OpenDocument(uri, text string) {
    s.Document[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
    s.Document[uri] = text
}
