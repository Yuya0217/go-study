package handler

import (
	openapi "go-layered-architecture-sample/generated"
	"go-layered-architecture-sample/internal/domain/entity"
	"go-layered-architecture-sample/internal/usecase"
	"go-layered-architecture-sample/internal/usecase/dtos"
	"go-layered-architecture-sample/pkg/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MedicineHandler struct {
	medicineUseCase usecase.MedicineUseCase
	logger          logger.Logger
}

func NewMedicineHandler(medicineUseCase usecase.MedicineUseCase, logger logger.Logger) *MedicineHandler {
	return &MedicineHandler{medicineUseCase, logger}
}

type MedicineIDParam struct {
	ID int `param:"id"`
}

type MedicineFindParams struct {
	MedicineName *string `query:"medicineName"`
	MedicineType *string `query:"medicineType"`
}

func (h *MedicineHandler) Find(c echo.Context) error {
	params := new(MedicineFindParams)

	if err := c.Bind(params); err != nil {
		return entity.NewAppError(entity.MedicineValidationFailed, nil)
	}

	medicines, err := h.medicineUseCase.Find(c.Request().Context(), dtos.MedicineFindParams{
		MedicineName: params.MedicineName,
		MedicineType: params.MedicineType,
	})

	if err != nil {
		return err
	}

	response := toMedicinesResponse(medicines)

	return c.JSON(http.StatusOK, response)
}

func (h *MedicineHandler) GetByID(c echo.Context) error {
	params := new(MedicineIDParam)

	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind parameters")
	}

	medicine, err := h.medicineUseCase.GetByID(c.Request().Context(), params.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := toMedicineResponse(medicine)

	return c.JSON(http.StatusOK, response)
}

func (h *MedicineHandler) Create(c echo.Context) error {
	req := new(openapi.MedicineRequest)

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind parameters")
	}

	medicine, err := h.medicineUseCase.Create(
		c.Request().Context(),
		dtos.Medicine{
			MedicineName: req.MedicineName,
			MedicineType: req.MedicineType,
		},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := toMedicineResponse(medicine)

	return c.JSON(http.StatusOK, response)
}

func (h *MedicineHandler) Update(c echo.Context) error {
	req := new(openapi.MedicineRequest)

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind parameters")
	}

	param := new(MedicineIDParam)

	if err := c.Bind(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind parameters")
	}

	medicine, err := h.medicineUseCase.Update(
		c.Request().Context(),
		dtos.Medicine{
			ID:           param.ID,
			MedicineName: req.MedicineName,
			MedicineType: req.MedicineType,
		},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := toMedicineResponse(medicine)

	return c.JSON(http.StatusOK, response)
}

func (h *MedicineHandler) Delete(c echo.Context) error {
	req := new(MedicineIDParam)

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind parameters")
	}

	err := h.medicineUseCase.Delete(c.Request().Context(), req.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
