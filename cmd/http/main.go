package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/Railssa1/quicknotes/internal/handlers"
	"github.com/Railssa1/quicknotes/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config := loadConfig()
	mux := http.NewServeMux()
	logger := newLogger(os.Stderr, slog.LevelInfo)
	slog.SetDefault(logger)

	dbpool, err := pgxpool.New(context.Background(), config.DBConnURL)
	if err != nil {
		slog.Error("Erro ao se conectar ao banco de dados")
		os.Exit(1)
	}

	slog.Info("Conexão com o banco de dados feita com sucesso")
	defer dbpool.Close()

	repository := repository.NewRepository(dbpool)
	noteHandler := handlers.NewNoteHandler(repository)

	staticHandler := http.FileServer(http.Dir("views/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", staticHandler))

	mux.Handle("/", handlers.HandlerWithError(noteHandler.ListNotes))
	mux.Handle("/notes/view", handlers.HandlerWithError(noteHandler.NoteView))
	mux.Handle("/notes/create", handlers.HandlerWithError(noteHandler.CreateNotes))
	mux.Handle("/notes/new", handlers.HandlerWithError(noteHandler.NoteNew))

	slog.Info(fmt.Sprintf("servidor rodando na porta %s\n", config.ServerPort))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), mux); err != nil {
		slog.Info(fmt.Sprintf("Erro ao servir servidor:%s", err))
	}

}
