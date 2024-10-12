package application_test

import (
	"testing"

	"github.com/iTukaev/backand_academy_example/configs"
	"github.com/iTukaev/backand_academy_example/internal/application"
	"github.com/iTukaev/backand_academy_example/internal/application/mocks"
	"github.com/iTukaev/backand_academy_example/internal/domain/game"
	"github.com/iTukaev/backand_academy_example/internal/domain/word"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestApp_Start(t *testing.T) {
	setupTest := func() (*mocks.MockioAdapter, *application.App) {
		cfg := &configs.Config{Words: []configs.Word{
			{Word: "mock", Hint: "mockery test stub"},
		}}

		aWord := word.New()

		err := aWord.Build(cfg)
		require.NoError(t, err)

		aGame := game.NewGame(aWord)

		mockIO := &mocks.MockioAdapter{}

		return mockIO, application.New(aGame, mockIO)
	}

	t.Parallel()

	t.Run("When player enters correct letters, then hi wins game", func(t *testing.T) {
		t.Parallel()

		mockIO, app := setupTest()

		mockIO.On("Output", mock.AnythingOfType("string")).Times(18)
		mockIO.On("Output", "Congratulations! You won!").Times(1)
		mockIO.On("Output", "Your word is: mock").Times(1)

		mockIO.On("Input").Return("m", nil).Once()
		mockIO.On("Input").Return("o", nil).Once()
		mockIO.On("Input").Return("c", nil).Once()
		mockIO.On("Input").Return("k", nil).Once()

		err := app.Start()
		require.NoError(t, err)
	})

	t.Run("When player enters incorrect letters, then hi looses game", func(t *testing.T) {
		t.Parallel()

		mockIO, app := setupTest()

		// Здесь было запарно, но я не обломался и посчитал сколько раз будет вызван Output.
		mockIO.On("Output", mock.AnythingOfType("string")).Times(40)
		mockIO.On("Output", "You are looser! But the deceased doesn't care anymore.").Times(1)

		mockIO.On("Input").Return("т", nil).Once()
		mockIO.On("Input").Return("о", nil).Once()
		mockIO.On("Input").Return("ч", nil).Once()
		mockIO.On("Input").Return("н", nil).Once()
		mockIO.On("Input").Return("о", nil).Once()
		mockIO.On("Input").Return("н", nil).Once()
		mockIO.On("Input").Return("е", nil).Once()
		mockIO.On("Input").Return(")", nil).Once()

		err := app.Start()
		require.NoError(t, err)
	})

	// Здесь можно добавить тесты с введённой повторно буквой и т.д.
}
