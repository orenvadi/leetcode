package main

func isNumber(s string) bool {
	var state State = InitialState{}
	for _, r := range s {
		state = state.Next(r)
	}
	return state.isEndState()
}

func isDot(r rune) bool {
	return r == '.'
}

func isSign(r rune) bool {
	switch r {
	case '+', '-':
		return true
	default:
		return false
	}
}

func isDigit(r rune) bool {
	switch r {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}

func isExp(r rune) bool {
	switch r {
	case 'e', 'E':
		return true
	default:
		return false
	}
}

type State interface {
	isEndState() bool
	Next(r rune) State
}

// InvalidState
type InvalidState struct{}

func (i InvalidState) isEndState() bool { return false }

func (i InvalidState) Next(r rune) State {
	return i
}

// InitialState
type InitialState struct{}

func (i InitialState) isEndState() bool { return false }

func (i InitialState) Next(r rune) State {
	switch {
	case isSign(r):
		return SignState{}
	case isDot(r):
		return DotFloatState{}
	case isDigit(r):
		return IntegerState{}
	default:
		return InvalidState{}
	}
}

// SignState
type SignState struct{}

func (i SignState) isEndState() bool { return false }

func (i SignState) Next(r rune) State {
	switch {
	case isDot(r):
		return DotFloatState{}
	case isDigit(r):
		return IntegerState{}
	default:
		return InvalidState{}
	}
}

// IntegerState
type IntegerState struct{}

func (i IntegerState) isEndState() bool { return true }

func (i IntegerState) Next(r rune) State {
	switch {
	case isDot(r):
		return FloatState{}
	case isDigit(r):
		return IntegerState{}
	case isExp(r):
		return ExpSignState{}
	default:
		return InvalidState{}
	}
}

// DotFloatState
type DotFloatState struct{}

func (i DotFloatState) isEndState() bool { return false }

func (i DotFloatState) Next(r rune) State {
	switch {
	case isDigit(r):
		return FloatState{}
	default:
		return InvalidState{}
	}
}

// FloatState
type FloatState struct{}

func (i FloatState) isEndState() bool { return true }

func (i FloatState) Next(r rune) State {
	switch {
	case isDigit(r):
		return FloatState{}
	case isExp(r):
		return ExpSignState{}
	default:
		return InvalidState{}
	}
}

// ExpSignState
type ExpSignState struct{}

func (i ExpSignState) isEndState() bool { return false }

func (i ExpSignState) Next(r rune) State {
	switch {
	case isSign(r):
		return SignedExpState{}
	case isDigit(r):
		return IntExpState{}
	default:
		return InvalidState{}
	}
}

// SignedExpState
type SignedExpState struct{}

func (i SignedExpState) isEndState() bool { return false }

func (i SignedExpState) Next(r rune) State {
	switch {
	case isDigit(r):
		return IntExpState{}
	default:
		return InvalidState{}
	}
}

// IntExpState
type IntExpState struct{}

func (i IntExpState) isEndState() bool { return true }

func (i IntExpState) Next(r rune) State {
	switch {
	case isDigit(r):
		return IntExpState{}
	default:
		return InvalidState{}
	}
}
