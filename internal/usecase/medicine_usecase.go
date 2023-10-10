package usecase

import (
	"context"

	"go-layered-architecture-sample/internal/gateway/rdb"
	"go-layered-architecture-sample/internal/usecase/dtos"
	"go-layered-architecture-sample/pkg/logger"

	"go-layered-architecture-sample/internal/domain/entity"
	"go-layered-architecture-sample/internal/domain/repository"
)

//go:generate mockgen -source=$GOFILE -destination=./mocks/mock_$GOFILE
type MedicineUseCase interface {
	Find(ctx context.Context, params dtos.MedicineFindParams) ([]*dtos.Medicine, error)
	GetByID(ctx context.Context, id int) (*dtos.Medicine, error)
	Create(ctx context.Context, input dtos.Medicine) (*dtos.Medicine, error)
	Update(ctx context.Context, input dtos.Medicine) (*dtos.Medicine, error)
	Delete(ctx context.Context, id int) error
}

type medicineUseCase struct {
	medicineRepository  repository.MedicineRepository
	transactionExecutor repository.TransactionExecutor
	logger              logger.Logger
}

func NewMedicineUseCase(
	medicineRepository repository.MedicineRepository,
	transactionExecutor repository.TransactionExecutor,
	logger logger.Logger,
) MedicineUseCase {
	return &medicineUseCase{
		medicineRepository:  medicineRepository,
		transactionExecutor: transactionExecutor,
		logger:              logger,
	}
}

func (uc *medicineUseCase) Find(ctx context.Context, params dtos.MedicineFindParams) ([]*dtos.Medicine, error) {
	medicines, err := uc.medicineRepository.Find(ctx, rdb.MedicineFindParams{
		MedicineName: params.MedicineName,
		MedicineType: params.MedicineType,
	})

	if err != nil {
		uc.logger.Errorf("failed to medicineRepository.Find: %w", err)
		return nil, entity.NewAppError(entity.MedicineFindFailed, nil)
	}

	outputs := dtos.EntityToMedicinesOutput(medicines)

	return outputs, nil
}

func (uc *medicineUseCase) GetByID(ctx context.Context, id int) (*dtos.Medicine, error) {
	medicine, err := uc.medicineRepository.GetByID(ctx, id)
	if err != nil {
		uc.logger.Errorf("failed to medicineRepository.GetByID: %w", err)
		return nil, entity.NewAppError(entity.MedicineNotFound, nil)
	}

	output := dtos.EntityToMedicineDTO(medicine)

	return output, nil
}

func (uc *medicineUseCase) Create(ctx context.Context, input dtos.Medicine) (*dtos.Medicine, error) {
	medicineEntity := dtos.MedicineDTOToEntity(input)

	createdMedicine, err := uc.medicineRepository.Create(ctx, medicineEntity)
	if err != nil {
		uc.logger.Errorf("failed to medicineRepository.Create: %w", err)
		return nil, entity.NewAppError(entity.MedicineCreateFailed, nil)
	}

	output := dtos.EntityToMedicineDTO(createdMedicine)

	return output, nil
}

func (uc *medicineUseCase) Update(ctx context.Context, input dtos.Medicine) (*dtos.Medicine, error) {
	medicineEntity := dtos.MedicineDTOToEntity(input)

	createdMedicine, err := uc.medicineRepository.Update(ctx, medicineEntity)
	if err != nil {
		uc.logger.Errorf("failed to medicineRepository.Update: %w", err)
		return nil, entity.NewAppError(entity.MedicineUpdateFailed, nil)
	}

	output := dtos.EntityToMedicineDTO(createdMedicine)

	return output, nil
}

func (uc *medicineUseCase) Delete(ctx context.Context, id int) error {
	err := uc.medicineRepository.Delete(ctx, id)
	if err != nil {
		uc.logger.Errorf("failed to medicineRepository.Delete: %w", err)
		return entity.NewAppError(entity.MedicineDeleteFailed, nil)
	}

	return nil
}
