package cmd

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ezequieljn/ezequiel-lopes-morse-code/adapters/cli"
	adaptersHttp "github.com/ezequieljn/ezequiel-lopes-morse-code/adapters/http"
	"github.com/ezequieljn/ezequiel-lopes-morse-code/application"
)

func Run() error {
	mode := flag.String("mode", "cli", "input type")
	version := flag.String("version", "one", "decode version")
	decode := flag.String("decode", "", "Texto a ser codificado")
	space := flag.String("space", "   ", "Espaço entre as palavras")
	port := flag.String("port", "8080", "Porta do servidor")
	flag.Parse()
	decoderService, err := application.NewDecoderFactory(*version, *space)
	if err != nil {
		return err
	}
	if *mode == "cli" {
		slog.Info("Running CLI")
		cliApp := cli.NewCLI(decoderService)
		return cliApp.Run(*decode)
	}
	if *mode == "http" {
		slog.Info("Running HTTP")
		http.HandleFunc("POST /", adaptersHttp.DecodeHandler(*version, *space))
		slog.Info(fmt.Sprintf("Servidor HTTP rodando na porta :%s", *port))
		if err := http.ListenAndServe(fmt.Sprintf(":%s", *port), nil); err != nil {
			slog.Error(err.Error())
			return err
		}
	}
	slog.Warn("Invalid mode")
	return nil
}
