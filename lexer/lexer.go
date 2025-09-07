package lexer

type Lexer struct {
	input			string
	position		int
	readProsition	int
	ch				byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readProsition >= len(l.input) {
	} else {
		l.ch = l.input[l.readProsition]
	}
	l.position = l.readProsition
}