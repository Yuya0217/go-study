package handler

import (
	openapi "go-layered-architecture-sample/generated"
	"go-layered-architecture-sample/internal/usecase/dtos"
)

func toMedicineResponse(medicine *dtos.Medicine) *openapi.Medicine {
	return &openapi.Medicine{
		Id:           medicine.ID,
		MedicineName: medicine.MedicineName,
		MedicineType: medicine.MedicineType,
	}
}

func toMedicinesResponse(medicines []*dtos.Medicine) []*openapi.Medicine {
	dtos := make([]*openapi.Medicine, len(medicines))

	for i, medicine := range medicines {
		dtos[i] = toMedicineResponse(medicine)
	}

	return dtos
}
