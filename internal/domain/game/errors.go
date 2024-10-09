package game

type ErrLetterAlreadyExists struct{}

func (ErrLetterAlreadyExists) Error() string {
	return "letter already exists"
}

type ErrLetterNotFound struct{}

func (ErrLetterNotFound) Error() string {
	return "letter not found"
}
