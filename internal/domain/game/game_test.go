package game_test

import (
	_ "embed"
	"os"
	"path/filepath"
	"testing"

	"github.com/iTukaev/backand_academy_example/configs"
	"github.com/iTukaev/backand_academy_example/internal/domain/game"
	"github.com/iTukaev/backand_academy_example/internal/domain/word"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/test_config.yaml
var testConfig []byte

func TestGame_IsGameOver(t *testing.T) {
	dir := t.TempDir()
	configFile := filepath.Join(dir, "test_config.yaml")

	err := os.WriteFile(configFile, testConfig, 0644)
	require.NoError(t, err)

	cfg, err := configs.Init(configFile)
	require.NoError(t, err)

	aWord := word.New()
	err = aWord.Build(cfg)
	require.NoError(t, err)

	testCases := []struct {
		name  string
		game  *game.Game
		steps func(*game.Game)

		isGameOver bool
	}{
		{
			name: "Game is not over",
			game: game.NewGame(aWord),
			steps: func(g *game.Game) {
				g.SetWord()
			},

			isGameOver: false,
		},
		{
			name: "Game is over with fail",
			game: game.NewGame(aWord),
			steps: func(g *game.Game) {
				g.SetWord()
				_ = g.GuessLetter('w')
				_ = g.GuessLetter('q')
				_ = g.GuessLetter('t')
				_ = g.GuessLetter('u')
				_ = g.GuessLetter('i')
				_ = g.GuessLetter('o')
				_ = g.GuessLetter('p')
				_ = g.GuessLetter('k')
			},

			isGameOver: true,
		},
		{
			name: "Game is over with win",
			game: game.NewGame(aWord),
			steps: func(g *game.Game) {
				g.SetWord()
				_ = g.GuessLetter('y')
				_ = g.GuessLetter('e')
				_ = g.GuessLetter('a')
				_ = g.GuessLetter('r')
			},

			isGameOver: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.steps(tc.game)

			require.Equal(t, tc.isGameOver, tc.game.IsGameOver())
		})
	}
}

func TestGame_GuessLetter(t *testing.T) {
	dir := t.TempDir()
	configFile := filepath.Join(dir, "test_config.yaml")

	err := os.WriteFile(configFile, testConfig, 0644)
	require.NoError(t, err)

	cfg, err := configs.Init(configFile)
	require.NoError(t, err)

	aWord := word.New()
	err = aWord.Build(cfg)
	require.NoError(t, err)

	testCases := []struct {
		name   string
		game   *game.Game
		steps  func(*game.Game)
		letter rune

		err error
	}{
		{
			name: "Correct letter",
			game: game.NewGame(aWord),
			steps: func(g *game.Game) {
				g.SetWord()
			},
			letter: 'y',

			err: nil,
		},
		{
			name: "Letter already exists",
			game: game.NewGame(aWord),
			steps: func(g *game.Game) {
				g.SetWord()
				_ = g.GuessLetter('y')
			},
			letter: 'y',

			err: &game.ErrLetterAlreadyExists{},
		},
		{
			name: "Letter not found",
			game: game.NewGame(aWord),
			steps: func(g *game.Game) {
				g.SetWord()
			},
			letter: 't',

			err: &game.ErrLetterNotFound{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.steps(tc.game)

			gotErr := tc.game.GuessLetter(tc.letter)

			if tc.err == nil {
				require.NoError(t, gotErr)
			} else {
				require.ErrorAs(t, gotErr, tc.err)
			}
		})
	}
}
