package main

import (
	"log/slog"
	"os"

	"github.com/iTukaev/backand_academy_example/configs"
	"github.com/iTukaev/backand_academy_example/internal/application"
	"github.com/iTukaev/backand_academy_example/internal/domain/game"
	"github.com/iTukaev/backand_academy_example/internal/domain/word"
	"github.com/iTukaev/backand_academy_example/internal/infrastructure"
	"github.com/spf13/pflag"
)

func main() {
	configPath := pflag.StringP("config", "c", "", "path to the configuration file")

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	cfg, err := configs.Init(*configPath)
	if err != nil {
		logger.Error("Init config", err)
		os.Exit(1)
	}

	aWord := word.New()
	if err = aWord.Build(cfg); err != nil {
		logger.Error("", err.Error())
		os.Exit(1)
	}

	aGame := game.NewGame(aWord)

	io := infrastructure.New(os.Stdin, os.Stdout, logger)

	app := application.New(aGame, io)

	if err = app.Start(); err != nil {
		logger.Error("Start game", err)
		os.Exit(1)
	}
}
