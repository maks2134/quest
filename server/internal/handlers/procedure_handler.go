package handlers

import (
	"github.com/gofiber/fiber/v3"
	"strconv"
	"tech-quest/internal/domain/models"
	"tech-quest/internal/services"
	"tech-quest/pkg/errors"
)

type ProcedureHandler struct {
	service *services.ProcedureService
}

func NewProcedureHandler(service *services.ProcedureService) *ProcedureHandler {
	return &ProcedureHandler{service: service}
}

// GetAll возвращает все процедуры
// @Summary Получить все процедуры
// @Description Возвращает список всех процедур, отсортированных по sort_order
// @Tags procedures
// @Accept json
// @Produce json
// @Success 200 {array} models.Procedure
// @Router /procedures [get]
func (h *ProcedureHandler) GetAll(c fiber.Ctx) error {
	procedures, err := h.service.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(procedures)
}

// GetByID возвращает процедуру по ID
// @Summary Получить процедуру по ID
// @Description Возвращает процедуру по указанному ID
// @Tags procedures
// @Accept json
// @Produce json
// @Param id path int true "ID процедуры"
// @Success 200 {object} models.Procedure
// @Failure 404 {object} errors.Error
// @Router /procedures/{id} [get]
func (h *ProcedureHandler) GetByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errors.NewSimpleError(fiber.StatusBadRequest, "invalid id parameter")
	}
	procedure, err := h.service.GetByID(id)
	if err != nil {
		return err
	}
	return c.JSON(procedure)
}

// GetByType возвращает процедуры по типу
// @Summary Получить процедуры по типу
// @Description Возвращает список процедур указанного типа
// @Tags procedures
// @Accept json
// @Produce json
// @Param type path string true "Тип процедуры"
// @Success 200 {array} models.Procedure
// @Router /procedures/type/{type} [get]
func (h *ProcedureHandler) GetByType(c fiber.Ctx) error {
	procedureType := c.Params("type")
	if procedureType == "" {
		return errors.NewSimpleError(fiber.StatusBadRequest, "type parameter is required")
	}
	procedures, err := h.service.GetByType(procedureType)
	if err != nil {
		return err
	}
	return c.JSON(procedures)
}

// Create создает новую процедуру
// @Summary Создать новую процедуру
// @Description Создает новую процедуру с указанными данными
// @Tags procedures
// @Accept json
// @Produce json
// @Param procedure body models.Procedure true "Данные процедуры"
// @Success 201 {object} models.Procedure
// @Failure 400 {object} errors.Error
// @Router /procedures [post]
func (h *ProcedureHandler) Create(c fiber.Ctx) error {
	var procedure models.Procedure
	if err := c.Bind().Body(&procedure); err != nil {
		return errors.NewSimpleError(fiber.StatusBadRequest, "invalid request body: "+err.Error())
	}
	if err := h.service.Create(&procedure); err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(procedure)
}

// Update обновляет существующую процедуру
// @Summary Обновить процедуру
// @Description Обновляет существующую процедуру по ID
// @Tags procedures
// @Accept json
// @Produce json
// @Param id path int true "ID процедуры"
// @Param procedure body models.Procedure true "Обновленные данные процедуры"
// @Success 200 {object} models.Procedure
// @Failure 400 {object} errors.Error
// @Failure 404 {object} errors.Error
// @Router /procedures/{id} [put]
func (h *ProcedureHandler) Update(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errors.NewSimpleError(fiber.StatusBadRequest, "invalid id parameter")
	}
	var procedure models.Procedure
	if err := c.Bind().Body(&procedure); err != nil {
		return errors.NewSimpleError(fiber.StatusBadRequest, "invalid request body: "+err.Error())
	}
	procedure.ID = id
	if err := h.service.Update(&procedure); err != nil {
		return err
	}
	return c.JSON(procedure)
}

// Delete удаляет процедуру по ID
// @Summary Удалить процедуру
// @Description Удаляет процедуру по указанному ID
// @Tags procedures
// @Accept json
// @Produce json
// @Param id path int true "ID процедуры"
// @Success 204 "No Content"
// @Failure 404 {object} errors.Error
// @Router /procedures/{id} [delete]
func (h *ProcedureHandler) Delete(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errors.NewSimpleError(fiber.StatusBadRequest, "invalid id parameter")
	}
	if err := h.service.Delete(id); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
