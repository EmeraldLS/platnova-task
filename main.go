package main

import (
	"flag"
	"log/slog"

	"github.com/emeraldls/platnova-task/internal/generator"
	"github.com/emeraldls/platnova-task/internal/rest"
	"github.com/emeraldls/platnova-task/internal/types"
	"github.com/emeraldls/platnova-task/internal/utils"
	"github.com/unidoc/unipdf/v3/creator"
)

func main() {
	apiKey := flag.String("api_key", "", "API required to generate PDF")
	file := flag.String("json_file", "account_statement.json", "JSON file to generate PDF")
	useApi := flag.Bool("use_api", false, "Use the API to generate the PDF, it needs .env file with the API key & PORT")

	flag.Parse()

	if *useApi {
		slog.Info("Using the API to generate the PDF")
		rest.SetupRoutes()
	}

	if *apiKey == "" {
		slog.Info("The APIKey is required, it can be passed as a flag")
		return
	}

	c := types.NewClient(creator.New(), *apiKey)
	stmt, err := utils.ReadJSONFile(*file)
	if err != nil {
		slog.Error("error has reading file", "err", err)
		return
	}

	fn, err := generator.GenerateAccountStatementPDF(*c, *stmt)
	if err != nil {
		slog.Error("error generating document", "err", err)
		return
	}

	slog.Info("Statement Generated: ", "filename", fn+".pdf")

}