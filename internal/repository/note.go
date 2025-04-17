package repository

import (
	"context"
	"math/big"
	"time"

	"github.com/Railssa1/quicknotes/internal/models"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	e "github.com/Railssa1/quicknotes/internal/error"
)

type NoteRepository interface {
	List() ([]models.Note, error)
	GetById(id int) (*models.Note, error)
	Create(title, content, color string) (*models.Note, error)
	Update(id int, title, content, color string) (*models.Note, error)
	Delete(id int) error
}

func NewRepository(dbpool *pgxpool.Pool) NoteRepository {
	return &noteRepository{
		db: dbpool,
	}
}

type noteRepository struct {
	db *pgxpool.Pool
}

func (nr *noteRepository) List() ([]models.Note, error) {
	var notes []models.Note
	rows, err := nr.db.Query(context.Background(), "SELECT * FROM notes")
	if err != nil {
		return notes, e.NewRepositoryError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var note models.Note
		err = rows.Scan(&note.Id, &note.Title, &note.Content, &note.Color, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			return notes, e.NewRepositoryError(err)
		}

		notes = append(notes, note)
	}
	return notes, nil
}

func (nr *noteRepository) GetById(id int) (*models.Note, error) {
	var note models.Note

	row := nr.db.QueryRow(context.Background(), `SELECT * FROM notes WHERE id = $1`, id)

	if err := row.Scan(&note.Id, &note.Title, &note.Content, &note.Color, &note.CreatedAt, &note.UpdatedAt); err != nil {
		return &note, e.NewRepositoryError(err)
	}

	return &note, nil
}

func (nr *noteRepository) Create(title, content, color string) (*models.Note, error) {
	var note models.Note
	note.Title = pgtype.Text{String: title, Valid: true}
	note.Content = pgtype.Text{String: content, Valid: true}
	note.Color = pgtype.Text{String: color, Valid: true}

	query := `INSERT INTO notes (title, content, color) 
			  VALUES ($1, $2, $3)
			  RETURNING id, created_at`
	row := nr.db.QueryRow(context.Background(), query, note.Title, note.Content, note.Color)
	if err := row.Scan(&note.Id, &note.CreatedAt); err != nil {
		return &note, e.NewRepositoryError(err)
	}

	return &note, nil
}

func (nr *noteRepository) Update(id int, title, content, color string) (*models.Note, error) {
	var note models.Note

	if len(title) > 0 {
		note.Title = pgtype.Text{String: title, Valid: true}
	}

	if len(content) > 0 {
		note.Content = pgtype.Text{String: content, Valid: true}
	}

	if len(color) > 0 {
		note.Color = pgtype.Text{String: color, Valid: true}
	}

	note.Id = pgtype.Numeric{Int: big.NewInt(int64(id)), Valid: true}
	note.UpdatedAt = pgtype.Date{Time: time.Now(), Valid: true}

	query := `UPDATE notes set title = COALESCE($1, title), content = COALESCE($2, content), color = COALESCE($3, content), updated_at = $4 where id = $5`

	_, err := nr.db.Exec(context.Background(), query, note.Title, note.Content, note.Color, note.UpdatedAt, id)
	if err != nil {
		return &note, e.NewRepositoryError(err)
	}

	return &note, nil
}

func (nr *noteRepository) Delete(id int) error {
	_, err := nr.db.Exec(context.Background(), "DELETE FROM notes WHERE id = $1", id)

	if err != nil {
		return e.NewRepositoryError(err)
	}
	return nil
}
