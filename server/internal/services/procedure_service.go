package services

import (
	"tech-quest/internal/domain/models"
	"tech-quest/internal/repository"
	"tech-quest/pkg/errors"
)

type ProcedureService struct {
	repo *repository.ProcedureRepository
}

func NewProcedureService(repo *repository.ProcedureRepository) *ProcedureService {
	return &ProcedureService{repo: repo}
}

func (s *ProcedureService) GetAll() ([]models.Procedure, error) {
	procedures, err := s.repo.GetAll()
	if err != nil {
		return nil, errors.NewError(
			500,
			errors.ErrorDetail{
				Code:   errors.ServerErrorCode,
				Detail: "failed to get procedures: " + err.Error(),
			},
		)
	}
	return procedures, nil
}

func (s *ProcedureService) GetByID(id int) (*models.Procedure, error) {
	procedure, err := s.repo.GetByID(id)
	if err != nil {
		if err == errors.ErrNotFound {
			return nil, errors.NewError(
				404,
				errors.ErrorDetail{
					Code:   errors.NotFoundCode,
					Detail: "procedure not found",
				},
			)
		}
		return nil, errors.NewError(
			500,
			errors.ErrorDetail{
				Code:   errors.ServerErrorCode,
				Detail: "failed to get procedure: " + err.Error(),
			},
		)
	}
	return procedure, nil
}

func (s *ProcedureService) GetByType(procedureType string) ([]models.Procedure, error) {
	procedures, err := s.repo.GetByType(procedureType)
	if err != nil {
		return nil, errors.NewError(
			500,
			errors.ErrorDetail{
				Code:   errors.ServerErrorCode,
				Detail: "failed to get procedures by type: " + err.Error(),
			},
		)
	}
	return procedures, nil
}

func (s *ProcedureService) Create(procedure *models.Procedure) error {
	if procedure.Title == "" {
		return errors.NewError(
			400,
			errors.ErrorDetail{
				Code:   errors.ValidationErrorCode,
				Detail: "title is required",
				Attr:   "title",
			},
		)
	}
	if procedure.Type == "" {
		return errors.NewError(
			400,
			errors.ErrorDetail{
				Code:   errors.ValidationErrorCode,
				Detail: "type is required",
				Attr:   "type",
			},
		)
	}
	err := s.repo.Create(procedure)
	if err != nil {
		return errors.NewError(
			500,
			errors.ErrorDetail{
				Code:   errors.ServerErrorCode,
				Detail: "failed to create procedure: " + err.Error(),
			},
		)
	}
	return nil
}

func (s *ProcedureService) Update(procedure *models.Procedure) error {
	if procedure.ID == 0 {
		return errors.NewError(
			400,
			errors.ErrorDetail{
				Code:   errors.ValidationErrorCode,
				Detail: "id is required",
				Attr:   "id",
			},
		)
	}
	if procedure.Title == "" {
		return errors.NewError(
			400,
			errors.ErrorDetail{
				Code:   errors.ValidationErrorCode,
				Detail: "title is required",
				Attr:   "title",
			},
		)
	}
	err := s.repo.Update(procedure)
	if err != nil {
		if err == errors.ErrNotFound {
			return errors.NewError(
				404,
				errors.ErrorDetail{
					Code:   errors.NotFoundCode,
					Detail: "procedure not found",
				},
			)
		}
		return errors.NewError(
			500,
			errors.ErrorDetail{
				Code:   errors.ServerErrorCode,
				Detail: "failed to update procedure: " + err.Error(),
			},
		)
	}
	return nil
}

func (s *ProcedureService) Delete(id int) error {
	if id == 0 {
		return errors.NewError(
			400,
			errors.ErrorDetail{
				Code:   errors.ValidationErrorCode,
				Detail: "id is required",
				Attr:   "id",
			},
		)
	}

	err := s.repo.Delete(id)
	if err != nil {
		if err == errors.ErrNotFound {
			return errors.NewError(
				404,
				errors.ErrorDetail{
					Code:   errors.NotFoundCode,
					Detail: "procedure not found",
				},
			)
		}
		return errors.NewError(
			500,
			errors.ErrorDetail{
				Code:   errors.ServerErrorCode,
				Detail: "failed to delete procedure: " + err.Error(),
			},
		)
	}
	return nil
}
