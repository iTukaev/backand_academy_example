package word

import (
	"math/rand"
	"time"

	"github.com/iTukaev/backand_academy_example/configs"
)

type Word struct {
	Word       string
	Hint       string
	Dictionary []wordWithHint
}

type wordWithHint struct {
	Word string
	Hint string
}

func New() Word {
	return Word{}
}

func (w *Word) Build(cfg *configs.Config) error {
	for _, word := range cfg.Words {
		w.Dictionary = append(w.Dictionary, wordWithHint{Word: word.Word, Hint: word.Hint})
	}

	return nil
}

func (w *Word) SetRandomWord() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	word := w.Dictionary[rnd.Intn(len(w.Dictionary))]

	w.Word = word.Word
	w.Hint = word.Hint
}
