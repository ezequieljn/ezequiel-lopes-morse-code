package application

import (
	"errors"
	"log/slog"

	"github.com/ezequieljn/morse-code/pkg/morse"
)

type DecoderService interface {
	Decode(text string) (string, error)
}

func NewDecoderFactory(version, space string) (DecoderService, error) {
	if version == "one" {
		slog.Info("Code Morse", "version", version)
		return morse.NewMorseOne(space), nil
	}
	if version == "two" {
		slog.Info("Code Morse", "version", version)
		return morse.NewMorseTwo(space), nil
	}
	if version == "three" {
		slog.Info("Code Morse", "version", version)
		return morse.NewMorseThree(space), nil
	}
	return nil, errors.New("version invalid")
}
