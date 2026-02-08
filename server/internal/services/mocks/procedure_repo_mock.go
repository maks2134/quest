package mocks

import (
	"tech-quest/internal/domain/models"
)

type ProcedureRepoMock struct {
	GetAllFn    func() ([]models.Procedure, error)
	GetByIDFn   func(int) (*models.Procedure, error)
	GetByTypeFn func(string) ([]models.Procedure, error)
	CreateFn    func(*models.Procedure) error
	UpdateFn    func(*models.Procedure) error
	DeleteFn    func(int) error
}

func (m *ProcedureRepoMock) GetAll() ([]models.Procedure, error) {
	return m.GetAllFn()
}

func (m *ProcedureRepoMock) GetByID(id int) (*models.Procedure, error) {
	return m.GetByIDFn(id)
}

func (m *ProcedureRepoMock) GetByType(t string) ([]models.Procedure, error) {
	return m.GetByTypeFn(t)
}

func (m *ProcedureRepoMock) Create(p *models.Procedure) error {
	return m.CreateFn(p)
}

func (m *ProcedureRepoMock) Update(p *models.Procedure) error {
	return m.UpdateFn(p)
}

func (m *ProcedureRepoMock) Delete(id int) error {
	return m.DeleteFn(id)
}
