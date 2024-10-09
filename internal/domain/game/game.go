package game

import (
	"strings"

	"github.com/iTukaev/backand_academy_example/internal/domain/word"
)

type Game struct {
	attempts    int
	maxAttempts int

	letters  map[rune]struct{}
	word     word.Word
	template []rune

	hangStates []string
}

func NewGame(word word.Word) *Game {
	g := &Game{
		letters:    make(map[rune]struct{}),
		word:       word,
		hangStates: hangmanStates,
	}

	g.maxAttempts = len(g.hangStates)

	return g
}

func (g *Game) GuessLetter(r rune) error {
	if _, ok := g.letters[r]; ok {
		return ErrLetterAlreadyExists{}
	}

	if strings.ContainsRune(g.word.Word, r) {
		g.letters[r] = struct{}{}
		g.setLetterToTemplate(r)

		return nil
	}

	g.attempts++

	return ErrLetterNotFound{}
}

func (g *Game) IsGameOver() bool {
	return g.attempts >= g.maxAttempts || string(g.template) == g.word.Word
}

func (g *Game) IsUserWon() bool {
	return string(g.template) == g.word.Word
}

func (g *Game) SetWord() {
	g.word.SetRandomWord()

	g.template = make([]rune, 0, len(g.word.Word))

	for range g.word.Word {
		g.template = append(g.template, '_')
	}
}

func (g *Game) HangState() string {
	return g.hangStates[g.attempts%g.maxAttempts]
}

func (g *Game) Hint() string {
	return g.word.Hint
}

func (g *Game) Template() string {
	return string(g.template)
}

func (g *Game) AttemptsLeft() int {
	return g.maxAttempts - g.attempts
}

func (g *Game) setLetterToTemplate(r rune) {
	for i, l := range g.word.Word {
		if l == r {
			g.template[i] = r
		}
	}
}
