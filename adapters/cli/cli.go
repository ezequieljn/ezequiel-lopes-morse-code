package cli

import (
	"log/slog"

	"github.com/ezequieljn/morse-code/application"
)

type CLI struct {
	Service application.DecoderService
}

func NewCLI(service application.DecoderService) *CLI {
	return &CLI{Service: service}
}

func (cli *CLI) Run(decode string) error {
	decoded, err := cli.Service.Decode(decode)
	if err != nil {
		return err
	}
	slog.Info("Decoded", "value", decoded)
	return nil
}
