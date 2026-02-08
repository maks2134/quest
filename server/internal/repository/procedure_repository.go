package repository

import (
	"github.com/jmoiron/sqlx"
	"tech-quest/internal/domain/models"
	"tech-quest/pkg/errors"
)

type ProcedureRepository struct {
	db *sqlx.DB
}

func NewProcedureRepository(db *sqlx.DB) *ProcedureRepository {
	return &ProcedureRepository{db: db}
}

func (r *ProcedureRepository) GetAll() ([]models.Procedure, error) {
	var procedures []models.Procedure
	query := `
		SELECT id, title, type, content, sort_order, is_expanded, created_at, updated_at
		FROM procedures
		ORDER BY sort_order ASC
	`
	err := r.db.Select(&procedures, query)
	if err != nil {
		return nil, err
	}
	return procedures, nil
}

func (r *ProcedureRepository) GetByID(id int) (*models.Procedure, error) {
	var procedure models.Procedure
	query := `
		SELECT id, title, type, content, sort_order, is_expanded, created_at, updated_at
		FROM procedures
		WHERE id = $1
	`
	err := r.db.Get(&procedure, query, id)
	if err != nil {
		return nil, err
	}
	return &procedure, nil
}

func (r *ProcedureRepository) GetByType(procedureType string) ([]models.Procedure, error) {
	var procedures []models.Procedure
	query := `
		SELECT id, title, type, content, sort_order, is_expanded, created_at, updated_at
		FROM procedures
		WHERE type = $1
		ORDER BY sort_order ASC
	`
	err := r.db.Select(&procedures, query, procedureType)
	if err != nil {
		return nil, err
	}
	return procedures, nil
}

func (r *ProcedureRepository) Create(procedure *models.Procedure) error {
	query := `
		INSERT INTO procedures (title, type, content, sort_order, is_expanded)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`
	return r.db.QueryRow(
		query,
		procedure.Title,
		procedure.Type,
		procedure.Content,
		procedure.SortOrder,
		procedure.IsExpanded,
	).Scan(&procedure.ID, &procedure.CreatedAt, &procedure.UpdatedAt)
}

func (r *ProcedureRepository) Update(procedure *models.Procedure) error {
	query := `
		UPDATE procedures
		SET title = $1, type = $2, content = $3, sort_order = $4, is_expanded = $5, updated_at = CURRENT_TIMESTAMP
		WHERE id = $6
		RETURNING updated_at
	`
	return r.db.QueryRow(
		query,
		procedure.Title,
		procedure.Type,
		procedure.Content,
		procedure.SortOrder,
		procedure.IsExpanded,
		procedure.ID,
	).Scan(&procedure.UpdatedAt)
}

func (r *ProcedureRepository) Delete(id int) error {
	query := `DELETE FROM procedures WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.ErrNotFound
	}
	return nil
}
