package application

import (
	"errors"
	"fmt"
	"strings"

	"github.com/iTukaev/backand_academy_example/internal/domain/game"
)

type ioAdapter interface {
	Input() (string, error)
	Output(content string)
}

type App struct {
	game *game.Game
	io   ioAdapter

	isHintNeeded bool
}

func New(game *game.Game, io ioAdapter) *App {
	return &App{game: game, io: io}
}

func (a *App) Start() error {
	a.game.SetWord()

	a.greetings()

	for !a.game.IsGameOver() {
		a.newRound()

		r, err := a.handleUserInput()
		if err != nil {
			return fmt.Errorf("handle user input: %w", err)
		}

		err = a.game.GuessLetter(r)
		switch {
		case err == nil:
			a.successGuess()
		case errors.As(err, &game.ErrLetterAlreadyExists{}):
			a.repeatedLetterGuess()
		default:
			a.failGuess()
		}
	}

	if a.game.IsUserWon() {
		a.io.Output("Congratulations! You won!")
		a.io.Output("Your word is: " + a.game.Template())
	} else {
		a.io.Output("You are looser! But the deceased doesn't care anymore.")
	}

	return nil
}

func (a *App) greetings() {
	a.io.Output("Hello user!")
	a.io.Output("Game started.")
}

func (a *App) newRound() {
	a.io.Output("Your word is: " + a.game.Template())

	if a.isHintNeeded {
		a.io.Output("Hint: " + a.game.Hint())
	} else {
		a.io.Output("If you need a hint enter '?' at any time.")
	}
}

func (a *App) successGuess() {
	a.io.Output("Great! You guessed the letter correctly.")
}

func (a *App) repeatedLetterGuess() {
	a.io.Output("Stupid guy, this letter was used.")
}

func (a *App) failGuess() {
	if !a.game.IsGameOver() {
		a.io.Output("Your death is getting closer...")
		a.io.Output(a.game.HangState())
	}
}

func (a *App) handleUserInput() (rune, error) {
	for {
		a.io.Output("Your letter:")

		s, err := a.io.Input()
		if err != nil {
			return 0, fmt.Errorf("input error: %w", err)
		}

		r := []rune(strings.ToLower(s))
		if len(r) == 0 {
			a.io.Output("Stupid user, one more time.")
			continue
		}

		if r[0] == '?' {
			a.isHintNeeded = true
		}

		// Здесь надо сделать валидацию ввода буква/цифра/говно, но мне лень :)
		return r[0], nil
	}
}
