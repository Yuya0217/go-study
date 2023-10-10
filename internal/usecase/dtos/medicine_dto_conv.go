package dtos

import "go-layered-architecture-sample/internal/domain/entity"

func MedicineDTOToEntity(m Medicine) *entity.Medicine {
	return &entity.Medicine{
		ID:           m.ID,
		MedicineName: m.MedicineName,
		MedicineType: m.MedicineType,
	}
}
func EntityToMedicineDTO(m *entity.Medicine) *Medicine {
	return &Medicine{
		ID:           m.ID,
		MedicineName: m.MedicineName,
		MedicineType: m.MedicineType,
	}
}

func EntityToMedicinesOutput(es []*entity.Medicine) []*Medicine {
	outputs := make([]*Medicine, len(es))

	for i, medicine := range es {
		outputs[i] = EntityToMedicineDTO(medicine)
	}

	return outputs
}
