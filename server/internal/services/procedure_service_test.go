package services

import (
	stderrors "errors"
	"testing"

	"github.com/stretchr/testify/require"

	"tech-quest/internal/domain/models"
	"tech-quest/internal/services/mocks"
	appErrors "tech-quest/pkg/errors"
)

func TestProcedureService_GetAll(t *testing.T) {
	tests := []struct {
		name      string
		mockFn    func() ([]models.Procedure, error)
		wantErr   bool
		wantCount int
	}{
		{
			name: "success",
			mockFn: func() ([]models.Procedure, error) {
				return []models.Procedure{
					{ID: 1, Title: "A"},
					{ID: 2, Title: "B"},
				}, nil
			},
			wantCount: 2,
		},
		{
			name: "repository error",
			mockFn: func() ([]models.Procedure, error) {
				return nil, stderrors.New("db error")
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.ProcedureRepoMock{
				GetAllFn: tt.mockFn,
			}
			service := NewProcedureService(repo)
			res, err := service.GetAll()
			if tt.wantErr {
				require.Error(t, err)

				var appErr *appErrors.Error
				require.True(t, stderrors.As(err, &appErr))
				require.Equal(t, 500, appErr.StatusCode)

				return
			}
			require.NoError(t, err)
			require.Len(t, res, tt.wantCount)
		})
	}
}

func TestProcedureService_GetByID(t *testing.T) {
	tests := []struct {
		name       string
		repoErr    error
		wantStatus int
	}{
		{
			name: "success",
		},
		{
			name:       "not found",
			repoErr:    appErrors.ErrNotFound,
			wantStatus: 404,
		},
		{
			name:       "repository error",
			repoErr:    stderrors.New("db error"),
			wantStatus: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.ProcedureRepoMock{
				GetByIDFn: func(id int) (*models.Procedure, error) {
					if tt.repoErr != nil {
						return nil, tt.repoErr
					}
					return &models.Procedure{ID: id}, nil
				},
			}
			service := NewProcedureService(repo)
			res, err := service.GetByID(1)
			if tt.wantStatus != 0 {
				require.Error(t, err)
				var appErr *appErrors.Error
				require.True(t, stderrors.As(err, &appErr))
				require.Equal(t, tt.wantStatus, appErr.StatusCode)
				if tt.wantStatus == 404 {
					require.Equal(t, appErrors.NotFoundCode, appErr.ErrorDetail[0].Code)
				}
				return
			}
			require.NoError(t, err)
			require.NotNil(t, res)
			require.Equal(t, 1, res.ID)
		})
	}
}

func TestProcedureService_GetByType(t *testing.T) {
	repo := &mocks.ProcedureRepoMock{
		GetByTypeFn: func(t string) ([]models.Procedure, error) {
			return []models.Procedure{
				{ID: 1, Type: t},
			}, nil
		},
	}
	service := NewProcedureService(repo)
	res, err := service.GetByType("manual")
	require.NoError(t, err)
	require.Len(t, res, 1)
	require.Equal(t, "manual", res[0].Type)
}

func TestProcedureService_Create_Validation(t *testing.T) {
	service := NewProcedureService(&mocks.ProcedureRepoMock{})
	err := service.Create(&models.Procedure{
		Type: "manual",
	})
	require.Error(t, err)
	var appErr *appErrors.Error
	require.True(t, stderrors.As(err, &appErr))
	require.Equal(t, 400, appErr.StatusCode)
	require.Equal(t, appErrors.ValidationErrorCode, appErr.ErrorDetail[0].Code)
	require.Equal(t, "title", appErr.ErrorDetail[0].Attr)
}

func TestProcedureService_Create_OK(t *testing.T) {
	called := false
	repo := &mocks.ProcedureRepoMock{
		CreateFn: func(p *models.Procedure) error {
			called = true
			p.ID = 1
			return nil
		},
	}
	service := NewProcedureService(repo)
	err := service.Create(&models.Procedure{
		Title: "Test",
		Type:  "manual",
	})
	require.NoError(t, err)
	require.True(t, called)
}

func TestProcedureService_Update(t *testing.T) {
	tests := []struct {
		name       string
		procedure  models.Procedure
		repoErr    error
		wantStatus int
	}{
		{
			name: "missing id",
			procedure: models.Procedure{
				Title: "Test",
			},
			wantStatus: 400,
		},
		{
			name: "not found",
			procedure: models.Procedure{
				ID:    1,
				Title: "Test",
			},
			repoErr:    appErrors.ErrNotFound,
			wantStatus: 404,
		},
		{
			name: "success",
			procedure: models.Procedure{
				ID:    1,
				Title: "Test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.ProcedureRepoMock{
				UpdateFn: func(p *models.Procedure) error {
					return tt.repoErr
				},
			}
			service := NewProcedureService(repo)
			err := service.Update(&tt.procedure)
			if tt.wantStatus != 0 {
				require.Error(t, err)

				var appErr *appErrors.Error
				require.True(t, stderrors.As(err, &appErr))
				require.Equal(t, tt.wantStatus, appErr.StatusCode)

				return
			}

			require.NoError(t, err)
		})
	}
}

func TestProcedureService_Delete(t *testing.T) {
	tests := []struct {
		name       string
		id         int
		repoErr    error
		wantStatus int
	}{
		{
			name:       "invalid id",
			id:         0,
			wantStatus: 400,
		},
		{
			name:       "not found",
			id:         1,
			repoErr:    appErrors.ErrNotFound,
			wantStatus: 404,
		},
		{
			name: "success",
			id:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.ProcedureRepoMock{
				DeleteFn: func(id int) error {
					return tt.repoErr
				},
			}

			service := NewProcedureService(repo)

			err := service.Delete(tt.id)

			if tt.wantStatus != 0 {
				require.Error(t, err)

				var appErr *appErrors.Error
				require.True(t, stderrors.As(err, &appErr))
				require.Equal(t, tt.wantStatus, appErr.StatusCode)

				return
			}

			require.NoError(t, err)
		})
	}
}
