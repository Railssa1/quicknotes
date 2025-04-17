package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"

	e "github.com/Railssa1/quicknotes/internal/error"
)

type noteHandler struct{}

func NewNoteHandler() *noteHandler {
	return &noteHandler{}
}

// Método para listar notas
func (nh *noteHandler) ListNotes(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		return errors.New("página não encontrada")
	}

	files := []string{
		"views/templates/base.html",
		"views/templates/pages/home.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		return errors.New("erro ao relizar parser dos arquivos")
	}
	t.Name()

	return t.ExecuteTemplate(w, "base", nil)
}

// Método para recuperar uma nota
func (nh *noteHandler) NoteView(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("id")
	if id == "" {
		return e.WithStatus(errors.New("anotação não encontrada"), http.StatusBadRequest)
	}
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/note-view.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		return errors.New("erro ao relizar parser dos arquivos")
	}

	return t.ExecuteTemplate(w, "base", id)
}

// Método para criar uma nova
func (nh *noteHandler) CreateNotes(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		w.Header().Set("allow", http.MethodPost)

		return errors.New("método não permitido")
	}
	fmt.Fprint(w, "Criando notas")
	return nil
}

// Método de formulário de criação de uma nova nota
func (nh *noteHandler) NoteNew(w http.ResponseWriter, r *http.Request) error {
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/note-new.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		return errors.New("erro ao relizar parser dos arquivos")
	}

	return t.ExecuteTemplate(w, "base", nil)
}
